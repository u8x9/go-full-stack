package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":  {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"nosplit": {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"double":  {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"chinese": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want: %v, got: %v\n", tc.want, got)
			}
		})
	}
}

