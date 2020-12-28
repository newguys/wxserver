package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
)

func main() {
	//绑定路由
	http.HandleFunc("/", checkout)
	//启动监听
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
	}
}

func checkout(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Printf("url解析失败,err:%s", err.Error())
		return
	}
	var token string = "Evildoer"
	signature := request.FormValue("signature")
	timestamp := request.FormValue("timestamp")
	nonce := request.FormValue("nonce")
	echostr := request.FormValue("echostr")

	//token timestamp nonce 排序
	var tempArray = []string{token, timestamp, nonce}
	sort.Strings(tempArray)

	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}

	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	fmt.Printf("resquest:signature%s,timestamp:%s,nonce:%s,echostr:%s", signature, timestamp, nonce, echostr)
	//获得加密后的字符串可与signature进行对比
	if sha1String == signature {
		_, err := response.Write([]byte(echostr))
		if err != nil {
			fmt.Printf("响应失败，err:%s", err.Error())
			return
		}
	} else {
		fmt.Println("验证失败")
	}

}
