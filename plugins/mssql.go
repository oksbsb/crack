package plugins

import (
	_ "github.com/denisenkom/go-mssqldb"
	"crack/model"
	"crack/util"
	"fmt"
	"database/sql"
)

func ScanMssql(s model.ScanResult) (result model.ScanResult) {

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8", s.Username,
		s.Password, s.Hostport, "mysql")

	db, err := sql.Open("mssql", dataSourceName)

	if err := db.Ping(); err == nil{
		s.Success = true
		defer db.Close()
		return  s
	}
	util.Info("%s %s",err.Error(),s.Hostport)
	s.Success = false

	return s
	
}
