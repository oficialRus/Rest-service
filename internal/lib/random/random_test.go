package random

import "testing"

const size = 6

func TestRandomResultSize(t *testing.T) {
	result := NewRandomString(size)

	if len(result) != size {
		t.Errorf("FAILED [ result != %d]", size)

	}

}
