package parser

import (
	"learngo/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents,err :=fetcher.Fetch(
		"http://album.zhenai.com/u")

	if err!=nil{
		panic(err)
	}
	result:= ParseCityList(contents)
	//fmt.Printf("%s\n",contents)
	//verify result
	const resultSize  =470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string {
		"阿坝","阿克苏","阿拉善盟",
	}

	if len(result.Requests) != resultSize{
		t.Errorf("result should have %d"+
			"resquest; but had %d",
			resultSize, len(result.Requests))
	}

	for i,url := range expectedUrls{
		if result.Requests[i].Url != url{
			t.Errorf("expected url #%d: %s; but"+
				"was %s",
				i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize{
		t.Errorf("result should have %d"+
			"resquest; but had %d",
			resultSize, len(result.Items))
	}
	for i,city := range expectedCities{
		if result.Items[i].(string) != city{
			t.Errorf("expected city #%d: %s; but"+
				"was %s",
				i, city, result.Items[i].(string))
		}
	}
}
