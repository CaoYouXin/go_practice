package main

import (
	"fmt"
	"net"
	"net/url"
)

func parseOneURL(str string) {
	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}

	fmt.Println("-------------")

	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	pwd, _ := u.User.Password()
	fmt.Println(pwd)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("%v=%v;\n", k, v)
	}

	fmt.Println()
}

func main() {
	parseOneURL("postgres://user:pass@host.com:5432/path?k=v#f")
	parseOneURL("http://www.baidu.com/?ssl=personal")
}
