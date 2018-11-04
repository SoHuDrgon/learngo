package main

import (
	"fmt"
	"github.com/guogang1990/learngo/retriever/mock"
	"github.com/guogang1990/learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string] string) string
}
func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
	map[string]string {
		"name": "docker",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}
const url = "http://www.baidu.com"
func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents": "another faked baidu.com.",
	})
	return s.Get(url)
}
func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake baidu.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	inspect(r)
	//Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever.")
	}
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
	//fmt.Println(download(r))
	fmt.Println("Try a session.")
	s := retriever
	fmt.Println(session(&s))
	
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v .UserAgent)
	}
}