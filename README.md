#crack 基于golang的多并发爆破工具

## Install

git clone  https://github.com/yanweijin/crack</br>

go build main.go </br>

## help

➜  crack ./main --help

  -P string
        filename list or username
  -U string
        filename list or username
  -i string
        ip addr 192.168.1.1/24，192.168.1-255，192.168.1.1
  -p int
        port 21,22，445，3306，1433 (default 22)
  -s string
        service ssh，smb，mssql，mysql
  -t int
        thread (default 100)


./main -i 192.168.1.1/24 -p 22 -U /tmp/user.txt -P /tmp/pass.txt  -s ssh</br>



