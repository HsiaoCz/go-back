package main

import "testing"

func TestGenIdentity(t *testing.T) {
	identity := GenIdnetity()
	t.Log(identity)
}
