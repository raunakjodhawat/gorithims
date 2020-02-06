package test

import (
	"github.com/raunakjodhawat/gorithims/src/rjson"
	"testing"
)

func TestIsJSON(t *testing.T){
	var b struct{}
	if !rjson.IsJSON(b) {
		t.Errorf("%T Sort failed for input", b)
	}
}
