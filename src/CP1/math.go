package main

func Add(num ...int)int{
	var result int
	if len(num)==0{
		println("No args")
		return 0
	}
	for _,i :=range num{
		result+=i
	}
	return result
}