package xgo

import (
	"fmt"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

func GoSafely(wg *sync.WaitGroup, ignoreRecover bool, handler func(), catchFunc func(r interface{})) {
	if wg != nil {
		wg.Add(1)
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if !ignoreRecover {
					fmt.Fprintf(os.Stderr, "%s goroutine panic: %v\n%s\n",
						time.Now(), r, string(debug.Stack()))
				}
				if catchFunc != nil {
					if wg != nil {
						wg.Add(1)
					}
					go func() {
						defer func() {
							if p := recover(); p != nil {
								if !ignoreRecover {
									fmt.Fprintf(os.Stderr, "recover goroutine panic:%v\n%s\n",
										p, string(debug.Stack()))
								}
							}

							if wg != nil {
								wg.Done()
							}
						}()
						catchFunc(r)
					}()
				}
			}
			if wg != nil {
				wg.Done()
			}
		}()
		handler()
	}()
}
