package array

import (
	"fmt"
	"strconv"
	"testing"
)

func TestFloat(t *testing.T) {
	a := getTbHotGoodsOption()

	for _, v := range a {
		t.Log(v.ToString())
	}
}

type TbHotGoodsOption struct {
	MinMoney float64 // 最小货单价
	MaxMoney float64 // 最大货单价
	MinCount int     // 统计即将达标的最小值
	Count    int     // 达标最小值
}

func getTbHotGoodsOption() []*TbHotGoodsOption {
	hg := []*TbHotGoodsOption{}
	hg = append(hg, &TbHotGoodsOption{
		MinMoney: 10,
		MinCount: 1000,
		Count:    3000,
	})

	hg = append(hg, &TbHotGoodsOption{
		MinMoney: 9.9,
		MaxMoney: 20,
		MinCount: 500,
		Count:    3000,
	})

	hg = append(hg, &TbHotGoodsOption{
		MinMoney: 20,
		MaxMoney: 40,
		MinCount: 300,
		Count:    1000,
	})

	hg = append(hg, &TbHotGoodsOption{
		MinMoney: 40,
		MinCount: 100,
		Count:    500,
	})
	return hg
}

// ToString 转字符串  要求规则小数精度最多为一位
func (e *TbHotGoodsOption) ToString() string {
	var out string = "[%s,%s)->%d"
	out = fmt.Sprintf(out,
		func(m float64) string {
			mstr := fmt.Sprintf("%.f", m)
			mfloat, _ := strconv.ParseFloat(mstr, 0)
			if m != mfloat {
				return fmt.Sprintf("%.1f", m)
			}
			return fmt.Sprintf("%.f", m)
		}(e.MinMoney),
		func(m float64) string {
			if m <= 0 {
				return "++"
			}
			mstr := fmt.Sprintf("%.f", m)
			mfloat, _ := strconv.ParseFloat(mstr, 0)
			if m != mfloat {
				return fmt.Sprintf("%.1f", m)
			}
			return fmt.Sprintf("%.f", m)
		}(e.MaxMoney), e.Count)
	return out
}
