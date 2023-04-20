package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	content, _ := getUrlContent("https://www.baidu.com")

	fmt.Println(content)
}

func getUrlContent(url string) (string, error) {

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	// 执行到该条语句时，函数和参数表达式得到计算，但直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行。
	// 不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
