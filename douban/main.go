// 爬取豆瓣电影 TOP250
package main

import (
	"learngo/douban/parse"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strings"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

// 开始爬取
func Start() {
	var movies []parse.DoubanMovie

	pages := parse.GetPages(BaseUrl)
	for _, page := range pages {
		doc, err := goquery.NewDocument(strings.Join([]string{BaseUrl, page.Url}, ""))
		if err != nil {
			log.Println(err)
		}

		movies = append(movies, parse.ParseMovies(doc)...)

		filePtr, err :=os.Create("Douban.json")
		if err != nil{
			fmt.Println("Create file failed",err.Error())
			return
		}
		defer filePtr.Close()
		encoder := json.NewEncoder(filePtr)

		err = encoder.Encode(movies)
		if err != nil {
			fmt.Println("Encoder failed", err.Error())

		} else {
			fmt.Println("Encoder success")
		}
	}

}

//func writeJsonFile(){
//	var movies []Parse.DoubanMovie
//	filePtr, err :=os.Create("Douban.json")
//	if err != nil{
//		fmt.Println("Create file failed",err.Error())
//		return
//	}
//	defer filePtr.Close()
//	encoder := json.NewEncoder(filePtr)
//
//	err = encoder.Encode(movies)
//	if err != nil {
//		fmt.Println("Encoder failed", err.Error())
//
//	} else {
//		fmt.Println("Encoder success")
//	}
//
//}


func main() {
	Start()
	readFile()

}
func readFile() {

	filePtr, err := os.Open("Douban.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	var movies []parse.DoubanMovie

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&movies)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())

	} else {
		fmt.Println("Decoder success")
		fmt.Println(movies)
	}
}
