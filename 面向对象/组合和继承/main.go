package main

import (
	"fmt"
	"math"
	"reflect"
)

type ShapeInterface interface {
	GetName() string
	PrintArea()
}

type AreaInterface interface {
	Area() float64
}

// ShapeArea 尝试接口组合
type ShapeArea interface {
	ShapeInterface
	AreaInterface
}

type Shape struct {
	name string
	area *float64
}

func (s *Shape) GetName() string {
	return s.name
}

func (s *Shape) Area() float64 {
	*s.area = 0
	return *s.area
}

// PrintArea 方法和 Area 有依赖关系，怎么拆好？
// 变量好像也不行，试试指针
func (s *Shape) PrintArea() {
	fmt.Println(reflect.TypeOf(*s.area))
	fmt.Printf("%s : Area %v\r\n", s.GetName(), &s.area)
}

// Rectangle 矩形求面积
type Rectangle struct {
	Shape
	w, h float64
}

func (r *Rectangle) Area() float64 {
	*r.area = r.w * r.h
	return *r.area
}

// Circle 圆形: 重新定义 Area 和PrintArea 方法
type Circle struct {
	Shape
	r float64
}

func (c *Circle) Area() float64 {
	*c.area = c.r * c.r * math.Pi
	return *c.area
}

//func (c *Circle) PrintArea() {
//	fmt.Printf("%s : Area %v\r\n", c.GetName(), c.area)
//}

func main() {

	s := Shape{name: "Shape"}
	c := Circle{Shape: Shape{name: "Circle"}, r: 10}
	r := Rectangle{Shape: Shape{name: "Rectangle"}, w: 5, h: 4}

	listshape := []ShapeArea{&s, &c, &r}

	for _, si := range listshape {
		si.PrintArea() //!! 猜猜哪个Area()方法会被调用 !!
	}

}
