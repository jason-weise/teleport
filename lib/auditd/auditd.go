/*
 *
 * Copyright 2022 Gravitational, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * /
 */

package auditd

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"os/user"
	"syscall"

	"github.com/mdlayher/netlink"
	"github.com/mdlayher/netlink/nlenc"

	"github.com/gravitational/trace"
	log "github.com/sirupsen/logrus"
)

const (
	AUDIT_GET        = 1000
	AUDIT_USER_END   = 1106
	AUDIT_USER_LOGIN = 1112
	AUDIT_USER_ERR   = 1109
)

const (
	success = "success"
	failed  = "failed"
)

// sshd
// type=USER_LOGIN msg=audit(1657493525.955:466): pid=16059 uid=0 auid=4294967295 ses=4294967295 subj==unconfined msg='op=login acct="jnyckowski" exe="/usr/sbin/sshd" hostname=? addr=127.0.0.1 terminal=sshd res=failed'UID="root" AUID="unset"
// type=USER_START msg=audit(1657493584.668:474): pid=16059 uid=0 auid=1000 ses=11 subj==unconfined msg='op=PAM:session_open grantors=pam_selinux,pam_loginuid,pam_keyinit,pam_permit,pam_umask,pam_unix,pam_systemd,pam_mail,pam_limits,pam_env,pam_env,pam_selinux,pam_tty_audit acct="jnyckowski" exe="/usr/sbin/sshd" hostname=127.0.0.1 addr=127.0.0.1 terminal=ssh res=success'UID="root" AUID="jnyckowski"
// type=USER_END msg=audit(1657744078.476:5916): pid=275303 uid=0 auid=1000 ses=118 subj==unconfined msg='op=PAM:session_close grantors=pam_selinux,pam_loginuid,pam_keyinit,pam_permit,pam_umask,pam_unix,pam_systemd,pam_mail,pam_limits,pam_env,pam_env,pam_selinux,pam_tty_audit acct="jnyckowski" exe="/usr/sbin/sshd" hostname=127.0.0.1 addr=127.0.0.1 terminal=ssh res=success'UID="root" AUID="jnyckowski"

type AuditDClient struct {
	conn *netlink.Conn

	pid          int
	execName     string
	hostname     string
	user         string
	teleportUser string
	address      string
	ttyName      string
}

type auditStatus struct {
	Mask                  uint32 /* Bit mask for valid entries */
	Enabled               uint32 /* 1 = enabled, 0 = disabled */
	Failure               uint32 /* Failure-to-log action */
	Pid                   uint32 /* pid of auditd process */
	RateLimit             uint32 /* messages rate limit (per second) */
	BacklogLimit          uint32 /* waiting messages limit */
	Lost                  uint32 /* messages lost */
	Backlog               uint32 /* messages waiting in queue */
	Version               uint32 /* audit api version number */ // feature bitmap
	BacklogWaitTime       uint32 /* message queue wait timeout */
	BacklogWaitTimeActual uint32 /* message queue wait timeout */
}

func NewAuditDClient(teleportUser, ttyName string) (*AuditDClient, error) {
	conn, err := netlink.Dial(syscall.NETLINK_AUDIT, nil)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	status, err := getAuditStatus(conn)
	if err != nil {
		return nil, trace.Errorf("failed to get audutd state: %v", err)
	}

	if status.Enabled != 1 {
		return nil, trace.Errorf("audutd is disabled")
	}
	log.Warnf("auditd is enabled")

	pid := os.Getpid()
	execName, err := os.Executable()
	if err != nil {
		log.WithError(err).Warn("failed to get executable name")
		execName = "?"
	}
	hostname, err := os.Hostname()
	if err != nil {
		log.WithError(err).Warn("failed to get hostname")
	}
	currentUser, err := user.Current()
	if err != nil {
		log.WithError(err).Warn("failed to get the current user")
	}

	addr := "127.0.0.1"
	if ttyName == "" {
		ttyName = "ssh"
	}

	return &AuditDClient{
		conn:         conn,
		pid:          pid,
		execName:     execName,
		hostname:     hostname,
		user:         currentUser.Username, //TODO: fix me
		teleportUser: teleportUser,
		address:      addr,
		ttyName:      ttyName,
	}, nil
}

