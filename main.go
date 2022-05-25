package main

import (
	"fmt"

	"main/util"
)

func main() {
	val, urls := util.HttpUrlParser("/a/b")
	fmt.Printf("%#v,%#v \n", val, urls)
	val, urls = util.HttpUrlParser("/a/b/c")
	fmt.Printf("%#v,%#v \n", val, urls)
	val, urls = util.HttpUrlParser("/a/12")
	fmt.Printf("%#v,%#v \n", val, urls)
	val, urls = util.HttpUrlParser("/a/bbbbbb")
	fmt.Printf("%#v,%#v \n", val, urls)

	//错误用例:
	val, urls = util.HttpUrlParser("bb/a/bbbbbb")
	fmt.Printf("%#v,%#v \n", val, urls)

	val, urls = util.HttpUrlParser("bb/a/bbbbbb/n/n/n/jj")
	fmt.Printf("%#v,%#v \n", val, urls)
}
