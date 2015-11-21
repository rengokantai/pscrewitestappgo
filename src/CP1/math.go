package main

func Add(num ...int)int {
	var result int
	if len(num) == 0 {
		println("No args")
		return 0
	}
	for _, i := range num {
		result += i
	}
	return result
}

func Subtract(init int, num ...int)int{
	var result = init
	if len(num)==0{
		print("No args")
		return 0
	}
	for _,i :=range num{
		result-=i
	}
	return result
}