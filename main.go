package main

import "net/http"

func main() {
	//静态资源处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//http.HandleFunc("/file/upload",handler)
}
