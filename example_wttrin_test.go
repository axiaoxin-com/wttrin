package wttrin_test

import (
	"github/axiaoxin-com/wttrin"
	"io/ioutil"
	"log"
)

// 原始请求http://wttr.in的示例
func ExampleWttrIn() {
	body, err := wttrin.WttrIn("成都?0ATp")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
	}
	defer body.Close()
	log.Println(string(content))
}