func getAuditStatus(conn *netlink.Conn) (*auditStatus, error) {
	resp, err := conn.Execute(netlink.Message{
		Header: netlink.Header{
			Type:  netlink.HeaderType(AUDIT_GET),
			Flags: netlink.Request | netlink.Acknowledge,
		},
		Data: nil,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}

	log.Warnf("AuditGetResp: %v\n", resp)

	msgs, err := conn.Receive()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	log.Warnf("msgs: %v\n", msgs)

	if len(msgs) != 1 {
		return nil, trace.Errorf("returned wrong messages number, expected 1, got: %d", len(msgs))
	}

	byteOrder := nlenc.NativeEndian()
	status := &auditStatus{}

	payload := bytes.NewReader(msgs[0].Data[:])
	if err := binary.Read(payload, byteOrder, status); err != nil {
		return nil, trace.Wrap(err)
	}

	log.Warnf("status: %+v\n", status)

	return status, nil
}

func (c *AuditDClient) SendLogin() error {
	log.Warnf("sending login audit event")

	const msgDataTmpl = "op=%s acct=\"%s\" teleportUser=\"%s\" exe=%s hostname=%s addr=%s terminal=%s res=%s"
	const op = "login"

	MsgData := []byte(fmt.Sprintf(msgDataTmpl, op, c.user, c.teleportUser, c.execName, c.hostname, c.address, c.ttyName, success))

	return c.sendMsg(AUDIT_USER_LOGIN, MsgData)
}

func (c *AuditDClient) SendLoginFailed() error {
	log.Warnf("sending login failed audit event")

	const msgDataTmpl = "op=%s acct=\"%s\" exe=%s hostname=%s addr=%s terminal=%s res=%s"
	const op = "login"

	MsgData := []byte(fmt.Sprintf(msgDataTmpl, op, c.user, c.execName, c.hostname, c.address, c.ttyName, failed))

	return c.sendMsg(AUDIT_USER_LOGIN, MsgData)
}

// type=USER_ERR msg=audit(1658343692.733:471): pid=7113 uid=0 auid=4294967295 ses=4294967295 subj=? msg='op=PAM:bad_ident grantors=? acct="?" exe="/usr/sbin/sshd" hostname=::1 addr=::1 terminal=ssh res=failed'UID="root" AUID="unset"

func (c *AuditDClient) SendInvalidUser() error {
	log.Warnf("sending invalid user audit event")

	const msgDataTmpl = "op=%s acct=\"%s\" exe=%s hostname=%s addr=%s terminal=%s res=%s"
	const op = "invalid_user"

	MsgData := []byte(fmt.Sprintf(msgDataTmpl, op, c.user, c.execName, c.hostname, c.address, c.ttyName, failed))

	return c.sendMsg(AUDIT_USER_ERR, MsgData)
}

// type=USER_END msg=audit(1657744078.476:5916): pid=275303 uid=0 auid=1000 ses=118 subj==unconfined msg='op=PAM:session_close grantors=pam_selinux,pam_loginuid,pam_keyinit,pam_permit,pam_umask,pam_unix,pam_systemd,pam_mail,pam_limits,pam_env,pam_env,pam_selinux,pam_tty_audit acct="jnyckowski" exe="/usr/sbin/sshd" hostname=127.0.0.1 addr=127.0.0.1 terminal=ssh res=success'UID="root" AUID="jnyckowski"

func (c *AuditDClient) SendSessionEnd() error {
	log.Warnf("sending login audit event")

	//const msgDataTmpl = "op=PAM:session_close grantors=pam_selinux,pam_loginuid,pam_keyinit,pam_permit,pam_umask,pam_unix,pam_systemd,pam_mail,pam_limits,pam_env,pam_env,pam_selinux,pam_tty_audit acct=\"jnyckowski\" exe=\"/usr/sbin/sshd\" hostname=127.0.0.1 addr=127.0.0.1 terminal=ssh res=success'UID=\"root\" AUID=\"jnyckowski\""
	const msgDataTmpl = "op=%s acct=\"%s\" exe=%s hostname=%s addr=%s terminal=%s res=%s"
	const op = "session_close"

	MsgData := []byte(fmt.Sprintf(msgDataTmpl, op, c.user, c.execName, c.hostname, c.address, c.ttyName, success))

	return c.sendMsg(AUDIT_USER_END, MsgData)
}

func (c *AuditDClient) sendMsg(eventType netlink.HeaderType, MsgData []byte) error {
	msg := netlink.Message{
		Header: netlink.Header{
			Type:  eventType,
			Flags: syscall.NLM_F_REQUEST | syscall.NLM_F_ACK,
		},
		Data: MsgData,
	}

	resp, err := c.conn.Execute(msg)
	if err != nil {
		return trace.Wrap(err)
	}

	if len(resp) != 1 {
		return fmt.Errorf("unexpected number of responses from kernel for status request: %d, %v", len(resp), resp)
	}
	log.Infof("reply: %v", resp)

	return nil
}

func (c *AuditDClient) Close() error {
	return c.conn.Close()
}
