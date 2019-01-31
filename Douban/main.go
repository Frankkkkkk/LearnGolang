// 爬取豆瓣电影 TOP250
package main

import (
	"LearnGolang/Douban/parse"
	"github.com/PuerkitoBio/goquery"
	"log"
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
	}

}

func main() {
	Start()

}
