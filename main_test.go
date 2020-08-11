package main

import "testing"

var testCalcpi = 100
var testCalcpiAnswer = "3,141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067"

func TestCalcpi(t *testing.T) {

	pi, err := CaclPi(testCalcpi)
	if err != nil {
		t.Errorf("test for CaclPi Failed: have error")
	}

	if pi != testCalcpiAnswer {
		t.Errorf("test for CaclPi Failed: incorrect answer %v when %v", testCalcpiAnswer, testCalcpi)
	}

}
