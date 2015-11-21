package main
import (
	"testing"
	"time"
)

func TestAdd(t *testing.T){
	result :=Add(1,2)
	time.Sleep(3*time.Second)
	if result!=3{
		t.Log("")
		t.Fail()
	}

	varresult :=Add(1,2,3,4)
	if varresult!=10{
		t.Error("Vartesterror.")
	}
}