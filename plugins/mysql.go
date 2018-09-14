package plugins

import (
	"crack/model"
	"crack/util"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ScanMysql(s model.ScanResult) (result model.ScanResult) {
	util.Info("start login  MYSQL   %s %s %s", s.Hostport, s.Username, s.Password)

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", s.Username,
		s.Password, s.Hostport, "mysql")

	db, err := sql.Open("mysql", dataSourceName)

	if err := db.Ping(); err == nil {
		s.Success = true
		util.Success("Found MYSQL %s %s %s", s.Hostport, s.Username, s.Password)
		defer db.Close()
		return s
	}
	util.Info("%s %s", err.Error(), s.Hostport)
	s.Success = false
	util.Info("MYSQL password error %s %s %s", s.Hostport, s.Username, s.Password)

	return s

}
