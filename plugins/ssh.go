package plugins

import (
	"golang.org/x/crypto/ssh"

	"crack/model"
	"crack/util"
	"fmt"
	"time"
)

func ScanSsh(s model.ScanResult) (result model.ScanResult) {
	config := &ssh.ClientConfig{
		User: s.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		Timeout: time.Second * 5,
	}
	util.Info("start login  %s %s %s", s.Hostport, s.Username, s.Password)
	c, err := ssh.Dial("tcp", fmt.Sprintf("%v", s.Hostport), config)
	if err == nil {
		defer c.Close()
		util.Success("Found SSH %s %s %s", s.Hostport, s.Username, s.Password)
		s.Success = true
		return s
	}
	util.Info("SSH password error %s %s %s", s.Hostport, s.Username, s.Password)
	s.Success = false
	return s
}
