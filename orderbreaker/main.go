package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := bufio.NewReader(os.Stdin)
	fmt.Println("请输入要打乱的文字：")
	str, err := r.ReadString(10)
	if err != nil {
		panic(err.Error())
	}
	if str == "" + string(13) + string(10) {
		fmt.Println("你什么都没输入哦~")
		return
	}
	time.Sleep(time.Millisecond * 10)
	str = strings.TrimRight(str, string(13) + string(10))
	strSlice := strings.Split(str, "")
	conv := make([]int, len(strSlice))
	for i:=0;;{
		k := rand.Intn(len(strSlice) + 1)
		i2 := 0
		for _, v := range conv {
			if v == k {
				i2++
			}
		}
		if i2 == 0 {
			conv[i] = k
			i++
		}
		if i > len(conv) - 1 {
			break
		}
	}
	newStr := make([]string, len(strSlice))
	for i, v := range conv {
		newStr[i] = strSlice[v-1]
	}
	str = strings.Join(newStr, "")
	fmt.Println("结果：")
	fmt.Println(str)
	fmt.Println("按回车以退出...")
	r.ReadString(10)
}