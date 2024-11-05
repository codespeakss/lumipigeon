package main

import (
	"fmt"
	"log"
	"lumipigeon/link"
)

func printVersion() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("BuildTime: %s\n", BuildTime)
	fmt.Printf("GitBranch: %s\n", GitBranch)
	fmt.Printf("GitCommit: %s\n", GitCommit)
}

func init() {
	printVersion()

	//var e error
}

func main() {
	link.InitConn(link.DefaultUrlPrefix, link.DefaultUsername, link.DefaultPassword)

	// 持续收集当前 node信息 并发送至平台端

	log.Println("[main] waiting select{}")
	select {}
}
