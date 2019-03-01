package main

type Todo struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Title    string `json:"Tittle"`
	Subtitle string `json:"Subtitle"`
	Other    string `json:"Other"`
	Desc     string `json:"Desc"`
	Year     string `json:"Year"`
	Area     string `json:"Area"`
	Tag      string `json:"Tag"`
	Star     string `json:"Star"`
	Comment  string `json:"Comment"`
	Quote    string `json:"Quote"`
}
type Todos []Todo