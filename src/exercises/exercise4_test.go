package exercises_test

import (
	"testing"

	"../exercises"
)

func TestNormal(t *testing.T) {
	var person exercises.Person = &exercises.Normal{Name: "Jose"}
	result := person.Pay(100)
	if result != 100 {
		t.Errorf("expected 100 but got %v", result)
	}
}

func TestJubilado(t *testing.T) {
	var person exercises.Person = &exercises.Jubilado{Name: "Pepe"}
	result := person.Pay(100)
	if result != 75 {
		t.Errorf("expected 75 but got %v", result)
	}
}

func TestInvitado(t *testing.T) {
	var person exercises.Person = &exercises.Invitado{Name: "Luis"}
	result := person.Pay(100)
	if result != 0 {
		t.Errorf("expected 0 but got %v", result)
	}
}
