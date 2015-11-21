package main
import (
	"testing"
)

func TestAdd(t *testing.T){
	result :=Add(1,2)
	if result!=3{
		t.Log("")
		t.Fail()
	}

	varresult :=Add(1,2,3,4)
	if varresult!=10{
		t.Error("Vartesterror.")
	}
}