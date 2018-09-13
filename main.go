package main

import (
	"crack/model"
	"crack/plugins"
	"crack/util"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	passchan    = make(chan model.ScanResult, 0)
	ipli        = make(chan string, 0)
	portresult  <-chan model.ServerResult
	t           int
	scan        func(result model.ScanResult) model.ScanResult
	defaultport int
)

func usage() {
	log.Fatalf("Usage: main.exe --help")
}

func main() {

	user := flag.String("U", "", "filename list or username")
	pass := flag.String("P", "", "filename list or username")
	ip := flag.String("i", "", "ip addr 192.168.1.1/24，192.168.1-255，192.168.1.1")
	port := flag.String("p", "", "port 21,22，445，3306，1433")
	service := flag.String("s", "scan", "service ftp  ssh，smb，mssql，mysql, postgresql   default  only scan tcp port")
	flag.IntVar(&t, "t", 100, "thread")

	flag.Parse()
	if *ip == "" {
		os.Exit(1)
	}

	switch *service {

	case "ssh":
		defaultport = 22
		scan = plugins.ScanSsh
	case "smb":
		defaultport = 445
		scan = plugins.ScanSmb
	case "ftp":
		defaultport = 21
		scan = plugins.ScanFtp
	case "mysql":
		defaultport = 3306
		scan = plugins.ScanMysql
	case "mssql":
		defaultport = 1433
		scan = plugins.ScanMssql
	case "postgresql":
		defaultport = 5432
		scan = plugins.ScanPostgres
	default:
		 scan = nil
	}

	if *port == "" {
		go makeiplist(*ip, fmt.Sprintf("%d", defaultport))
	} else {
		go makeiplist(*ip, *port)

	}
	portresult = util.Portcheck(t, ipli)


	 if scan == nil {
		for _ = range portresult {
		}
		os.Exit(1)
	}

	go makchan(user, pass, scan)

	f, err := os.OpenFile("result.txt",  os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	resutlt := util.Passattack(t, passchan)
	for i := range resutlt {
		if i.Success == true {
			f.WriteString(fmt.Sprintf("%s\t%s\t%s\r\n",i.Hostport,i.Username,i.Password))
		}

	}

}

func makeiplist(ip string, port string) {
	ipip, err := util.Makeip(ip)
	if err != nil {
		usage()
		os.Exit(1)
	}

	for _, i := range ipip {
		ipli <- fmt.Sprintf("%s:%s", i, port)
	}
	close(ipli)

}

func makchan(user *string, pass *string, scan func(result model.ScanResult) model.ScanResult) {

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
