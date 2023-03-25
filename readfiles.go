package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	fmt.Println("getOldStrArr_tenurl文件中的原始内容长度：", len(oldstrarr))
	return oldstrarr

}

func getUrlStrArr() []string {
	//一次性读取文件
	read2, err := ioutil.ReadFile("urlallgo.php")
	if err != nil {
		fmt.Printf("文件打开失败", err.Error())
		return nil
	}
	// fmt.Printf(string(read))
	// fmt.Println("test文件中的原始内容：", read)

	gostrarr := strings.Split(string(read2), "\n")
	fmt.Println("urlallgo文件中的原始内容长度：", len(gostrarr))
	return gostrarr

}

func getAllUrlStrArr() []string {
	var urlsInRead []string
	oldstrarrs := getOldStrArr()
	gotstrarrs := getUrlStrArr()
	for index, _ := range gotstrarrs {
		// fmt.Printf("index=%d, value=%c\n", index, gotstrarrs[index])
		m1 := strings.Replace(gotstrarrs[index], "\n", "", -1)
		m2 := strings.Replace(m1, " ", "", -1)
		for index2, _ := range oldstrarrs {
			// fmt.Printf("index=%d, value=%c\n", index, oldstrarrs[index2])
			m21 := strings.Replace(oldstrarrs[index2], "\n", "", -1)
			m22 := strings.Replace(m21, " ", "", -1)
			m3 := strings.Replace(m2, "***", m22, 1)
			urlsInRead = append(urlsInRead, m3) //添加到数组中

		}
	}
	return urlsInRead
}

func getAllUrlFromFile() []string {

	//一次性读取文件
	read2, err := ioutil.ReadFile("end_all_url_.php")
	if err != nil {
		fmt.Printf("文件打开失败", err.Error())
		return nil
	}
	// fmt.Printf(string(read))
	// fmt.Println("test文件中的原始内容：", read)

	allUrls := strings.Split(string(read2), "\n")
	fmt.Println("end_all_url_文件中的原始内容长度：", len(allUrls))
	return allUrls

}

func readToOnefile() {
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
