/**
 * @Author: cyj19
 * @Date: 2021/12/1 14:12
 */

// 命令行TCP聊天室，socket编程中TCP Server的通用代码
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {

	// 监听端口
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

// User 用户
type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

func (u *User) string() string {
	return u.Addr + " , UID: " + strconv.Itoa(u.ID) + " EnterAt: " + u.EnterAt.Format("2006-01-02 15:04:05+8000")
}

// Message 消息
type Message struct {
	OwnerID int    // 发送方
	Content string // 消息内容
}

var (
	enteringChannel = make(chan *User)       // 新用户登记
	leavingChannel  = make(chan *User)       // 用户离开登记
	messageChannel  = make(chan *Message, 8) // 广播普通消息
)

// broadcaster 用于记录聊天用户并进行消息广播
// 1、用户进入 2、用户普通消息  3、用户离开
func broadcaster() {
	users := make(map[*User]struct{})

	for {
		select {
		case user := <-enteringChannel:
			users[user] = struct{}{}
		case user := <-leavingChannel:
			delete(users, user)
			// 关闭该用户的消息channel，避免goroutine泄露
			close(user.MessageChannel)
		case msg := <-messageChannel:
			// 给所有用户发送消息
			for user := range users {
				if msg.OwnerID == user.ID {
					continue
				}
				user.MessageChannel <- msg.Content
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		_ = conn.Close()
	}()

	// 1、构建新用户的实例
	user := &User{
		ID:             GenUserId(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	// 2、向用户发送消息
	go sendMessage(conn, user.MessageChannel)

	// 3、向用户发送欢迎消息，给所有用户告知新用户到来
	user.MessageChannel <- "Welcome " + user.string()
	msg := &Message{
		OwnerID: user.ID,
		Content: "user: " + strconv.Itoa(user.ID) + " has enter",
	}
	messageChannel <- msg

	// 4、将该用户记录到全局的用户队列
	enteringChannel <- user
	// 控制超时自动剔除用户
	userActive := make(chan struct{})
	go func() {
		d := 5 * time.Minute
		timer := time.NewTimer(d)
		for {
			select {
			case <-timer.C:
				_ = conn.Close()
				// 超时退出该goroutine，防止泄露
				return
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	// 5、循环读取用户的输入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg.Content = strconv.Itoa(user.ID) + ": " + input.Text()
		messageChannel <- msg
		userActive <- struct{}{}
	}

	// 读取错误
	if err := input.Err(); err != nil {
		log.Println(err)
	}

	// 6、用户离开，从用户队列中剔除
	leavingChannel <- user
	msg.Content = "user: " + strconv.Itoa(user.ID) + " has left"
	messageChannel <- msg

}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		// 向客户端发送消息
		_, _ = fmt.Fprintln(conn, msg)
	}
}

// 生成ID
var (
	globalId int
	idLocker sync.Mutex
)

func GenUserId() int {
	idLocker.Lock()
	defer idLocker.Unlock()
	globalId++
	return globalId
}
