package wttrin_test

import (
	"github/axiaoxin-com/wttrin"
	"log"
)

// 返回ASCII图形的天气预报示例
func ExampleASCII() {
	// 默认样式控制
	result, err := wttrin.ASCII("zh", "成都")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	// 自定义样式控制
	result, err = wttrin.ASCII("zh", "成都", "0", "A", "Q")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
