package counter

import "testing"

func TestCountToUtf81(t *testing.T) {
	s := "hello"
	cset := "utf8"
	expected := 5 * 1

	actual, _ := Count(s, cset)
	if actual != expected {
		t.Fatalf("failed test. expected = %v actual = %v", expected, actual)
	}
}
func TestCountToShiftJIS1(t *testing.T) {
	s := "hello"
	cset := "sjis"
	expected := 5 * 1

	actual, _ := Count(s, cset)
	if actual != expected {
		t.Fatalf("failed test. expected = %v actual = %v", expected, actual)
	}
}

func TestCountToShiftJIS2(t *testing.T) {
	s := "あいうえお"
	cset := "sjis"
	expected := 5 * 2

	actual, _ := Count(s, cset)
	if actual != expected {
		t.Fatalf("failed test. expected = %v actual = %v", expected, actual)
	}
}

func TestCountToShiftJIS3(t *testing.T) {
	s := "日本語andアルファベット混合"
	cset := "sjis"
	expected := 3*1 + 12*2

	actual, _ := Count(s, cset)
	if actual != expected {
		t.Fatalf("failed test. expected = %v actual = %v", expected, actual)
	}
}
