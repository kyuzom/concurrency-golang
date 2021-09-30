# concurrency-golang

Golang concurrency patterns, channels and best practices.

## Fan-in, Fan-out patterns

Fan-in Fan-out is a way of Multiplexing and Demultiplexing in golang. Fan-in refers to processing multiple input data and combining into a single entity. Fan-out is the exact opposite, dividing the data into multiple smaller chunks, distributing the work amongst a group of workers to parallelize CPU use and I/O.

It’s a way to converge and diverge data into a single data stream from multiple streams or from one stream to multiple streams or pipelines.

## Preventing Goroutine Leaks

The runtime handles multiplexing the goroutines onto any number of operating system threads so that we don’t often have to worry about that level of abstraction. But they do cost resources, and goroutines are not garbage collected by the runtime, so regardless of how small their memory footprint is, we don’t want to leave them lying about our process. So how do we go about ensuring they’re cleaned up?

Let’s start from the beginning and think about this step by step: why would a goroutine exist? In Chapter 2, we established that goroutines represent units of work that may or may not run in parallel with each other. The goroutine has a few paths to termination:

* When it has completed its work.
* When it cannot continue its work due to an unrecoverable error.
* When it’s told to stop working.

We get the first two paths for free—these paths are your algorithm—but what about work cancellation? This turns out to be the most important bit because of the network effect: if you’ve begun a goroutine, it’s most likely cooperating with several other goroutines in some sort of organized fashion. We could even represent this interconnectedness as a graph: whether or not a child goroutine should continue executing might be predicated on knowledge of the state of many other goroutines. The parent goroutine (often the main goroutine) with this full contextual knowledge should be able to tell its child goroutines to terminate.
