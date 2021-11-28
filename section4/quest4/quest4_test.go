package quest4_test

import (
	"program/section4/quest4"
	"testing"
)

func TestBaito_kyuyo(t *testing.T) {
	result := quest4.Baito_kyuyo(10)
	want := 1850
	if result != want {
		t.Errorf("want: %d, result: %d", want, result )
	}
}