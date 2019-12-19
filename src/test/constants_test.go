package Utility

import (
	"github.com/raunakjodhawat/gorithims/src/Utility"
	"testing"
)

func TestToPrintFlag(t *testing.T) {

	allCreditsTrue := Utility.GetCredits(true)
	allCreditsFalse := Utility.GetCredits(false)
	allCreditsNil := Utility.GetCredits()
	if len(allCreditsTrue) < 2 || len(allCreditsFalse) < 2 || len(allCreditsNil) < 2{
		t.Errorf("There should atlest be two credits")
	}

}

func TestAllComplexitySize(t *testing.T){
	GetAllComplexityTrue := Utility.GetAllComplexity(true)
	GetAllComplexityFalse := Utility.GetAllComplexity(false)
	GetAllComplexityNil := Utility.GetAllComplexity()

	if len(GetAllComplexityTrue) != 2 ||  len(GetAllComplexityFalse) != 2 || len(GetAllComplexityNil) != 2 {
		t.Errorf("There should two set of time complexity and space complexity tasks")
	}

	// Dont change order
	// 4-5 specific points time
	// 4-5 specific points for space
	// last element size 1
}

