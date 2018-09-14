package plugins

import (
	"crack/model"
	"crack/util"
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func ScanFtp(s model.ScanResult) (result model.ScanResult) {

	util.Info("start login  FTP  %s %s %s", s.Hostport, s.Username, s.Password)
	conn, err := ftp.DialTimeout(fmt.Sprintf("%v", s.Hostport), time.Second*8)
	if err == nil {
		err = conn.Login(s.Username, s.Password)
		if err == nil {
			defer conn.Logout()
			s.Success = true
			util.Success("Found FTP %s %s %s", s.Hostport, s.Username, s.Password)
			return s
		}
	}
	util.Info("FTP password error %s %s %s", s.Hostport, s.Username, s.Password)
	s.Success = false
	return s
}
