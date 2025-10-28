package main

import (
	"fmt"
	"os"
	"voceapi/bot"
)

const url = ""
const user = ""
const passwd = ""
const key = ""

func main() {
	// g := group.New(url)
	// g.Login(user, passwd)
	// data, _ := os.ReadFile("YT7575550466029.jpg")
	// path, _ := g.Upload(data, "YT7575550466029.jpg")
	// for {
	// 	var input string
	// 	fmt.Scanln(&input)
	// 	g.SendFile(1, path)
	// }

	b := bot.New(
		url,
		key,
		"/api/bot",
		80,
		message,
	)
	b.Run()
}

func message(b *bot.Bot, gid int64, msg string) {
	fmt.Println(gid, msg)
	data, _ := os.ReadFile("YT7575550466029.jpg")
	path, _ := b.Upload(data, "YT7575550466029.jpg")
	fmt.Println(path)
	b.SendFile(gid, path)
}
