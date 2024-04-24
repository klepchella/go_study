package main

import "fmt"

func plus_one(a int, b int) int {
	return a + b
}

func plus_two(c int, d int, e int) int {
	return c + d + e
}

func plus_tree(aa int, bb int) (int, int, int) {
	return aa + bb, aa, bb
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	// изменяем по адресу
	*iptr = 0
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name, age: 42}
	// p.age = 42
	return &p
}

func main() {
	res := plus_one(1, 2)
	fmt.Println(res)

	res2 := plus_two(10, 20, 30)
	fmt.Println(res2)

	res3_1, res3_2, res3_3 := plus_tree(100, 200)
	fmt.Println(res3_1, res3_2, res3_3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(fact(25))

	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	fmt.Println(newPerson("Jon"))
}
