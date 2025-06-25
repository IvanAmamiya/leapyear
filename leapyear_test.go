package main

import (
    "testing"
)

func TestLeapYear_CommonYear(t *testing.T) {
    if leapyear(2001) != false {
        t.Errorf("2001 should not be a leap year")
    }
}

func TestLeapYear_DivisibleBy4(t *testing.T) {
    if leapyear(1996) != true {
        t.Errorf("1996 should be a leap year")
    }
}

func TestLeapYear_DivisibleBy100(t *testing.T) {
    if leapyear(1900) != false {
        t.Errorf("1900 should not be a leap year")
    }
}

func TestLeapYear_DivisibleBy400(t *testing.T) {
    if leapyear(2000) != true {
        t.Errorf("2000 should be a leap year")
    }
}
