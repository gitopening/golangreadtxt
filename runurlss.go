package main

import (
	"fmt"
	"os"
	"sync" //importing custom package
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("error:%s", err)
		}
	}()
}

var waitgroup sync.WaitGroup

func main() {
	gotoStrarr := getAllUrlStrArr()
	// if err = ioutil.WriteFile("./endphp", []string(gotoStrarr), 0666); err != nil {
	// 	fmt.Println("Writefile Error =", err)
	// 	return
	// }
	// os.Create("./endphp") //创建文件
	dstFile, err := os.Create("./end_all_url_.php")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	// io.WriteString("./endphp", gotoStrarr)
	// dstFile.WriteString(string(gotoStrarr))
	for _, value := range gotoStrarr {
		dstFile.WriteString(value + "\n")
	}

}
