package standardAlgorithims

import (
	"github.com/raunakjodhawat/gorithims/src/standardAlgorithims/utility"
	"testing"
)

func TestToPrintFlag(t *testing.T) {
	allCreditsTrue := utility.GetCredits(true)
	allCreditsFalse := utility.GetCredits(false)
	allCreditsNil := utility.GetCredits()
	if len(allCreditsTrue) < 2 || len(allCreditsFalse) < 2 || len(allCreditsNil) < 2 {
		t.Errorf("There should atlest be two credits")
	}

}

func TestAllComplexitySize(t *testing.T) {
	GetAllComplexityTrue := utility.GetAllComplexity(true)
	GetAllComplexityFalse := utility.GetAllComplexity(false)
	GetAllComplexityNil := utility.GetAllComplexity()

	if len(GetAllComplexityTrue) != 2 || len(GetAllComplexityFalse) != 2 || len(GetAllComplexityNil) != 2 {
		t.Errorf("There should two set of time complexity and space complexity tasks")
	}
}
