/*
	1.学习函数可变参数使用（主函数 main function must have no arguments and no return values）
	2.学习任意类型参数使用
*/
package main

import "fmt"

func main() {
	varParameter(1,2,3,4,5,6,7,8)
	allTypeVarParameters(1,"abcd",12.345)
}
//可变参数使用
func varParameter(args ...int){
	for i,arg  := range args {
		fmt.Printf("arg %v = %v\n",i,arg)
	}
}

//任意(Any)类型参数使用,使用interface
func allTypeVarParameters(args ...interface{}) {
	for i,arg := range args {
		switch arg.(type) {		//////////////////////////类型查询
			case int :
				fmt.Printf("arg %d type is int\n",i)
			case string:
				fmt.Printf("arg %d type is string\n",i)
		default:
				fmt.Printf("arg %d type is others\n",i)
		}
	}
}

