// Package wttrin 封装http://wttr.in
//
// 不同形式的天气预报URL格式：
//
//   ASCII： http://wttr.in/<location>?<query>
//   图片： http://wttr.in/<location>_<query>.png
//   单行文本： http://wttr.in/<location>?format=<format>&<query>
//
// location 是需要查询天气预报的地点，可以使用英文字符串，也可以直接传中文
//
// query 是控制天气预报内容的参数，所有控制参数直接拼接在一起即可，如0_A_T_p表示对输出内容做对应的4种处理
//
// 不加query参数，默认返回html内容，包含全部信息，从上到下为：预报的地点；当前天气信息；今天、明天、后天的表格天气信息；底部详细地点信息；底部wttrin作者项目信息
//
// 以下信息是阅读wttr.in源码写的，其中有点参数看起来还在实现中，因此仅供参考。
//
// query 控制参数解释:
//
//   A: 返回带颜色的ASCII内容，不加返回的是HTML
//   n: 表格天气信息只返回中午和夜间，不加返回的是早上、中午、傍晚、夜间的信息
//   m: 温度展示位摄氏度
//   u: 温度展示为华氏度，风速等其他值和单位也有对应变化
//   M: 风速展示为m/s 不加展示为km/h
//   I: 反转html或ASCII中的颜色
//   t: 设置png图片透明值为150
//   transparency: 指定png图片透明值，transparency=123
//   T: 返回的ASCII内容不带颜色
//   p: 对内容设置一定的padding
//   0: 返回当前天气信息
//   1: 返回当前+今天的天气信息
//   2: 返回当前+今天+明天的天气信息
//   3: 返回当前+今天+明天+后天的天气信息
//   q: 不显示底部详细地址和顶部地址信息的前缀信息
//   Q: 不显示顶部和底部的地点信息
//   F: 不显示底部作者项目信息
//   lang: 天气语言翻译 "az", "bg", "bs", "cy", "cs", "eo", "es", "fi", "ga", "hi", "hr", "hy", "is", "ja", "jv", "ka", "kk", "ko", "ky", "lt", "lv", "mk", "ml", "nl", "fy", "nn", "pt", "pt-br", "sk", "sl", "sr", "sr-lat", "sv", "sw",  "te", "uz", "zh", "zu", "he"
//
// format 参数解释
//
//   %l: 地点
//   %c: 天气图标
//   %C: 天气文字
//   %t: 温度
//   %w: 风速
//   %m: 月相图标 0-新月 1-眉月 2-上弦月 3-盈凸月 4-满月 5-亏凸月 6-下弦月 7-残月
//   %M: 新月后第几天
//   %h: 湿度
//   %p: 降水量
//   %o: 降水几率
//   %P: 气压
//   %s: 日落时间
//   1: 预定义格式(⛅️ +15°C)
//   2: 预定义格式(⛅️ 🌡️+15°C 🌬️↗11 km/h)
//   3: 预定义格式(成都: ⛅️ +15°C)
//   4: 预定义格式(成都: ⛅️ 🌡️+15°C 🌬️↗11 km/h)
package wttrin

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/axiaoxin-com/logging"
)

// WttrIn 获取 GET 请求 http://wttr.in 的返回 Body
func WttrIn(locationQuery string) (io.ReadCloser, error) {
	wttrinURL := "http://wttr.in/" + locationQuery
	logging.Debugs(nil, "wttrin request url:", wttrinURL)
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(wttrinURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("wttrin response status error:" + resp.Status)
	}
	return resp.Body, nil
}

// Line 单行天气信息
func Line(lang, location, format string, q ...string) (string, error) {
	if lang == "" {
		lang = "zh"
	}
	if location == "" {
		location = "成都"
	}
	if format == "" {
		format = "当前%l:\n天气%c %C\n温度🌡️ %t\n风速🌬️ %w\n湿度💦 %h\n气压🧭 %P\n降水☔️ %p\n月相🌑 +%M%m"
	}
	query := "m_M"
	if len(q) > 0 {
		query = strings.Join(q, "_")
	}

	locationQuery := fmt.Sprintf("%s?lang=%s&format=%s&%s", location, lang, url.QueryEscape(format), query)
	resp, err := WttrIn(locationQuery)
	if err != nil {
		return "", err
	}
	result, err := ioutil.ReadAll(resp)
	if err != nil {
		return "", err
	}
	line := string(result)
	if isServiceUnavailable(line) {
		return "", errors.New(line)
	}
	if strings.Contains(line, "Unknown location; please try") {
		return "", errors.New("wttrin Line get location failed")
	}
	return line, nil
}

// ASCII 图形天气信息
func ASCII(lang, location string, q ...string) (string, error) {
	if lang == "" {
		lang = "zh"
	}
	if location == "" {
		location = "成都"
	}
	query := "0_A_T_F_m_M_p"
	if len(q) > 0 {
		query = strings.Join(q, "_")
	}
	locationQuery := fmt.Sprintf("%s?lang=%s&%s", location, lang, query)
	resp, err := WttrIn(locationQuery)
	if err != nil {
		return "", err
	}
	result, err := ioutil.ReadAll(resp)
	if err != nil {
		return "", err
	}
	ascii := string(result)
	if isServiceUnavailable(ascii) {
		return "", errors.New(ascii)
	}
	return ascii, nil
}

// Image 图片天气信息
func Image(lang, location string, q ...string) (io.ReadCloser, error) {
	if lang == "" {
		lang = "zh"
	}
	if location == "" {
		location = "成都"
	}
	query := "0_m_p_q"
	if len(q) > 0 {
		query = strings.Join(q, "_")
	}
	locationQuery := fmt.Sprintf("%s_lang=%s_%s.png", location, lang, query)
	return WttrIn(locationQuery)
}

func isServiceUnavailable(respText string) bool {
	str1 := "======================================================================================"
	str2 := "https://twitter.com/igor_chubin"
	return strings.Contains(respText, str1) && strings.Contains(respText, str2)
}
