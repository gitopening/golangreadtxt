package main

import (
	"fmt"
	"log"
	"net/http"
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

	fmt.Println("组合数组最终--内容长度:", len(gotoStrarr))

	for i := 1; i <= len(gotoStrarr); i++ {
		waitgroup.Add(1)
		//计数器+1 可以认为是队列+1
		h, err := http.Get(gotoStrarr[i])

		if err != nil {
			//panic(err)
			// fmt.Printf("index=%d, value=%c\n", index, gotstrarrs[index])
			log.Println("is err-", i, "-value=  ", gotoStrarr[i])
			continue

		}

		if h.StatusCode == http.StatusOK {
			//如果获取状态不为 200,输出状态程序结束
			log.Println("run is ok- -", i, "-value=  ", gotoStrarr[i])
		}

	}
	waitgroup.Wait() //进行阻塞等待 如果 队列不跑完 一直不终止

}
