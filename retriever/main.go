package main

import (
	"LearnGolang/retriever/mock"
	"LearnGolang/retriever/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake website"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}

	inspect(r)

	//type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)
}
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	fmt.Println("Type switch:")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)

	}
}

//fmt.Println(download(r))
