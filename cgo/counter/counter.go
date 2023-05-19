package counter

// #include "../struct.h"
import "C"

import (
  "fmt"
  "unsafe"
)

type Counter struct {
  Data []C.struct_Counter
}

func New(counters unsafe.Pointer, len int) Counter {
    slice := (*[1 << 30]C.struct_Counter)(counters)[:len:len]
    return Counter{slice}
}

func (c Counter) Print() {
  for i, d := range(c.Data) {
    fmt.Printf("%d:  %d\n", i, d.id_)
  }
}
