package main

import "testing"

func TestRWLockedMap(t *testing.T) {
	mapTest := NewCommandsMap()
	hotkey1, hotkey2 := Hotkey{0, 45}, Hotkey{0, 46}
	hotkey11, hotkey22 := Hotkey{0, 45}, Hotkey{0, 46}

	mapTest.Set(hotkey1, "something")
	if mapTest.Get(hotkey1) != "something" {
		t.Error("Wrong value")
	}
	if mapTest.Get(hotkey11) != "something" {
		t.Error("Wrong value")
	}

	if mapTest.IsEmpty(hotkey1) {
		t.Error("IsEmpty: False positive")
	}
	if mapTest.IsEmpty(hotkey11) {
		t.Error("IsEmpty: False positive")
	}

	if !mapTest.IsEmpty(hotkey2) {
		t.Error("IsEmpty: False negative")
	}
	if !mapTest.IsEmpty(hotkey22) {
		t.Error("IsEmpty: False negative")
	}

	mapTest.Delete(hotkey11)
	if !mapTest.IsEmpty(hotkey1) {
		t.Error("False negative")
	}
}
