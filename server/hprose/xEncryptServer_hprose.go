package main

import (
	"flag"
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	"github.com/xEncrypt/encrypt"
)

func hello(name string) string {
	return "hello " + name
}

func main() {
	Host := flag.String("host", "0.0.0.0", "listen ip")
	Port := flag.String("port", "1036", "listen port")
	flag.Parse()
	IPAndPort := "tcp4://" + *Host + ":" + *Port + "/"
	fmt.Println("xEncrypt start...")
	fmt.Println(IPAndPort)
	server := rpc.NewTCPServer(IPAndPort)
	server.AddFunction("DoEncrypt", encrypt.DoEncrypt)
	server.AddFunction("DoDecrypt", encrypt.DoDecrypt)
	server.AddFunction("hello", hello)
	server.Start()
}
