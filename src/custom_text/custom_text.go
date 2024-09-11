package main

import (
	"encoding/json"
	"fmt"
)

const myUrl string = "https://gobyexample.com/"

type course struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Platform string  `json:"website"`
}

func main() {

	fmt.Println("wel come rakesh")
	//enodeMyJson()
	decodeMyJson()
}

func enodeMyJson() {

	myCourses := []course{{"Rakesh", 12.3, "www.google.com"}, {"Ravi", 12.99, "www.abc.com"}, {"Ramya", 12887.3, "www.hero.com"}}
	myCourses2 := []course{{"kulla", 12.3, "www.google.com"}, {"Ravi", 12.99, "www.abc.com"}, {"yyy", 12887.3, "www.hero.com"}}
	myCourses = append(myCourses, myCourses2...)
	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}

func decodeMyJson() {
	jsonFromWeb := []byte(` 
        {
                "name": "yyy",
                "price": 12887.3,
                "website": "www.hero.com"
        }`)
	var lcoCourse course
	checkValid := json.Valid(jsonFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is: %T\n", k, v, v)
	}
}
