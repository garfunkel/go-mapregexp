package mapregexp

import (
	"testing"
)

func TestFindStringSubmatchMap(t *testing.T) {
	re := MustCompile(`Name: (?P<name>[a-zA-Z]+),\s*Age: (?P<age>\d+)`)
	groups := re.FindStringSubmatchMap("blabla bla Name: Garfunkel, \t\nAge: 6?")
	name, ok := groups["name"]

	if !ok {
		t.Fatal("'name' group not found")
	}

	if name != "Garfunkel" {
		t.Fatal("'name' group value unexpected")
	}

	age, ok := groups["age"]

	if !ok {
		t.Fatal("'age' group not found")
	}

	if age != "6" {
		t.Fatal("'age' group value unexpected")
	}

	if len(groups) != 2 {
		t.Fatal("expected two matched groups")
	}
}
