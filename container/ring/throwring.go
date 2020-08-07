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

	opt options
}

type Opt interface {
	apply(*options)
}

// **尝试**获取即将被覆盖数据,buffer越大，丢失几率越低
func WithTakeThrow(ThrowBufferSize int) Opt {
	return newOptFuncImpl(func(o *options) {
		o.IsTakeThrow = true
		o.ThrowBufferSize = ThrowBufferSize
		o.ThrowC = make(chan interface{}, ThrowBufferSize)
	})
}

var DefaultOpt = options{}

//size mast be 2^n
func NewThrowRing(size int, o ...Opt) *ThrowRing {
	if size <= 0 && (size&(size-1)) != 0 {
		panic("param size mast be 2^n and can not be zero")
	}

	t := &ThrowRing{
		buf:  make([]interface{}, size),
		size: size,
		mu:   sync.Mutex{},
		opt:  DefaultOpt,
	}

	for _, v := range o {
		v.apply(&t.opt)
	}
	return t
}

func (q *ThrowRing) Add(elem interface{}) {

	q.mu.Lock()
	q.buf[q.head] = elem
	//next head
	q.head = (q.head + 1) & (q.size - 1)
	if q.head == q.tail {
		q.beforeThrow()
		q.tail = (q.tail + 1) & (q.size - 1)
	}
	q.mu.Unlock()
}

func (q *ThrowRing) Get() interface{} {
	q.mu.Lock()

	ret := q.buf[q.tail]
	q.buf[q.tail] = nil
	q.tail = (q.tail + 1) & (q.size - 1)

	q.mu.Unlock()

	return ret

}

func (q *ThrowRing) String() {
	fmt.Printf("ThrowRing: %+v \n ", q)
}

func (q *ThrowRing) GetThrowC() <-chan interface{} {
	if q.opt.IsTakeThrow {
		return q.opt.ThrowC
	}
	return nil
}

func (q *ThrowRing) beforeThrow() {
	if q.opt.IsTakeThrow &&
		len(q.opt.ThrowC) < q.opt.ThrowBufferSize {
		q.opt.ThrowC <- q.buf[q.tail]
	}

}

type options struct {
	IsTakeThrow     bool
	ThrowBufferSize int
	ThrowC          chan interface{}
}

type optFunc func(o *options)

type optFuncImpl struct {
	f optFunc
}

func (odo *optFuncImpl) apply(o *options) {
	odo.f(o)
}

func newOptFuncImpl(f optFunc) *optFuncImpl {
	return &optFuncImpl{f: f}
}
