package parser

import (
	"LearnGolang/crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func Parsecity(
	contents []byte)engine.ParseResult {
		re:= regexp.MustCompile(cityRe)
		matches:=re.FindAllSubmatch(contents,-1)

		result := engine.ParseResult{}
		limit :=1
		for _,m := range matches{
			result.Items =append(
				result.Items, "User "+string(m[2]))
			result.Requests =append(
				result.Requests, engine.Request{
					Url:    string(m[1]),
					ParserFunc: func(c []byte) engine.ParseResult {
						return ParseProfile(c, string(m[2]))
					},
				})
			limit --
			if limit ==0{
				break
			}
		}
		return result
	
}
