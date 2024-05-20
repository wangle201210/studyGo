package main

import (
	"testing"
)

func TestRace1(t *testing.T) {
	race1()
}

func TestRace2(t *testing.T) {
	for i := 0; i < 100; i++ {
		race2()
	}
}

func TestRace3(t *testing.T) {
	for i := 0; i < 100; i++ {
		race3()
	}
}

func TestRace4(t *testing.T) {
	for i := 0; i < 100; i++ {
		race4()
	}
}

func TestRace5(t *testing.T) {
	for i := 0; i < 100; i++ {
		race5()
	}
}

func TestRace6(t *testing.T) {
	race6()
}

func TestRace8(t *testing.T) {
	race8()
}
