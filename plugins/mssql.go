package plugins

import (
	"crack/model"
	"crack/util"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"strconv"
	"strings"
	"log"
)

func ScanMssql(s model.ScanResult) (result model.ScanResult) {

	util.Info("start login  MSSQL   %s %s %s", s.Hostport, s.Username, s.Password)

	Host := strings.Split(s.Hostport, ":")[0]
	Port, _ := strconv.Atoi(strings.Split(s.Hostport, ":")[1])

	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", Host,
		Port, s.Username, s.Password, "master")

	db, err := sql.Open("mssql", dataSourceName)
	if err !=nil {
		s.Success = false
		util.Info("MSSQL password error %s %s %s", s.Hostport, s.Username, s.Password)
		return s
	}
	if err := db.Ping(); err == nil {
		s.Success = true
		util.Success("Found MSSQL %s %s %s", s.Hostport, s.Username, s.Password)
		defer db.Close()
		return s
	}
	s.Success = false
	util.Info("MSSQL password error %s %s %s", s.Hostport, s.Username, s.Password)

	return s

}
