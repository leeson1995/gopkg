package ring

import (
	"fmt"
	"testing"
	"time"
)

func TestGetSet(t *testing.T) {
	q := NewThrowRing(5)
	var count int
	go func() {

		for {
			g, ok := q.Get().(TickRecord)
			if ok {
				fmt.Println("get ", g)
			} else {
				fmt.Println("nil nil nil")
			}

			time.Sleep(time.Millisecond)
		}

	}()

	for {
		q.Add(TickRecord{count})
		(q.String())
		count++
		time.Sleep(time.Millisecond * 2)
	}
}

type TickRecord struct {
	t int
}

func BenchmarkGetSetBench(t *testing.B) {
	q := NewThrowRing(5)
	var count int
	go func() {

		for {
			g, ok := q.Get().(TickRecord)
			if ok {
				fmt.Println("get ", g)
			}

			time.Sleep(time.Millisecond * 2)
		}

	}()

	for count < 1000 {
		q.Add(TickRecord{count})
		(q.String())
		count++
		time.Sleep(time.Millisecond * 3)
	}
}
