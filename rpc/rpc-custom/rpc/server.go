package rpc

import (
	"log"
	"net"
	"reflect"
)

/*
	RPC服务端核心功能
	1. 维护函数map
	2. 解析客户端的数据
	3. 调用相应的函数，得到编码后返回值打包给客户端
*/

// 声明服务端
type Server struct {
	// 地址
	addr string
	// map 维护已注册的函数
	funcs map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{
		addr:  addr,
		funcs: make(map[string]reflect.Value),
	}
}

// 注册用于RPC交互的函数
func (s *Server) Register(rpcName string, f interface{}) {
	// 从funcs中查找，已有键
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	// 把f函数添加到函数map中
	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal

}

// 运行服务端，等待调用
func (s *Server) Run() {
	// 创建服务端
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("监听:%s err:%v \n", s.addr, err)
		return
	}

	for {
		// 循环等待连接
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		// 创建会话
		session := NewSession(conn)
		// 解析客户端的数据
		reqData, err := session.Read()
		if err != nil {
			return
		}
		rpcData, err := decode(reqData)
		if err != nil {
			return
		}

		// 根据函数名，调用对应函数
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			log.Printf("函数:%s 不存在 \n", rpcData.Name)
			return
		}

		// 遍历参数，把参数类型改为[]reflect.Value
		args := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			v := reflect.ValueOf(arg)
			args = append(args, v)

		}

		// 通过反射调用函数
		out := f.Call(args)
		outArgs := make([]interface{}, 0, len(out))
		// 遍历结果，因为RPCData的Args是[]interface{}，需要修改结果类型
		for _, item := range out {
			outArgs = append(outArgs, item.Interface())
		}

		// 编码结果
		respRPCData := RPCData{rpcData.Name, outArgs}
		respByte, err := encode(respRPCData)
		if err != nil {
			return
		}
		// 将编码结果返回给客户端
		err = session.Write(respByte)
		if err != nil {
			return
		}
	}
}
