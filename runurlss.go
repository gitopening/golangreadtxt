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
	gotoStrarr := getAllUrlFromFile()

	// length := len(gotoStrarr)
	nReq := 20 // 10个并发请求
	var wg sync.WaitGroup
	wg.Add(nReq)
	for i := 0; i < 20; i++ {
		// goto_url := gotoStrarr[i]
		for _, value := range gotoStrarr {
			go func() {
				defer wg.Done()
				reqHTTP(value)
			}()
		}
	}
	wg.Wait()

	// reqHTTP("http://192.168.37.144:40185/www.01fy.cn")

}

func reqHTTP(url string) {
	// 简单的http Get
	resp, err := http.Get(url)
	if err != nil {
		// log.Printf("-----Http request error: %s\n", err, "-----Http request error: %s\n", url)
		log.Printf("-----Http request error:", err, "----", url)
		return
	}
	// data := make([]byte, resp.ContentLength)
	// resp.Body.Read(data)
	// log.Printf("Http response: %s\n", data)

	if resp.StatusCode == http.StatusOK {
		//如果获取状态不为 200,输出状态程序结束
		log.Println("run is ok-  " + url)
	}
}
