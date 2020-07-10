package main

import (
	"fmt"
	"log"
)

func main()  {
	l := [...]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	fmt.Printf("l 的数据类型为 %T\n", l)

	log.Printf("列表长度为%d  %d", len(l), cap(l))

	k := l[50:60]
	log.Printf("切片长度为%d  %d", len(k), cap(k))
	log.Printf("切片长度为%d  %d", len(l[10:20]), cap(l[10:20]))

	f := append(l[10:20],11)

	log.Printf("F长度为%d  %d", len(f), cap(f))

	log.Println(f)
	f[1] = 5686565
	log.Println(f)
	log.Println(l)
	log.Printf("列表长度为%d  %d", len(l), cap(l))

}
