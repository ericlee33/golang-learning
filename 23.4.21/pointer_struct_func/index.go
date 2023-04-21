package main

import "fmt"

type Person struct {
	X int
}

// 需要传入指针，才能 +1
func (p *Person) PlusXWithOne() int {
	p.X = p.X + 1
	return p.X
}

func main() {
	p := Person{X: 1}
	// 编译器会隐式地帮我们用&p去调用PlusXWithOne这个方法。这种简写方法只适用于“变量”，包括struct里的字段比如p.X，以及array和slice内的元素比如perim[0]。我们不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的内存地址就无法获取得到：

	// 	这里的几个例子可能让你有些困惑，所以我们总结一下：在每一个合法的方法调用表达式中，也就是下面三种情况里的任意一种情况都是可以的：

	// 要么接收器的实际参数和其形式参数是相同的类型，比如两者都是类型T或者都是类型*T：

	// Point{1, 2}.Distance(q) //  Point
	// pptr.ScaleBy(2)         // *Point
	// 或者接收器实参是类型T，但接收器形参是类型*T，这种情况下编译器会隐式地为我们取变量的地址：

	// p.ScaleBy(2) // implicit (&p)
	// 或者接收器实参是类型*T，形参是类型T。编译器会隐式地为我们解引用，取到指针指向的实际变量：

	// pptr.Distance(q) // implicit (*pptr)
	p.PlusXWithOne()
	fmt.Println(p.X)
}
