package main

import (
	"crack/model"
	"crack/util"
	"flag"
	"fmt"
	"crack/plugins"
	"os"
	"log"
)

var (
	passchan = make(chan model.ScanResult, 0)
	ipli     = make(chan string, 0)
 	portresult   <- chan model.ServerResult
	t int
)

func usage()  {
	log.Fatalf("Usage: main.exe --help")
}

func main() {

	user := flag.String("U", "", "filename list or username")
	pass := flag.String("P", "", "filename list or username")
	ip := flag.String("i", "", "ip addr 192.168.1.1/24，192.168.1-255，192.168.1.1")
	port := flag.Int("p", 22, "port 21,22，445，3306，1433")
	service := flag.String("s", "scan", "service  ssh，smb，mssql，mysql, postgresql   default  only scan tcp port")
	flag.IntVar (&t,"t", 100, "thread")

	flag.Parse()
	if *ip=="" {
		os.Exit(1)
	}

	go makeiplist(ip, port)
	portresult = util.Portcheck(t, ipli)

	switch *service {
	case "ssh":
		go makchan(user, pass,plugins.ScanSsh)
	case "smb":
		go makchan(user, pass,plugins.ScanSmb)
	case "ftp":
		go makchan(user, pass,plugins.ScanFtp)
	case "mysql":
		go makchan(user, pass,plugins.ScanMysql)
	case "mssql":
		go makchan(user, pass,plugins.ScanMssql)
	case "postgresql":
		go makchan(user, pass,plugins.ScanPostgres)
	default :
		for _ = range portresult {}
		return
	}

	resutlt := util.Passattack(t, passchan)
	for i := range resutlt {
		if i.Success == true {
			fmt.Println(i)
		}

	}

}


func makeiplist(ip *string, port *int) {
	ipip, err := util.Makeip(*ip)
	if err != nil {
		usage()
		os.Exit(1)
	}
	if *port > 65535 {
		usage()
		os.Exit(1)
	}
	if *port < 10 {
		usage()
		os.Exit(1)
	}
	for _, i := range ipip {
		ipli <- fmt.Sprintf("%s:%d", i, *port)
	}
	close(ipli)

}

func makchan(user *string, pass *string,scan func(result model.ScanResult) model.ScanResult) {

	userlist, err := util.Makelist(*user)
	if err != nil {
		fmt.Println("user file error")
		os.Exit(1)
	}
	passlist, err := util.Makelist(*pass)
	if err != nil {
		fmt.Println("user file error")
		os.Exit(1)
	}
	if len(userlist) == 0 {
		fmt.Println("passw file error")
		os.Exit(1)
	}

	if len(passlist) == 0 {
		fmt.Println("passw file error")
		os.Exit(1)
	}

	for i := range portresult {
		for _, u := range userlist {
			for _, p := range passlist {
				if i.Open == true {
					passchan <- model.ScanResult{Hostport: i.Hostport, Username: u, Password: p, Server: scan}
				}
			}
		}
	}
	close(passchan)
}
