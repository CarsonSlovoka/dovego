package utils

import (
	"fmt"
	"testing"
)

type testItem struct {
	actual   interface{}
	expected interface{}
}

func (t *testItem) test() bool {
	if t.actual != t.expected {
		return false
	}
	return true
}

func TestContains(t *testing.T) {
	mySlice := []interface{}{"data1", 5, "foo"}
	mySlice2 := []string{"data1", "foo"}
	for numCase, testCase := range []testItem{
		{ContainsAny(mySlice, 5), true},
		{ContainsAny(mySlice, "5"), false},
		{ContainsAny(mySlice, "foo"), true},
		{ContainsAny(mySlice, "bar"), false},
		{ContainsAny(mySlice2, "bar"), false},
		{ContainsAny(mySlice2, "foo"), true},

		{ContainsAll(mySlice, 5), true},
		{ContainsAll(mySlice, 5, 4), false},
		{ContainsAll(mySlice, "data1", "foo"), true},
		{ContainsAll(mySlice, "data1", "foo2"), false},
		{ContainsAll(mySlice2, "data1", "foo"), true},
		{ContainsAll(mySlice2, "data1", "foo2"), false},
	} {
		if ok := testCase.test(); !ok {
			fmt.Printf("%+v != %+v", testCase.actual, testCase.expected)
			t.Error(numCase)
		}
	}
}
