package main

import (
	"encoding/json"
	"fmt"
)

// capital first letter : exporting it
// small first letter : it is private no export
type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // i don't want password in my JSON
	Tags     []string `json:"tags,omitempty"` // if it's nil don't show it
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS", 299, "udemy", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "udemy", "ash123", []string{"web-dev", "js"}},
		{"Angular", 299, "udemy", "bcd123", nil},
	}

	//package  this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "RectJS",
		"price": 199,
		"website": "udemy",
		"tags": ["web-dev","js"]
	}
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	// without struct that's why used interface
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("%v = %v : %T\n", k, v, v)
	}
}
