package time

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	//diff := 16
	//t1 := "2021-08-18 16:17:52"
	//t2 := "2021-08-19 12:00:00"
	t1 := time.Date(2021, 8, 9, 14, 0, 0, 100, time.Local)
	t2 := time.Date(2021, 8, 19, 12, 0, 0, 100, time.Local)
	diff := TimeSub(t2, t1)
	fmt.Println(diff)
	//t1 := time.Date(2018, 1, 9, 24, 0, 1, 100, time.Local)
	//t2 := time.Date(2018, 1, 9, 23, 59, 22, 100, time.Local)
	////
	//fmt.Println(TimeSub(t1, t2))
	//
	//t1 = time.Date(2018, 1, 10, 0, 0, 1, 100, time.UTC)
	//t2 = time.Date(2018, 1, 9, 23, 59, 22, 100, time.UTC)

	//fmt.Println(TimeSub(t1, t2))

	pm, _ := strconv.ParseFloat("0.25", 64)
	premium := pm * 100
	a := int(premium) * diff
	out := float64(a) / 100
	outStr := strconv.FormatFloat(out, 'f', -1, 64)
	fmt.Println(outStr)

	return

	var Order = make(map[string]interface{})

	Order["order_id"] = "20190707212318"

	Order["order_price"] = 21.3

	Goods := make([]map[string]interface{}, 2)

	Goods[0] = make(map[string]interface{})
	Goods[0]["goods_name"] = "手机"
	Goods[0]["goods_price"] = 23.1

	Goods[1] = make(map[string]interface{})
	Goods[1]["goods_name"] = "电脑"
	Goods[1]["goods_price"] = 123.1

	GoodsColor := make([]map[string]interface{}, 2)

	GoodsColor[0] = make(map[string]interface{})
	GoodsColor[0]["good_color"] = "红色"

	GoodsColor[1] = make(map[string]interface{})
	GoodsColor[1]["good_color"] = "蓝色"

	Goods[0]["goods_color"] = GoodsColor
	Goods[1]["goods_color"] = GoodsColor

	Order["good"] = Goods

	data, _ := json.Marshal(Order)

	fmt.Println(string(data))
}

func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int(t1.Sub(t2).Hours() / 24)
}
