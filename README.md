#crack 基于golang的多并发爆破工具

## Install

git clone  https://github.com/yanweijin/crack</br>

go build main.go </br>

## help

./main --help</br>
  -P string</br>
        filename list or username</br>
  -U string</br>
        filename list or username</br>
  -i string</br>
        ip addr 192.168.1.1/24，192.168.1-255，192.168.1.1</br>
  -p int
        port 21,22，445，3306，1433 (default 22)</br>
  -s string
        service ssh，smb，mssql，mysql</br>
  -t int
        thread (default 100)</br>

./main -i 192.168.1.1/24 -p 22 -U /tmp/user.txt -P /tmp/pass.txt  -s ssh</br>



