/*
 * Created by Leeson on 2020/03/09.
 * xgo
 */

package xgo

import (
	"fmt"
	"golang.org/x/exp/rand"
	"math"
	"os"
	"runtime/debug"
	"time"
)

const (
	defaultRetryTimes = math.MaxInt32
	defaultInterval   = 100
)

func Go(fn func()) {
	go goSafe(fn, nil)
}

func GoWithCleaner(fn func(), cleaner ...func()) {
	go goSafe(fn, cleaner...)
}

func goSafe(fn func(), cleaner ...func()) {
	defer Recover(cleaner...)
	fn()
}

type (
	RetryOption func(*retryOptions)

	retryOptions struct {
		retryMode RetryMode
		times     int
		t         *time.Ticker
	}
)

//退避策略：线性退避、随机退避、指数退避
type RetryMode int

const (
	None RetryMode = iota
	LinearBackoff
	RandomBackoff
	ExponentialBackoff
)

func WithTimes(times int) RetryOption {
	return func(options *retryOptions) {
		options.times = times
	}
}
func WithBackoff(mode RetryMode) RetryOption {
	return func(opt *retryOptions) {
		opt.retryMode = mode
		opt.t = time.NewTicker(defaultInterval * time.Millisecond)
	}
}

func GoWithRetries(fn func(), opts ...RetryOption) {
	var options = newRetryOptions()
	for _, opt := range opts {
		opt(options)
	}
	for i := 0; i < options.times; i++ {
		go goSafe(fn)
		options.backoff()
	}

}

func newRetryOptions() *retryOptions {
	return &retryOptions{
		times:     defaultRetryTimes,
		retryMode: None,
	}
}
func (options *retryOptions) backoff() {
	if options.retryMode != None {
		<-options.t.C
		switch options.retryMode {
		case LinearBackoff:
		case RandomBackoff:
			options.t.Reset(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}
}

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		fmt.Fprintf(os.Stderr, "recover occurs %+v , stack:%s \n", p, string(debug.Stack()))
	}
}
