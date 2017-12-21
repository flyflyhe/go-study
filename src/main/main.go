package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"regexp"
)

type Handle struct {
}

func (e Handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	text := "<a href='/test' ></a><a href='/home'></a>"	
	reg := regexp.MustCompile(`<a.*?href=['"]*([^'"\s<>]+)['"]*.*?>`)
	match := reg.FindAllStringSubmatch(text, -1)
	for _,v := range match {
		fmt.Println(v[1])
	}
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	handle := &Handle{}
	s := &http.Server{
		Addr:           ":9090",
		Handler:        handle,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
