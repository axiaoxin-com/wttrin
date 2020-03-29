package wttrin

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestWttrIn(t *testing.T) {
	body, err := WttrIn("成都?0ATp")
	if err != nil {
		t.Error(err)
	}
	content, err := ioutil.ReadAll(body)
	if err != nil {
		t.Error(err)
	}
	defer body.Close()
	t.Log(string(content))
}

func TestLine(t *testing.T) {
	result, err := Line("zh", "成都", "地点%l 天气图标%c 天气文字%C 温度%t 风速%w 月相图标%m 新月后第几天%M 湿度%h 降水量%p 降水几率%o 气压%P 日落时间%s")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

	result, err = Line("zh", "成都", "")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestASCII(t *testing.T) {
	result, err := ASCII("zh", "成都")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)

	result, err = ASCII("zh", "成都", "0", "A", "Q")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestImage(t *testing.T) {
	result, err := Image("zh", "成都")
	if err != nil {
		t.Error(err)
	}
	f, err := os.Create("./wttrin_noquery.png")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	if _, err := io.Copy(f, result); err != nil {
		t.Error(err)
	}

	result, err = Image("zh", "成都", "A", "p", "F")
	if err != nil {
		t.Error(err)
	}
	f1, err := os.Create("./wttrin_query.png")
	if err != nil {
		t.Error(err)
	}
	defer f1.Close()
	if _, err := io.Copy(f1, result); err != nil {
		t.Error(err)
	}
}
