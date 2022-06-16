package main

import "fmt"

func main() {

	fmt.Println(nil == nil)

	//值参数不影响原始值
	//引用参数影响原始值
	a := 10
	fmt.Println(a) //10
	fun1(a)        //值传递
	fmt.Println(a) //100

	fun2(&a)       //引用传递
	fmt.Println(a) //200

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1) //1,2,3,4
	fun3(arr1)        //值传递，内部值有变化
	fmt.Println(arr1) //原始值不变，1,2,3,4

	fun4(&arr1)       //引用传递
	fmt.Println(arr1) //200,2,3,4

}

func fun1(num int) { // 值传递：num = a = 10
	num = 100
}

func fun2(p1 *int) { //传递的是a的地址，就是引用传递
	*p1 = 200
}

func fun3(arr2 [4]int) { // 值传递
	arr2[0] = 100
	fmt.Println(arr2)
}

func fun4(p2 *[4]int) { // 引用传递
	p2[0] = 200
	fmt.Println(p2) //*&
}
