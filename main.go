package main

import (
	"fmt"
	"net/http"
	"object_storage/handler"
)

func main() {
	//静态资源处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)

	// 监听端口
	fmt.Println("上传服务正在启动, 监听端口:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}

	//http.HandleFunc("/objects/",object.Handler)
	//log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil))
}
