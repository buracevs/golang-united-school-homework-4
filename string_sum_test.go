package string_sum

import (
	"testing"
)

func TestStringSumPlus(t *testing.T) {
	res, _ := StringSum("22+11")
	if res != "33" {
		t.Errorf("Expected 2 but was %v", res)
	}
}
func TestStringSumMin(t *testing.T) {
	res, _ := StringSum("22-11")
	if res != "11" {
		t.Errorf("Expected 2 but was %v", res)
	}
}

func TestStringSumPlusNeg(t *testing.T) {
	res, _ := StringSum("-22+11")
	if res != "-11" {
		t.Errorf("Expected 2 but was %v", res)
	}
}

func TestStringSumPlusNegSp(t *testing.T) {
	res, _ := StringSum(" -22 - 11")
	if res != "-33" {
		t.Errorf("Expected 2 but was %v", res)
	}
}

func TestStringSumMinspaces(t *testing.T) {
	res, _ := StringSum("22 - 11")
	if res != "11" {
		t.Errorf("Expected 2 but was %v", res)
	}
}

func TestStringSumTooMuchOperands1(t *testing.T) {
	_, err := StringSum("- 22 - 11 - 5")
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestStringEmptyString(t *testing.T) {
	_, err := StringSum("    ")
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestStringOneDigit(t *testing.T) {
	_, err := StringSum("55")
	if err == nil {
		t.Errorf("Expected error")
	}
}
