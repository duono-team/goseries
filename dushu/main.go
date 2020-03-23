package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	str := ""
	if len(args) == 0 {
		str = "340282366920938463463374607431768211456"
	} else {
		str = args[0]
	}
	fmt.Println("数字：", str)
	count := 0
	yu := 0
	for i, _ := range str {
		yu = (i + 1) % 4
		if yu == 0 {
			count++
		}
	}
	// dotSet := "个十百千万亿兆京垓秭穰沟涧正载极恒河沙阿僧只那由他不可思议无量大数"
	doku := ""
	countCopy := count
	yuCopy := yu
	index := 0
	for i := 0; i <= count; i++ {
		switch yuCopy {
		case 0:
			doku += toZh(str[index]) + "千"
			index++
			doku += toZh(str[index]) + "百"
			index++
			doku += toZh(str[index]) + "十"
			index++
			if countCopy == 0 {
				doku += toZh(str[index])
			} else {
				doku += toZh(str[index]) + appdx(countCopy)
				countCopy--
			}
			index++
			yuCopy = 0
		case 1:
			if countCopy == 0 {
				doku += toZh(str[index])
			} else {
				doku += toZh(str[index]) + appdx(countCopy)
				countCopy--
			}
			index++
			yuCopy = 0
		case 2:
			doku += toZh(str[index]) + "十"
			index++
			if countCopy == 0 {
				doku += toZh(str[index])
			} else {
				doku += toZh(str[index]) + appdx(countCopy)
				countCopy--
			}
			index++
			yuCopy = 0
		case 3:
			doku += toZh(str[index]) + "百"
			index++
			doku += toZh(str[index]) + "十"
			index++
			if countCopy == 0 {
				doku += toZh(str[index])
			} else {
				doku += toZh(str[index]) + appdx(countCopy)
				countCopy--
			}
			index++
			yuCopy = 0
		}
	}

	fmt.Printf("这个数字一共有%v位，有%v个四位余%v位\n读作：%v\n", len(str), count, yu, doku)
}

func toZh(s byte) string {
	switch s {
	case "0"[0]:
		return "零"
	case "1"[0]:
		return "一"
	case "2"[0]:
		return "二"
	case "3"[0]:
		return "三"
	case "4"[0]:
		return "四"
	case "5"[0]:
		return "五"
	case "6"[0]:
		return "六"
	case "7"[0]:
		return "七"
	case "8"[0]:
		return "八"
	case "9"[0]:
		return "九"
	default:
		return "E"
	}
}

func appdx(i int) string {
	switch i {
	case 1:
		return "万"
	case 2:
		return "亿"
	case 3:
		return "兆"
	case 4:
		return "京"
	case 5:
		return "垓"
	case 6:
		return "秭"
	case 7:
		return "穰"
	case 8:
		return "沟"
	case 9:
		return "涧"
	case 10:
		return "正"
	case 11:
		return "极"
	default:
		return "E"
	}
}
