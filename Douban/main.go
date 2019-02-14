// 爬取豆瓣电影 TOP250
package main

import (
	"LearnGolang/Douban/Parse"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

var (
	BaseUrl = "https://movie.douban.com/top250"
)

// 开始爬取
func Start() {
	var movies []Parse.DoubanMovie

	pages := Parse.GetPages(BaseUrl)
	for _, page := range pages {
		doc, err := goquery.NewDocument(strings.Join([]string{BaseUrl, page.Url}, ""))
		if err != nil {
			log.Println(err)
		}

		movies = append(movies, Parse.ParseMovies(doc)...)
	}

}

func main() {
	Start()

}
