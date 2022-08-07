package adyen

import "testing"

func TestDetectCardType(t *testing.T) {
	if DetectCardType("5123459046058920") != "mc" {
		t.Fail()
	}

	if DetectCardType("4000180000000002") != "visa" {
		t.Fail()
	}
}
