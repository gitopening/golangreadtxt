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
	// var list []string
	// urls := make([]urs[readtxtfiles()]string ,0)
	// fmt.Print(readtxtfiles(), nil)
	readtxtfiles()

	// var url string = "http://www.nowamagic.net/librarys/veda/all/"
	// for i := 1; i <= len(names); i++ {
	// 	waitgroup.Add(1) //计数器+1 可以认为是队列+1
	// 	// go reslove(names[i], i)

	// 	h, err := http.Get(names[i])
	// 	if h.StatusCode == http.StatusOK { //如果获取状态不为 200,输出状态程序结束
	// 		log.Println("is ok-  " + names[i])
	// 	}
	// 	if err != nil {
	// 		log.Println("is err-  " + names[i])
	// 	}

	// }
	// waitgroup.Wait() //进行阻塞等待 如果 队列不跑完 一直不终止

}

// func reslove(url string, page int) {
// 	p := strconv.Itoa(page)
// 	url += p
// 	defer waitgroup.Done() //如果跑完就进行 队列-1
// 	log.Println("start " + url)
// 	h, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 		return
// 	}
// 	if h.StatusCode != http.StatusOK { //如果获取状态不为 200,输出状态程序结束
// 		panic(err)
// 		return
// 	}
// 	defer h.Body.Close()
// 	buf := make([]byte, 1024) //创建一个字节数组 长度为 1024
// 	file_open, err := os.OpenFile("./html/"+p+".html", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
// 	if err != nil {
// 		panic(err)
// 		return
// 	}
// 	defer func() {
// 		time.Sleep(time.Duration(1 * 1e9))
// 		file_open.Sync()
// 		file_open.Close()
// 	}()
// 	for { //无限循环,读取网页数据
// 		num, _ := h.Body.Read(buf)
// 		//如果获取数量为0，说明已经取到头了
// 		if num == 0 {
// 			break
// 		}
// 		file_open.WriteString(string(buf[:num]))
// 	}
// 	log.Println("end  " + url)
// }

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

	// kindex := 1

	for {
		m1, err1 := h.ReadString('\n')

		// fmt.Printf("最终网址", m2)
		// log.Println("最终网址  " + m2)
		// if err1 == io.EOF && err2 == io.EOF {
		// 	return
		// }
		// for {
		// 	m2, err2 := h2.ReadString('\n')
		// 	m3 := strings.Replace(m2, "***", m1, -1)
		// 	names = append(names, m3)
		// 	fmt.Printf("最终网址", m3)
		if err1 == io.EOF {
			return
		}
		// log.Println("end stirng is ok-  " + m1)
		// break
		// log.Println("is ok-  " + m1)
		for {
			m2, err2 := h2.ReadString('\n')
			if err2 == io.EOF {
				return
			}

			m3 := strings.Replace(m2, "***", m1, 1)
			m3 = strings.Replace(m3, "\n", "", -1)
			m3 = strings.Replace(m3, " ", "", -1)
			// kindex++

			log.Println("end stirng is ok-  " + m3)
			// log.Fatalf("end stirng is ok-  %d ", kindex)
			// break
			// 这里是分界线
			// h, err := http.Get(m3)
			// // mhttpclient(m3)
			// if err != nil {
			// 	log.Println("--------prog is err-  " + m3)
			// log.Fatalf("--------prog is err- %d --%s ", h.StatusCode, m3)
			// log.Fatalf("--end is ok- %d --%s ", kindex, m3)
			// 	// log.Fatalf("--------prog is err- --%s ", m3)
			// 	continue
			// }

			// if h.StatusCode == http.StatusOK { //如果获取状态不为 200,输出状态程序结束
			// 	log.Println("is ok-  " + m3)
			// }
			// 这里是分界线
			// log.Println("--+++-prog is err-  " + m3)
			// client := http.Client{}

			// request, err := http.NewRequest("GET", m3, nil)
			// if err != nil {
			// 	log.Fatalf("http.NewRequest error: %s", err)
			// 	continue
			// }
			// // ----- modify start --------
			// // q := request.URL.Query()
			// // q.Add("keyword", "完美世界")
			// // request.URL.RawQuery = q.Encode()
			// // ----- modify end --------
			// request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
			// request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			// request.Header.Set("Connection", "keep-alive")

			// resp, err := client.Do(request)
			// if err != nil {
			// 	log.Println("--------prog is err-  " + m3)
			// 	fmt.Println(err)
			// 	continue
			// }
			// defer resp.Body.Close()
			// if resp.StatusCode != 200 {
			// 	log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
			// }

			// log.Println("is ok-  " + m3)

		}

		// }

		// fmt.Print(m1)

	}
	// return names

}

// func mhttpclient(url string) {
// 	// url := "https://m.999xs.com/search.php?keyword=完美世界"
// 	client := http.Client{}

// 	request, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatalf("http.NewRequest error: %s", err)
// 	}
// 	// ----- modify start --------
// 	// q := request.URL.Query()
// 	// q.Add("keyword", "完美世界")
// 	// request.URL.RawQuery = q.Encode()
// 	// ----- modify end --------
// 	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
// 	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	request.Header.Set("Connection", "keep-alive")

// 	resp, err := client.Do(request)
// 	if err != nil {
// 		fmt.Println(err)
// 		log.Println("--------prog is err-  " + url)
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != 200 {
// 		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
// 	}
// 	log.Println("--------prog is err-  " + url)

// }
