package getter

//go:generate boilergen

type MyModel struct {
	aaa string
	bbb int
	ccc float64
	ddd struct {
		eee string
		fff []int
		ggg map[string]int64
	}
}
