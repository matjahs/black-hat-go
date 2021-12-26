package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func portRange(from, to int) []int {
	N := to - from + 1
	ports := make([]int, N)
	for i := 0; i < N; i++ {
		ports[i] = from + i
	}
	return ports
}

func Test_worker(t *testing.T) {
	tests := []struct {
		name  string
		ports []int
		exp   []int
	}{
		{
			name:  "default case",
			ports: portRange(70, 90),
			exp:   []int{80},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ports := make(chan int)
			results := make(chan int)
			var actual []int

			go worker(ports, results)

			go func() {
				for _, p := range tt.ports {
					ports <- p
				}
			}()

			for i := 0; i < len(tt.ports); i++ {
				r := <-results
				if r != 0 {
					actual = append(actual, r)
				}
			}

			close(ports)
			close(results)

			assert.Equalf(t, tt.exp, actual, "")
		})
	}
}
