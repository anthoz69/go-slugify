package slugify

import (
	"testing"
)

func TestSlugify(t *testing.T) {
	s := "test->àèâ<-test"
	slug := Marshal(s)
	expected := "test-aea-test"
	if slug != expected {
		t.Fatal("Return string is not slugified as expected", expected, slug)
	}
}

func TestLowerOption(t *testing.T) {
	s := "Test->àèâ<-Test"
	slug := Marshal(s, true)
	expected := "test-aea-test"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
	slug = Marshal(s, false)
	expected = "Test-aea-Test"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
	slug = Marshal(s)
	expected = "Test-aea-Test"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
}

func TestSlugifyTH(t *testing.T) {
	s := "test->àèâ<-test-ฉันไปกินข้าว กับ    เพื่อน   "
	slug := Marshal(s)
	expected := "test-aea-test-ฉันไปกินข้าว-กับ-เพื่อน"
	if slug != expected {
		t.Fatal("Return string is not slugified as expected", expected, slug)
	}
}

func TestLowerOptionTH(t *testing.T) {
	s := "TestT->àèâ<-Test ฉันไปกินข้าว"
	slug := Marshal(s, true)
	expected := "testt-aea-test-ฉันไปกินข้าว"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
	slug = Marshal(s, false)
	expected = "TestT-aea-Test-ฉันไปกินข้าว"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
	slug = Marshal(s)
	expected = "TestT-aea-Test-ฉันไปกินข้าว"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
	slug = Marshal(`เมาท์ก่อนหน้าอัลไซเมอร์จัมโบ้เอเซีย เทรลเลอร์ช็อป ปิ้งคอมเพล็กซ์แฟรี่ บลูเบอร์รี`)
	expected = "เมาท์ก่อนหน้าอัลไซเมอร์จัมโบ้เอเซีย-เทรลเลอร์ช็อป-ปิ้งคอมเพล็กซ์แฟรี่-บลูเบอร์รี"
	if slug != expected {
		t.Error("Return string is not slugified as expected", expected, slug)
	}
}
