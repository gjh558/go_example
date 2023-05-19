package main

import (
	"log"
	"syscall"

	"freewheel.com/proto/person"
	"google.golang.org/protobuf/proto"
)

func main() {
	b, err := syscall.Mmap(-1, 0, syscall.Getpagesize(), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		log.Printf("Mmap: %v", err)
	}
	log.Printf("mapped size: %d\n", len(b))
	defer func() {
		syscall.Munmap(b)
	}()

	p := person.Person{
		Age:    10,
		Name:   10,
		Salary: 10,
		Child: &person.Child{
			Age:    10,
			Name:   10,
			Salary: 10,
		},
	}

	s := b[0:0:4096]
	log.Printf("s len %d, cap %d\n", len(s), cap(s))
	b1 := proto.MarshalOptions{Deterministic: true}

	s, err = b1.MarshalAppend(s, &p)
	log.Printf("%d: %+v\n", len(s), s)

	// s1 := make([]byte, len(b1.Bytes()))
	//
	// copy(s1, b1.Bytes())
	// log.Printf("s1: %+v", s1)
	// var p2 person.Person
	// proto.Unmarshal(s1, &p2)
	// log.Printf("unmarshal: %+v", p2)
}
