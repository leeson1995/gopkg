// Copyright 2016 ~ 2019 Leeson(https://github.com/Leeson1995).
// All rights reserved.  Use of this source code is
// governed by Apache License 2.0.

package ring

import (
	"fmt"
	"sync"
)

type ThrowRing struct {
	buf []interface{}
	//最新数据位置
	head,
	//消费偏移
	tail,
	//ring 长度
	size int
	mu sync.Mutex
}

func NewThrowRing(size int) *ThrowRing {
	return &ThrowRing{
		buf:  make([]interface{}, size),
		size: size,
		mu:   sync.Mutex{},
	}
}

func (q *ThrowRing) Add(elem interface{}) {

	q.mu.Lock()
	q.buf[q.head] = elem
	//next head
	q.head = (q.head + 1) % (q.size)
	if q.head == q.tail {
		q.tail = (q.tail + 1) % (q.size)
	}
	q.mu.Unlock()
}

func (q *ThrowRing) Get() interface{} {
	q.mu.Lock()

	ret := q.buf[q.tail]
	q.buf[q.tail] = nil
	q.tail = (q.tail + 1) % (q.size)

	q.mu.Unlock()

	return ret

}

func (q *ThrowRing) String() {
	fmt.Printf("ThrowRing: %+v \n ", q)
}
