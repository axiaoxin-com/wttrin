package wttrin_test

import (
	"github/axiaoxin-com/wttrin"
	"log"
)

// 单行文字的天气预报示例：
func ExampleLine() {
	// 自定义格式
	result, err := wttrin.Line("zh", "成都", "地点%l 天气图标%c 天气文字%C 温度%t 风速%w 月相图标%m 新月后第几天%M 湿度%h 降水量%p 降水几率%o 气压%P 日落时间%s")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)

	// 默认格式
	result, err = wttrin.Line("zh", "成都", "")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
}
