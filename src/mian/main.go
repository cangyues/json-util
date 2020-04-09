package main

import (
	"fmt"
	"qson"
)

func main() {

	q := qson.ParseJSONArray("[{\"name\":\"cangyue\"}]")

	fmt.Println(q.ToString())
}
