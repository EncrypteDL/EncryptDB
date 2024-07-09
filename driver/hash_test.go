package driver

import (
	"encoding/hex"
	"testing"

	"github.com/amzn/ion-go/ion"
)

func TestToEldbhash(t *testing.T) {
	// Test cases with input values and expected hash results
	tests := []struct {
		value    interface{}
		expected string
	}{
		{value: "test string", expected: "98f8e013db44e5aaf1f1e9a9c28944af05c6ff5bb9a4907c5c80c52acafeffde"},
		{value: 12345, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12346, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12347, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12348, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12349, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12350, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12351, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12352, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12353, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12354, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
		{value: 12355, expected: "827ccb0eea8a706c4c34a16891f84e7b"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			result, err := toEldbhash(test.value)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			resultHex := hex.EncodeToString(result.hash)
			if resultHex != test.expected {
				t.Errorf("expected %v, got %v", test.expected, resultHex)
			}
		})
	}
}
