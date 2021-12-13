/**
 * @Author: cyj19
 * @Date: 2021/12/13 16:01
 */

// 逃逸分析例子
// 使用命令进行分析：go build -gcflags '-m -l' main.go

package main

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func main() {
	_ = GetUserInfo(User{
		ID:     88022777,
		Name:   "cyj19",
		Avatar: "https://avatars.githubusercontent.com/u/88022777?v=4",
	})
	//str := new(string)
	//*str = "cyj19"
	//fmt.Println(str)
}

func GetUserInfo(u User) *User {
	return &u
}
