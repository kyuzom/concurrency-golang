# concurrency-golang

Golang concurrency patterns, channels and best practices.

## Fan-in, Fan-out patterns

Fan-in Fan-out is a way of Multiplexing and Demultiplexing in golang. Fan-in refers to processing multiple input data and combining into a single entity. Fan-out is the exact opposite, dividing the data into multiple smaller chunks, distributing the work amongst a group of workers to parallelize CPU use and I/O.

It’s a way to converge and diverge data into a single data stream from multiple streams or from one stream to multiple streams or pipelines.
