package ssq

import "fmt"

type SsqData struct {
	Data       map[int]bool //单式大底
	Len        int          //单式注数
	Red        []int        //红球可选
	Blue       []int        //红球可选
	RedDan     []int        //红球胆码
	BlueDan    []int        //蓝球胆码
	RedKilled  []int        //红球杀号
	BlueKilled []int        //蓝球杀号
}

func NewSsq() *SsqData {
	return &SsqData{
		Data:       make(map[int]bool),
		Red:        make([]int, 33),
		Blue:       make([]int, 16),
		RedDan:     make([]int, 0),
		BlueDan:    make([]int, 0),
		RedKilled:  make([]int, 0),
		BlueKilled: make([]int, 0),
	}
}

//初始化数据
func (d SsqData) Init() {
	//初始化大底
	rs := make(map[int]bool)
	d := make([]int, 33)
	for i := 0; i < 6; i++ {
		d[i] = 1
	}

	for {
		cg := false
		for j := 0; j < 32; j++ {
			if d[j] == 1 && d[j+1] == 0 {
				d[j], d[j+1] = d[j+1], d[j]

				s := 0
				e := j - 1

				ones := 0
				for a := s; a <= e; a++ {
					if d[a] == 1 {
						ones++
					}
				}

				for b := 0; b < ones; b++ {
					d[b] = 1
				}
				for c := ones; c <= e; c++ {
					d[c] = 0
				}

				var this int = 0
				for index := 0; index < 33; index++ {
					if d[index] == 1 {
						this += 1 << (33 - index - 1)
					}
				}
				rs[this] = true
				cg = true
				break
			}
		}
		if cg == false {
			break
		}
	}
	d.Data = rs

	//初始化红蓝球
	for i := 1; i <= 33; i++ {
		d.Red[i-1] = i
	}
	for i := 1; i <= 16; i++ {
		d.Blue[i-1] = i
	}
	return
}

func (d SsqData) Print() {
	fmt.Sprintf("当前大底：%v 注\n", d.Len)
	fmt.Println("红球可选：", d.Red)
	fmt.Println("红球胆码：", d.RedDan)
	fmt.Println("红球杀号：", d.RedKilled)

	fmt.Println("蓝球可选：", d.Blue)
	fmt.Println("蓝球胆码：", d.BlueDan)
	fmt.Println("蓝球杀号：", d.BlueKilled)
	return
}
