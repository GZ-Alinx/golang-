package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("127.0.0.1","8888")) // 地址和端口拼接
	fmt.Println(net.LookupAddr("127.0.0.1")) // 通过ip查找域名
	fmt.Println(net.LookupHost("www.baidu.com")) // 通过域名找地址
	fmt.Println(net.LookupHost("106.227.9.9")) // 通过域名找地址



	// ip IP  类型处理  三种方式  在TYPE的处
	var ip net.IP = net.ParseIP("1.1.1.1")
	fmt.Println(ip) // 解析ip


	// ip段 CIDR  格式：  ip/mask

	ip, ipnet, err := net.ParseCIDR("192.168.0.10/32")
	fmt.Println(ipnet,ip,err)
	ip, ipnet, err = net.ParseCIDR("192.168.0.10/24")
	fmt.Println(ipnet,ip,err)

	fmt.Println(ipnet.IP)
	fmt.Println(ipnet.Mask)
	ipnet.Contains(net.ParseIP("192.168.0.88"))
	fmt.Println(ipnet,ip,err)


	// Addr 网络地址
	addrs,err := net.InterfaceAddrs()
	for idx,addr := range addrs {
		fmt.Println(idx,addr)
	}


	// Interfaces 网络地址 本地获取
	intfs,err := net.Interfaces()
	for idx,intf := range intfs {
		fmt.Println(intf.Name)
		fmt.Println(intf.Index)
		fmt.Println(idx)

		// 本地地址信息获取
		intf.Addrs()
		intf.MulticastAddrs()
	}

}
