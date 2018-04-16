package main

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
	//"sync/atomic"
	//"time"
)

// Stub is ...
type Stub struct {
	DoEncrypt func(string, string) (string, error)
	DoDecrypt func(string, string) (string, error)
	//AsyncHello func(func(string, error), string) `name:"hello"`
}

func main() {
	client := rpc.NewClient("tcp://127.0.0.1:1036/")
	var stub *Stub
	client.UseService(&stub)
	/*
		stub.AsyncHello(func(result string, err error) {
			fmt.Println(result, err)
		}, "async world")
	*/
	fmt.Println(stub.DoEncrypt("email4", "mytestname"))

	data := "0vIc8x3pNm32gj4jncf/0yV941CXQflA/j8qxLBdF3WuNFnKZZjbBhtCsFsFDyknZMAATS2fygyuGGYW3oWKwfAravq73xOrKvU+LGqei+hLl24/ilL8alO2lS2fmvUs7uO+olL3ShfgO7++CpoJoIiiii9GdXfkuwEI/v6u5k8="
	fmt.Println(stub.DoDecrypt("email4", data))
	/*
		start := time.Now()
		var n int32 = 500000
		done := make(chan bool)
		for i := 0; i < 500000; i++ {
			stub.AsyncHello(func(result string, err error) {
				if atomic.AddInt32(&n, -1) == 0 {
					done <- true
				}
			}, "async world")
		}
		<-done
		stop := time.Now()
		fmt.Println((stop.UnixNano() - start.UnixNano()) / 1000000)
	*/
}
