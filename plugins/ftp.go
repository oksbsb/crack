package plugins

import (
	"crack/model"
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func ScanFtp(s model.ScanResult) (result model.ScanResult) {



	conn, err := ftp.DialTimeout(fmt.Sprintf("%v", s.Hostport), time.Second*2)
	if err == nil {
		err = conn.Login(s.Username, s.Password)
		if err == nil {
			defer conn.Logout()
			s.Success = true
			return s
		}
	}
	s.Success = false
	return s
}
