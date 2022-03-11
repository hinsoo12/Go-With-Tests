package main

import "testing"

func TestMain(t *testing.T) {
	achievedMSG := greeting("Nahom!")
	requiredMSG := "Hello, Nahom! How are you doing?"

	if achievedMSG != requiredMSG {
		t.Errorf("Achieved Message is %q Required Message is %q", achievedMSG, requiredMSG)
	}
}