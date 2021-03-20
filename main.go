package main

import (
	"fmt"
)

func main() {
	// var hs1 thongtin = &study{"cntt1","to kim manh",18}
	// var tc1 thongtin = &teacher{"cntt1","su van luong","giai tich",30}
	hs1 := study{
		"cntt1",
		"to kim manh",
		18,
	}
	hs1.Show()
}

type study struct {
	class string
	name  string
	age   int
}

type teacher struct {
	class   string
	name    string
	subject string
	age     int
}

func (hs *study) Show() {
	fmt.Println(
		hs.class,
		hs.name,
	)
}

func (tc *teacher) show() {
	// fmt.Println(
	// 	tc.class,
	// 	tc.name,
	// 	tc.subject,
	// )
	tc.name.Show()
}

// func (hs *study) old() {
// 	fmt.Println(hs.age)
// }

// func (tc *teacher) old() {
// 	fmt.Println(
// 		"tuoi cua toi la:",
// 		tc.age,
// 	)
// }

// type thongtin interface{
// 	show()
// 	old()
// }
