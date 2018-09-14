package plugins

import (
	"crack/model"
	"crack/util"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
)

func ScanPostgres(s model.ScanResult) (result model.ScanResult) {
	util.Info("start login  Postgresql   %s %s %s", s.Hostport, s.Username, s.Password)


	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=%v", s.Username,
		s.Password, s.Hostport, "postgres", "disable")

	db, err := sql.Open("postgres", dataSourceName)

	if err := db.Ping(); err == nil {
		s.Success = true
		util.Success("Found Postgresql %s %s %s", s.Hostport, s.Username, s.Password)
		defer db.Close()
		return s
	}
	util.Info("%s %s", err.Error(), s.Hostport)
	s.Success = false
	util.Info("Postgresql password error %s %s %s", s.Hostport, s.Username, s.Password)

	return s

}
