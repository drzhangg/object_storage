package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//如果是get请求，跳转到上传页面
	case "GET":
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	case "POST":
		//接收文件流并存储文件到本地
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("failed to get data, err:%s\n", err.Error())
			return
		}
		defer file.Close()



	}

}
