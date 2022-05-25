package util

import (
	"strconv"
	"strings"
)

//写一个HTTP URL Parser.需要同时实现以下匹配规则:

// /a/<b:int> 指的是,如果一个URL是/a/1,那么命中该规则,返回b的值为int类型1
// /a/<b:string> 指的是,如果一个URL是/a/bbbb,那么命中该规则,返回b的值为string类型bbbbb
// /a/* 指的是,如果一个URL是/a/b或者/a/c,那么命中该规则,返回这个URL
// /a/b/c 指的是,当且仅当如果一个URL是/a/b/c,那么命中该规则,返回这个URL

//须知, URL里边的/a/b是形参,a和b的长度不是定长的
func HttpUrlParser(url string) (val interface{}, urls []string) {
	values := strings.Split(url, "/")
	//长度不符合规范
	if values[0] != "" {
		return nil, []string{}
	}
	if len(values) > 4 {
		return nil, []string{}
	}
	//判断参数类型
	i, err := strconv.ParseInt(values[2], 10, 64)
	if err != nil {
		val = values[2]
	} else {
		val = i
	}

	if len(values) == 3 {
		urls = append(urls, url)
	}
	if len(values) == 4 {
		urls = append(urls, url)
	}
	return
}
