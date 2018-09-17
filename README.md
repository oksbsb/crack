# crack 基于golang的多并发爆破工具

## Install

git clone  https://github.com/yanweijin/crack</br>

go build main.go </br>

## help

```
➜  crack git:(master) ✗ go run main.go -i 133.96.1.1/16  --help               
-P string
    filename list or username
-U string
    filename list or username
-i string
    ip addr 192.168.1.1/24，192.168.1-255，192.168.1.1
-p string
    port 21,22，445，3306，1433
-s string
    service ftp  ssh，smb，mssql，mysql, postgresql   default  only scan tcp port (default "scan")
-t int
    thread (default 100)

```

go run main.go -i 192.168.1.1/15  -U root  -P 123456 -s mysql




 


