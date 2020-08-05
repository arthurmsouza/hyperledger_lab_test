package main

import "testing"

func TestHello(t *testing.T) {

	emptyResult := hello("")

	if emptyResult != "Hello Dude!" {
		t.Errorf("Failed")
	} else {
		t.Logf("Success")
	}

	//test for valid
	result := hello("aluno")

	if result != "Hello aluno!" {
		t.Errorf("Failed Hello aluno")
	} else {
		t.Logf("Success Hello aluno")
	}

}
