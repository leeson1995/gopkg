// Copyright 2016 ~ 2019 Leeson(https://github.com/Leeson1995).
// All rights reserved.  Use of this source code is
// governed by Apache License 2.0.

package ring

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkGetSet(b *testing.B) {
	type TestStruct struct {
		t int
	}

	q := NewThrowRing(100)
	b.Run("Get", func(b *testing.B) {
		go func() {

			for {
				g, ok := q.Get().(TestStruct)
				if ok {
					b.Log("get ", g)
				}
			}

		}()
	})

	var count int

	b.Run("set ", func(b *testing.B) {
		for count < 1000 {
			q.Add(TestStruct{count})
			count++
		}
	})

}

func TestGetLag(t *testing.T) {
	type TickRecord struct {
		t int
	}
	q := NewThrowRing(10)
	var count int
	go func() {
		for {
			g, ok := q.Get().(TickRecord)
			if ok {
				fmt.Println("get ", g)
			}

			time.Sleep(time.Millisecond * 40)
		}

	}()

	for count < 100 {
		q.Add(TickRecord{count})
		count++
		time.Sleep(time.Millisecond * 20)
	}

}
func TestSetLag(t *testing.T) {
	type TickRecord struct {
		t int
	}
	q := NewThrowRing(10)
	var count int
	go func() {
		for {
			g, ok := q.Get().(TickRecord)
			if ok {
				fmt.Println("get ", g)
			}

			time.Sleep(time.Millisecond * 20)
		}

	}()

	for count < 100 {
		q.Add(TickRecord{count})
		count++
		time.Sleep(time.Millisecond * 40)
	}

}
func TestWithTakeThrow(t *testing.T) {
	type TickRecord struct {
		t int
	}
	q := NewThrowRing(100, WithTakeThrow(100))
	var count int
	go func() {
		c := q.GetThrowC()
		if c != nil {

			for {
				d := <-c
				fmt.Println("thrwo: ", d)
			}
		}

	}()
	go func() {
		for {
			g, ok := q.Get().(TickRecord)
			if ok {
				fmt.Println("get ", g)
			}

			time.Sleep(time.Millisecond * 40)
		}

	}()

	for count < 100 {
		q.Add(TickRecord{count})
		count++
		time.Sleep(time.Millisecond * 20)
	}

}
