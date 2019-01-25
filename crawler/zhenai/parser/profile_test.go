package parser

import (
	"LearnGolang/crawler/fetcher"
	"LearnGolang/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents,err :=fetcher.Fetch(
		"http://www.zhenai.com/zhenghun")
	if err != nil{
		panic(err)
	}
	result := ParseProfile(contents, "安静的雪")

	if len(result.Items) != 1{
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Age: 34,
		Marriage: "离异",
	}
	if profile != expected{
		t.Errorf("expected %v: but was %v",1, len(result.Items))
	}
}
