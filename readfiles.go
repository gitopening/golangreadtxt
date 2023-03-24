package main

import (
	"fmt"
	"io/ioutil"
	//importing custom package
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("error:%s", err)
		}
	}()
}

func main() {
	//一次性读取文件
	read, err := ioutil.ReadFile("tenurl.php")
	if err != nil {
		fmt.Printf("文件打开失败", err.Error())
		return
	}
	fmt.Printf(string(read))

}
