package parser

import (
	"learngo/crawler/fetcher"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents,err :=fetcher.Fetch(
		"http://album.zhenai.com/u/")
	if err != nil{
		panic(err)
	}
	result := ParseProfile(contents, "不想")

	if len(result.Items) != 1{
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:"不想",
		Age: 37,
	}
	if profile != expected{
		t.Errorf("expected %v: but was %v", expected,profile)
	}
}
