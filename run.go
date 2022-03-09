package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gspgsp/el-api/services"
	"math/rand"
	"time"
)

const (
	a = iota
	b = 1
	c
)

var (
	d int
	e = 123
	f string = "12hello world"
	g byte = 'h'
)

func main()  {
	num := services.GetNumber()
	gin.Default()

	//基础复习

	tt := []byte(`12我的abc？\n`)
	content := string([]rune(f))[:4]
	fmt.Println("a is:", a, b, c,d,e,f,g,content,string(tt)[0:6])

	second := time.Now().UnixNano()

	fmt.Println("just a test", num, second)

	var aa []int

	aa = []int{1,2,3,4}

	var bb = [...]int{1,2,3,4,5}

	cc := bb[0:3]

	dd := make([]int,3,5)

	for i := 0; i < 3; i++ {
		dd[i] = i
	}


	fmt.Println("aa is:",aa[2]," cc is:", cc, "dd is:",dd)
	fmt.Printf("cc:%+v\n,dd:%+v\n", cc,dd[0:2])

	ee := rand.Intn(100)
	fmt.Println("ee is:",ee)

	for i, v := range cc{
		fmt.Println("i: v: ",i,v)
	}
}


