package plugins

import (
	"crack/model"
	"crack/util"
	"github.com/stacktitan/smb/smb"
	"strconv"
	"strings"
)

func ScanSmb(s model.ScanResult) (result model.ScanResult) {

	Host := strings.Split(s.Hostport, ":")[0]
	Port, err := strconv.Atoi(strings.Split(s.Hostport, ":")[1])
	if err != nil {
		panic(err)
	}

	options := smb.Options{
		Host:        Host,
		Port:        Port,
		User:        s.Username,
		Password:    s.Password,
		Domain:      "",
		Workstation: "",
	}
	util.Info("start login SMB  %s %s %s", s.Hostport, s.Username, s.Password)
	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			util.Success("Found SMB %s %s %s", s.Hostport, s.Username, s.Password)
			s.Success = true
			return s

		}
	}
	util.Info("SMB password error %s %s %s", s.Hostport, s.Username, s.Password)
	s.Success = false
	return s

}
