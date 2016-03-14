# Notes

## Intro

Creating good concurrent programs is not an easy task.

In fact some well known languages and frameworks are either single threaded
(node.js, Emacs/elisp, OCaml) or limit scalability to make writing concurrent
programs easier (Python).

However nowdays most of the CPUs are multicore and I beleive the number of
cores will only increase in the future.

It is fairly easy to write concurrent programs in Go but it is important to
understand that concurrency != performance. In this talk we will look into
concurrency primitives in Go and how to use them to write scalable programs.

### "A word of warning" slide

You might make a function 5 times faster but if it takes only 1% of total time
it is not worth it.

Define import measurements and track them to identify improvements and
regressions.

I will show results of some micro-benchmark in this talk. But the only real
performance measurement is how well the final program performs under realistic
load conditions on realistic hardware (Laptop -> RPi, Laptop -> Server).

## Example 1

All results in the presentation are taken on a AWS EC2 c4.8xlarge instance (36
CPU cores, 60GB or RAM).


## Example 2

We will write a program that generates `n` random points in the top right quoter
of the square and count number of points `k` that belongs to the top right
quorter of the circle.

Then Pi = 4*k/n.

### Parallel benchmarks

ns/op should decrease if an operation is scalable. If one person eats 6 durians
in 1 hour then it takes 10 minutes per durian. If two people eat 12 durinas in 1
hour then it takes 5 minutes per durian.

#### mutex

It is fairly fast when there is no contention. If you do not expect huge number
of operations per second just use mutex.

#### channel

Slower than mutex with no contention. slightly faster with 4, 8, 16 cores and
gets slower after.

This is not very surprising. Channel is a generic multi producer multi consumer
queue that and it is hard to optimize. Less generic single producer multiple
consumer queue could perform much better.

It is still possible to write scalable programs using channels. But the
operations should be expensive comparing to channels overhead (which is actually
quite small).

#### no shared state

Here we can see that overhead of mutex with no contention is about 20
nanoseconds.

#### pii

by using X cores the program becomes N times faster.


### Example 3

#### concurrent reads

if all writes to map happen in package init() function or before starting read
goroutines then it is fine to read this map from multiple threads.

#### RWMutex

Why does it not scale? Because of the shared mutable state. RLock
increments number of readers, RUnlock decrements it.


#### Atomic

Writes must be in a single goroutine. Otherwise it is possible to miss some
writes. In case of multiple writers a mutex or channel should be used.





