package main

import "testing"

// 入力と出力が同じかを判定する
func TestHello(t *testing.T) {
	got := Hello("Kohei")
	want := "Hello, Kohei!"

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
