package goutils

import (
	"fmt"
	"testing"
)

func TestLatLng001(t *testing.T) {
	s1 := `17°26'03.13N`
	f, err := DMSStr2Float(s1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f)
}

func TestLatLng002(t *testing.T) {
	s1 := `S17°26'03.13`
	f, err := DMSStr2Float(s1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f)
}

func TestLatLng003(t *testing.T) {
	s1 := `E171°46'`
	f, err := DMSStr2Float(s1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f)
}

func TestLatLng004(t *testing.T) {
	s1 := `171°46'27.1345"W`
	f, err := DMSStr2Float(s1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f)
}

func TestLatLng005(t *testing.T) {
	s1 := `E171°`
	f, err := DMSStr2Float(s1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f)
}

func TestLatLng006(t *testing.T) {
	s1 := `N11°0'16`
	f, err := DMSStr2Float(s1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f)
}

func TestLatLng007(t *testing.T) {
	s1 := `171°46'27.2345"W`
	f, err := DMSStr2Float(s1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(f)
}
