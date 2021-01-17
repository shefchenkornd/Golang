package main

import (
	"sync"
	"testing"
)

func Test_worker(t *testing.T) {
	type args struct {
		wg *sync.WaitGroup
		mu *sync.Mutex
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}