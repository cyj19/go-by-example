package rpc

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

/*
	从自定义协议中读写数据
	网络字节流：header uint32 | data []byte
*/

type Session struct {
	Conn net.Conn
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		Conn: conn,
	}
}

// 向连接中写数据
func (s *Session) Write(data []byte) error {
	// 构建整个数据包
	buf := make([]byte, 4+len(data))
	// 前4个字节为协议头部，记录data的长度
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	// 写入数据
	copy(buf[4:], data)
	_, err := s.Conn.Write(buf)
	return err
}

// 从连接中读数据
func (s *Session) Read() ([]byte, error) {
	// 读协议头
	header := make([]byte, 4)
	_, err := io.ReadFull(s.Conn, header)
	if err != nil {
		return nil, err
	}
	dataLen := binary.BigEndian.Uint32(header)
	fmt.Println("数据长度为：", dataLen)
	// 读取数据
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.Conn, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
