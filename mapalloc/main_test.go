package main

import "testing"

func BenchmarkMapAllocValue(b *testing.B) {
	m := make(map[string]int, 1)
	m["foo"] = 1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m["foo"]++
	}
}

func BenchmarkMapAllocPointer(b *testing.B) {
	m := make(map[string]*int, 1)
	v := 1
	m["foo"] = &v

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		(*m["foo"])++
	}
}
