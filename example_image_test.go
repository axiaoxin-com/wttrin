package wttrin_test

import (
	"io"
	"log"
	"os"

	"github.com/axiaoxin-com/wttrin"
)

// 返回天气预报图片的示例
func ExampleImage() {
	// 默认样式的图片
	result, err := wttrin.Image("zh", "成都")
	if err != nil {
		log.Fatal(err)
	}
	// 保存图片
	f, err := os.Create("./wttrin_noquery.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := io.Copy(f, result); err != nil {
		log.Fatal(err)
	}

	// 自定义图片样式
	result, err = wttrin.Image("zh", "成都", "A", "p", "F")
	if err != nil {
		log.Fatal(err)
	}
	// 保存图片
	f1, err := os.Create("./wttrin_query.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	if _, err := io.Copy(f1, result); err != nil {
		log.Fatal(err)
	}
}
