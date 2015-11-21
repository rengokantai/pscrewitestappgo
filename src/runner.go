package main

import(
	"fmt"
	"math/rand"
	"testing/quick"
	"time"
	"reflect"
)

func main(){
	val, ok :=quick.Value(reflect.TypeOf(1),rand.New(rand.NewSource(time.Now().Unix())))

	if ok{
		fmt.Println(val.Int())
	}
}

