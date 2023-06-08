package goober_test

import (
	"testing"

	"github.com/nowayhecodes/goober"
)

func Test_AddAndSearch(t *testing.T) {
	st := goober.NewStorage[int](3)

	st.Add("SPa", 1)
	st.Add("SPb", 2)
	st.Add("SPc", 3)

	if l := len(st.Search("SP")); l != 3 {
		t.Fatal("Key SP is expected to have length of 3 but got ", l)
	}

}
