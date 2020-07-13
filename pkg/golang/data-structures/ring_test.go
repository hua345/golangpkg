// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package data_structures

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing_Next(t *testing.T) {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	ringLen := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < ringLen; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	for j := 0; j < ringLen*2; j++ {
		t.Log(r.Value)
		r = r.Next()
	}
}

func TestRing_Prev(t *testing.T) {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring backwards and print its contents
	for j := 0; j < n*2; j++ {
		r = r.Prev()
		t.Log(r.Value)
	}
}

func TestRing_Do(t *testing.T) {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		t.Log(p.(int))
	})
}

func TestRing_Move(t *testing.T) {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Move the pointer forward by three steps
	r = r.Move(3)

	// Iterate through the ring and print its contents
	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}
