package util

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Makeip(ip string) ([]string, error) {

	iplist := make([]string, 0)

	if strings.Contains(ip, "-") {
		base := strings.Split(ip, "-")
		baseip := base[0][0 : strings.LastIndex(base[0], ".")+1]
		start, _ := strconv.Atoi(base[0][strings.LastIndex(base[0], ".")+1:])
		end, _ := strconv.Atoi(base[1])
		if end > start {
			for index := start; index < end; index++ {
				iplist = append(iplist, baseip+strconv.Itoa(index))
			}
		} else {
			return iplist, errors.New("ip error -")
		}
	} else if strings.Contains(ip, "/") {
		ip, ipNet, err := net.ParseCIDR(ip)
		if err != nil {
			fmt.Println("invalid CIDR /")
			return iplist, err
		}
		for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
			iplist = append(iplist, ip.String())
		}
	} else {
		trial := net.ParseIP(ip)
		if trial.To4() == nil {
			return iplist, errors.New("ip error")
		}
		iplist = append(iplist, ip)
	}
	return iplist, nil
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
