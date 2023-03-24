package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
var names []string

func main() {
	oldStrarr := getOldStrArr()
	// fmt.Printf(string(oldStrarr))

	fmt.Print("error:%s", oldStrarr)

	readtxtfiles()

	// for i := 1; i <= len(names); i++ {
	// 	waitgroup.Add(1)
	// 	//计数器+1 可以认为是队列+1
	// 	h, err := http.Get(names[i])

	// 	if err != nil {
	// 		//panic(err)
	// 		log.Println("is err-  " + names[i])
	// 		continue

	// 	}

	// 	if h.StatusCode == http.StatusOK {
	// 		//如果获取状态不为 200,输出状态程序结束
	// 		log.Println("run is ok-  " + names[i])
	// 	}

	// }
	// waitgroup.Wait() //进行阻塞等待 如果 队列不跑完 一直不终止

}

func readtxtfiles() {

	i, v := os.Open("tenurl.php")
	i2, v2 := os.Open("urlallgo.php")
	if v2 != nil {
		// fmt.Printf("文件打开失败" ,v2)
		log.Println("文件打开失败", v2)
		return
	}

	if v != nil {
		// fmt.Printf("文件打开失败" + v)
		log.Println("文件打开失败 ", v)
		return
	}
	defer i.Close()
	h := bufio.NewReader(i)
	defer i2.Close()
	h2 := bufio.NewReader(i2)
	for {
		m1, err1 := h.ReadString('\n')
		if err1 == io.EOF {
			return
		}
		for {
			m2, err2 := h2.ReadString('\n')
			if err2 == io.EOF {
				return
			}

			m3 := strings.Replace(m2, "***", m1, 1)
			m3 = strings.Replace(m3, "\n", "", -1)
			m3 = strings.Replace(m3, " ", "", -1)
			// kindex++

			// log.Println("end stirng is ok-  " + m3)
			names = append(names, m3) //添加到数组中
		}
	}

	fmt.Println("原数组内容为:", names)

}
