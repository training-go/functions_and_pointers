package main

import "testing"

func TestSetPersonName(t *testing.T) {
	v := Person{"Nikola", 21}
	v.SetName("Pesho")
	if v.Name != "Nikola" {
		t.Errorf("The name has change, but it should't, want %s, got %s", "Nikola", v.Name)
	}
}

func TestSetPersonAge(t *testing.T) {
	v := Person{"Nikola", 21}
	v.SetAge(25)
	if v.Age != 25 {
		t.Errorf("The age was not set, want %d, got %d", 25, v.Age)
	}
}
