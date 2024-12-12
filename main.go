package main

import (
	"log"
	"os"

	"github.com/armon/go-socks5"
)

func main() {
	// 从环境变量获取认证信息
	username := os.Getenv("SOCKS5_USERNAME")
	password := os.Getenv("SOCKS5_PASSWORD")

	// 创建SOCKS5配置
	conf := &socks5.Config{}
	
	// 如果设置了用户名和密码，则启用认证
	if username != "" && password != "" {
		creds := socks5.StaticCredentials{
			username: password,
		}
		auth := socks5.UserPassAuthenticator{Credentials: creds}
		conf.AuthMethods = []socks5.Authenticator{auth}
	}

	// 创建SOCKS5服务器
	server, err := socks5.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	// 获取端口，如果未设置则使用默认端口1080
	port := os.Getenv("PORT")
	if port == "" {
		port = "1080"
	}

	// 启动服务器
	log.Printf("Starting SOCKS5 server on :%s", port)
	if err := server.ListenAndServe("tcp", ":"+port); err != nil {
		log.Fatal(err)
	}
}
