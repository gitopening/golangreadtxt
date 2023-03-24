package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	//importing custom package
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("error:%s", err)
		}
	}()
}

func getOldStrArr() []string {
	//一次性读取文件
	read, err := ioutil.ReadFile("tenurl.php")
	if err != nil {
		fmt.Printf("文件打开失败", err.Error())
		return nil
	}
	// fmt.Printf(string(read))
	// fmt.Println("test文件中的原始内容：", read)

	oldstrarr := strings.Split(string(read), "\n")
	fmt.Println("test文件中的原始内容长度：", len(oldstrarr))
	return oldstrarr

}
