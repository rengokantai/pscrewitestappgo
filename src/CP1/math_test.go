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
		//Log + FailNow=Fatal
	}

	varresult :=Add(1,2,3,4)
	if varresult!=10{
		t.Error("Vartesterror.")
	}
}

func TestSubtract(t *testing.T){
	result :=Subtract(3,2)
	if result!=1{
		t.Fatal("error")
	}
}

func TestFake(t *testing.T){
	//Or testing.Short()
	if(testing.Verbose()){
		t.Skip("Skip verbose")
	}
}