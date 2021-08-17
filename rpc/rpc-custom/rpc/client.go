package rpc

import (
	"net"
	"reflect"
)

/*
	RPC客户端核心功能
	1. 只拥有函数原型(函数原型相当于函数声明，有函数名，形参列表，且不需要函数体)
	2. 使用reflect.MakeFunc()完成原型到函数的调用 (有关MakeFunc可参考标准库文档 https://studygolang.com/pkgdoc)
*/

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

/*
	实现通用的RPC客户端
	rpcName：函数名
	fPtr：指向函数原型的指针
	example:
	var select func selectUser(int)
	Call("selectUser", &select)
*/
func (c *Client) Call(rpcName string, fPtr interface{}) {
	// 通过reflect.ValueOf().Elem()获取fPtr指向的函数原型
	fn := reflect.ValueOf(fPtr).Elem()

	// 从服务端获取rpcName对应的函数实现
	f := func(args []reflect.Value) []reflect.Value {

		// 处理参数，改为与RPCData.Args类型相同
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			// 返回反射持有的值并添加到切片
			inArgs = append(inArgs, arg.Interface())
		}

		// 连接服务端
		cliSession := NewSession(c.conn)
		// 编码
		reqRPCData := RPCData{
			Name: rpcName,
			Args: inArgs,
		}
		reqByte, err := encode(reqRPCData)
		if err != nil {
			panic(err)
		}
		// 向服务端写入数据
		err = cliSession.Write(reqByte)
		if err != nil {
			panic(err)
		}

		// 获取服务端的返回值
		respByte, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		// 解码
		respRPCData, err := decode(respByte)
		if err != nil {
			panic(err)
		}
		// 处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRPCData.Args))
		for i, arg := range respRPCData.Args {
			// 必须进行nil转换
			if arg == nil {
				// reflect.Zero()会返回类型的零值的value
				// .Out()会返回函数输出的参数类型
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}

		return outArgs
	}
	// 使用MakeFunc完成函数原型到函数定义的内部转换
	v := reflect.MakeFunc(fn.Type(), f)
	// 函数赋值给fPtr
	fn.Set(v)
}
