package propagation

import "testing"

func TestInterpolate(t *testing.T) {
	testCases := []struct {
		x1, y1, x2, y2, xVal, expected float64
	}{
		{1, 10, 5, 50, 3, 30},
		{0, 0, 10, 100, 5, 50},
		{2, 4, 8, 16, 9.55, 19.1},
		{10, 0, 0, 10, 3, 7},
		{0, 10, 10, 0, 4, 6},
		{1, 13.5, -1, 12, 0., 12.75},
		{1, 13.5, -1, 12, 0.75, 13.3125},
	}

	for _, tc := range testCases {
		actual := interpolate(tc.x1, tc.y1, tc.x2, tc.y2, tc.xVal)
		if actual != tc.expected {
			t.Errorf("interpolate(%f, %f, %f, %f, %f) = %f; want %f",
				tc.x1, tc.y1, tc.x2, tc.y2, tc.xVal, actual, tc.expected)
		}
	}
}
