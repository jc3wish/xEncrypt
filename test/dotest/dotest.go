package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	//"strconv"
	"sync/atomic"
	"time"
)

// Stub is ...
type Stub struct {
	DoEncrypt      func(string, string) (string, error)
	DoDecrypt      func(string, string) (string, error)
	AsyncDoEncrypt func(func(string, error), string, string) `name:"DoEncrypt"`
}

func main() {
	doAsyncTest()
}

func doSyncTest() {
	client := rpc.NewClient("tcp://127.0.0.1:1036/")
	var stub *Stub
	client.UseService(&stub)
	data, err := stub.DoEncrypt("email1", "jadsdsdfsdf@126.com")
	if err == nil {
		d, err2 := stub.DoDecrypt("email1", data)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(d)
		}
	} else {
		fmt.Println(err)
	}
}

func doAsyncTest() {
	client := rpc.NewClient("tcp://127.0.0.1:1036/")
	var stub *Stub
	client.UseService(&stub)
	start := time.Now()
	var n int32 = 50000
	done := make(chan bool)
	for i := 0; i < 50000; i++ {
		stub.AsyncDoEncrypt(func(result string, err error) {
			//fmt.Println(result)
			if atomic.AddInt32(&n, -1) == 0 {
				done <- true
			}
		}, "email1", "sdfsdfsdfsf")
	}
	<-done
	stop := time.Now()
	fmt.Println((stop.UnixNano() - start.UnixNano()) / 1000000)
}
