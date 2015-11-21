package main
import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkExample(b *testing.B) {
	temp, _ := template.New("").Parse("Go")


	var buf bytes.Buffer

	for i :=0; i < b.N; i++ {
		temp.Execute(&buf, nil)
		buf.Reset()
	}

}