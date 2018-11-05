package main

import (
	"github.com/guogang1990/learngo/errhanding/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

//统一出错处理
type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(
	handler appHandler) func(
	w http.ResponseWriter, r *http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		//定义自我保护
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s",
				err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
				//http.Error(
				//	writer,
				//	http.StatusText(http.StatusNotFound),
				//	http.StatusNotFound)
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

//定义希望给用户看到的错误
type userError interface {
	error
	Message() string
}

func main() {
	//http.HandleFunc("/list/",
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
