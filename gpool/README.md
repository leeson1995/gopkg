

Usage
-----------

Create and run a gpool:

```go
var fn1, fn2 func() // the function which you want to  execute, anonymous functions form closures is better
var fn3, fn4 func(ctx context.Context) // with context, will canceled when pool stop
var limit, jobCount int   // the number of goroutine and job
var wait bool                          // whether blocking

gp := gpool.New(limit, jobCount, wait)
gp.AddJob(fn1)
gp.AddJob(fn2)
gp.AddJobWithCtx(fn3)
gp.AddJobWithCtx(fn4)

if wait {
	gp.Wait()
}
```

termination:

```go
gp.Stop()
...
```
