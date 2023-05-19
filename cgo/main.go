package main

// #include "struct.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// func print(p uintptr) {
// 	pacingCtx := (*C.struct_Pacing)(unsafe.Pointer(p))
// 	req := pacingCtx.request_
//
// 	var cands *C.struct_Candidate = req.cands_
// 	//var counters *C.struct_Counter = req.counters_
// 	fmt.Printf("pacing request: context: %v\n", req.context_)
// 	candidates := (*[1 << 30]C.struct_Candidate)(unsafe.Pointer(cands))[:req.cands_len_:req.cands_len_]
// 	fmt.Printf("pacing request: candidate size: %d\n", req.cands_len_)
// 	fmt.Printf("len %d, %v\n", len(candidates), candidates[0])
// 	for _, cand := range candidates {
// 		fmt.Printf("candidate[%d]: %v\n", cand.id_, cand)
// 	}
// 	//fmt.Printf("counters: %d\n", counters.id_)
// 	// slice := (*[1 << 30]C.struct_Counter)(unsafe.Pointer(counters))[:req.cands_len_:req.cands_len_]
//
// 	var c counter.Counter = counter.New(unsafe.Pointer(req.counters_), 1)
// 	c.Print()
// }

func test() {
	p := C.alloc()
	size := *p
	fmt.Println(size)
	eles := (*[1 << 30]uint32)(unsafe.Pointer(p))[1 : size+1 : size+2]
	output(eles)
}

func output(arr []uint32) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	// p := C.new_pacing(1, 1)
	// p := &C.struct_Pacing{}
	// print(uintptr(unsafe.Pointer(p)))
	test()
}


s.Code =                 code
s.Expense =              expense
s.DailyExpense =        dailyExpense
s.LifeStage =           lifeStage
s.FillRate =            fillRate
s.ExtFillRate =         extFillRate
s.ExpectedProgress =    target
s.ExpectedExtProgress = extTarget
s.DailyFillRate =        dailyFillRate
