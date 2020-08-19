package calcpi

import "testing"

var testCalcpi = 400
var testCalcpiAnswer = "3,141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067982148086513282306647093844609550582231725359408128481117450284102701938521105559644622948954930381964428810975665933446128475648233786783165271201909145648566923460348610454326648213393607260249141273724587006606315588174881520920962829254091715364367892590360011330530548820466521384146951941511609"

var testInputParam = 1

func TestInputParam(t *testing.T) {
	_, err := To(testInputParam)
	if err == nil {
		t.Error("test for CaclPi Failed: took a bad parameter")
	}
}

func TestCalcpi(t *testing.T) {

	pi, err := To(testCalcpi)
	if err != nil {
		t.Errorf("test for CaclPi Failed: have error")
	}

	if pi != testCalcpiAnswer {
		t.Errorf("test for CaclPi Failed: incorrect answer %v when %v", testCalcpiAnswer, testCalcpi)
	}

}
