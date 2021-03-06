/**
* TCP  Client
 */

package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {

	client, err := rpc.Dial("tcp", "127.0.0.1:1234") //获取服务
	if err != nil {
		log.Fatal("错误:", err)
	}
	// 同步调用
	args := Args{17, 8}
	var reply int
	err = client.Call("RPC.Multiply", args, &reply)
	if err != nil {
		log.Fatal("RPC error:", err)
	}
	fmt.Printf("结果: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("RPC.Divide", args, &quot)
	if err != nil {
		log.Fatal("RPC error:", err)
	}
	fmt.Printf("结果: %d/%d=%d 余数 %d\n", args.A, args.B, quot.Quo, quot.Rem)

}
