package token

import "testing"

func TestLookupIdent_Function(t *testing.T) {
	tests := []struct {
		name         string
		expectedType Type
	}{
		{"def", FUNCTION},
		{"let", LET},
		{"return", RETURN},
		{"if", IF},
		{"else", ELSE},
		{"true", TRUE},
		{"false", FALSE},
		{"x", IDENT},
	}

	for i, test := range tests {
		ident := LookupIdent(test.name)
		if ident != test.expectedType {
			t.Fatalf("tests[%d] - ident wrong. expected=%q, actual=%q", i, test.expectedType, ident)
		}
	}
}
