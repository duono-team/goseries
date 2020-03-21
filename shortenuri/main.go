package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"net/url"
	"strings"
	"flag"
)

type ResUrlStruct struct {
	Result_url string
	Error string
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Please enter the uri that needs to be shortened")
		os.Exit(1)
	}

	strslice := make([]string, 1)
	strslice[0] = args[0]
	var d url.Values = make(map[string][]string)
	d["url"] = strslice
	res, err := http.PostForm("https://cleanuri.com/api/v1/shorten", d)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var s ResUrlStruct
	json.Unmarshal(body, &s)
	if s.Error != "" {
		fmt.Println(s.Error)
		return
	}
	fmt.Println(strings.Trim(s.Result_url, "\\"))
}