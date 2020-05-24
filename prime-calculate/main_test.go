package main

import (
	"strconv"
	"testing"
)

func TestIsPrime(t *testing.T) {
	cases := map[int]bool {
		1: false,
		2: true,
		3: true,
		4: false,
		5: true,
		7: true,
		9: false,
		11: true,
		17: true,
	}

	for k, v := range cases {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			r := isPrime(k)
			if r != v {
				t.Errorf("%d expect %v, got %v", k, v, r)
			}
		})
	}
}
