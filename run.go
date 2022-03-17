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
	e        = 123
	f string = "12hello world"
	g byte   = 'h'
)

type st struct {
	name string
	age  int
	next *st
}

func ech(p *st) {
	for p != nil {
		fmt.Println("p is:", *p)
		p = p.next
	}
}

func write(ch chan int)  {
	for i := 0; i < 50; i++ {
		ch <- i
		fmt.Println("i is:", i)
	}
}

func read(ch chan int)  {
	//for  {
	//	var r int
	//	r = <- ch
	//
	//	fmt.Println(" r is:", r)
	//}

	for  {
		<- ch
		fmt.Println(" r is:", time.Now())
	}
}

func main() {
	num := services.GetNumber()
	gin.Default()

	//基础复习

	tt := []byte(`12我的abc？\n`)
	content := string([]rune(f))[:4]
	fmt.Println("a is:", a, b, c, d, e, f, g, content, string(tt)[0:6])

	second := time.Now().UnixNano()

	fmt.Println("just a test", num, second)

	var aa []int

	aa = []int{1, 2, 3, 4}

	var bb = [...]int{1, 2, 3, 4, 5}

	cc := bb[0:3]

	dd := make([]int, 3, 5)

	for i := 0; i < 3; i++ {
		dd[i] = i
	}

	fmt.Println("aa is:", aa[2], " cc is:", cc, "dd is:", dd)
	fmt.Printf("cc:%+v\n,dd:%+v\n", cc, dd[0:2])

	ee := rand.Intn(100)
	fmt.Println("ee is:", ee)

	for i, v := range cc {
		fmt.Println("i: v: ", i, v)
	}

	//继续基础
	var ff = new(int)
	*ff = 123

	fmt.Println("ff is:", *ff)

	//继续基础

	//链表尾部插入数据
	//var h st
	//h.name = "jack"
	//h.age = 12
	//
	//var hp = &h
	//for i := 0; i < 10; i++ {
	//	st1 := st{
	//		name: fmt.Sprintf("st:%d", i),
	//		age: rand.Intn(500),
	//	}
	//
	//	hp.next = &st1
	//	hp = &st1
	//}
	//
	//ech(&h)

	//链表头部插入数据
	var h = new(st)
	h.name = "jack"
	h.age = 12

	for i := 0; i < 10; i++ {
		st1 := st{
			name: fmt.Sprintf("st:%d", i),
			age:  rand.Intn(500),
		}

		st1.next = h
		h = &st1 //修改h的地址
	}

	ech(h)

	//测试接口

	//var hh interface{}
	//var jj = 123
	//var kk int
	//hh = jj

	//kk = hh.(int)

	//switch v := hh.(type) {
	//case int:
	//	fmt.Println("kk is...:", v)
	//default:
	//	fmt.Println("kk is...:","...")
	//}

	//kk := reflect.ValueOf(hh)

	//再转为interface{}
	//ll := kk.Interface()
	//fmt.Println("kk is:", kk)
	//
	//var vv = int64(123)
	//kk.Elem().SetInt(vv)
	//fmt.Println("kk to string:",kk.Elem().Int())

	//switch v := ll.(type) {
	//case interface{}:
	//	fmt.Println("ll is:", v)
	//default:
	//	fmt.Println("ll is:...")
	//}

	//channel测试
	intChan := make(chan int)
	go write(intChan)
	go read(intChan)

	time.Sleep(10 * time.Second)
	defer func() {
		fmt.Println("10 秒结束了")
	}()
}
