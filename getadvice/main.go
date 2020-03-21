package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Advices struct {
	Advice string
	Slip_id string
}

type AdviceStruct struct {
	Slip Advices
}

func main() {
	res, err := http.Get("https://api.adviceslip.com/advice")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var s AdviceStruct
	json.Unmarshal(body, &s)
	fmt.Println(s.Slip.Advice)
}