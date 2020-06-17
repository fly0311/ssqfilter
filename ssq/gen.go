package ssq

import "sync"

type SsqData struct {
	Data map[int][]int
	Len  int
}

func (d SsqData) Init() {
	data := make(map[int][]int)
	wg := new(sync.WaitGroup)
	for i := 0; i < 33-6-1; i++ {

	}
}


func (d SsqData) GenData(start, end int) []int{
	for i:=start;

}
