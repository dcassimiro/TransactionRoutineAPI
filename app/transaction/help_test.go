package transaction

import "testing"

func Test_amout(t *testing.T) {
	tests := []struct {
		name     string
		op       int
		am       float32
		expected float32
	}{
		{"Test with op=4", 4, 10.0, 10.0},
		{"Test with op!=4", 3, 10.0, -10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := amount(tt.op, tt.am)
			if result != tt.expected {
				t.Errorf("For op=%d and am=%f, expected %f but got %f", tt.op, tt.am, tt.expected, result)
			}
		})
	}
}
