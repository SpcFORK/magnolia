# Thread Library (thread)

## Overview

`libthread` provides threading primitives and concurrency utilities for Oak, including locks, semaphores, channels, and thread spawning.

## Import

```oak
thread := import('thread')
{ spawn: spawn, Mutex: Mutex, Semaphore: Semaphore } := import('thread')
{ makeChannel: makeChannel, send: send, recv: recv } := import('thread')
```

## Core Functions

### `spawn(fn)`

Spawns a new thread executing the given function.

**Parameters:**
- `fn` - Function to execute in new thread

**Returns:** Thread handle

```oak
{ spawn: spawn } := import('thread')

worker := fn {
    println('Running in background thread')
    // Do work...
}

thread := spawn(worker)
```

### `makeChannel()`

Creates a new channel for thread communication.

**Returns:** Channel object with `send` and `recv` operations

```oak
{ makeChannel: makeChannel, spawn: spawn } := import('thread')

ch := makeChannel()

spawn(fn {
    send(ch, 'Hello from thread!')
})

message := recv(ch)
println(message) // => 'Hello from thread!'
```

### `send(channel, value)`

Sends a value through a channel (blocking until received).

```oak
{ makeChannel: makeChannel, send: send } := import('thread')

ch := makeChannel()
send(ch, 42)
send(ch, 'message')
send(ch, { data: [1, 2, 3] })
```

### `recv(channel)`

Receives a value from a channel (blocking until sent).

**Returns:** Value sent through channel

```oak
{ recv: recv } := import('thread')

value := recv(ch)
println('Received: ' + string(value))
```

## Synchronization Primitives

### `Mutex()`

Creates a mutual exclusion lock.

**Returns:** Mutex object with `lock()` and `unlock()` methods

```oak
{ Mutex: Mutex, spawn: spawn } := import('thread')

counter := 0
mu := Mutex()

fn increment(n) {
    each(range(n), fn {
        mu.lock()
        counter <- counter + 1
        mu.unlock()
    })
}

spawn(fn { increment(1000) })
spawn(fn { increment(1000) })

// Wait for completion...
println(counter) // => 2000 (protected by mutex)
```

### `Semaphore(n)`

Creates a semaphore with initial count.

**Parameters:**
- `n` - Initial semaphore count

**Returns:** Semaphore object with `acquire()` and `release()` methods

```oak
{ Semaphore: Semaphore } := import('thread')

// Limit to 3 concurrent operations
sem := Semaphore(3)

fn limitedTask {
    sem.acquire()
    // Do work (max 3 at a time)
    println('Running...')
    sem.release()
}
```

## Usage Examples

### Producer-Consumer Pattern

```oak
{ makeChannel: makeChannel, spawn: spawn, send: send, recv: recv } := import('thread')

ch := makeChannel()

// Producer
spawn(fn {
    each(range(10), fn(i) {
        println('Producing: ' + string(i))
        send(ch, i)
    })
    send(ch, ?) // Signal completion
})

// Consumer
spawn(fn {
    with std.loop() fn(again) {
        value := recv(ch)
        if value = ? -> {
            println('Done consuming')
        } else {
            println('Consuming: ' + string(value))
            again()
        }
    }
})
```

### Worker Pool

```oak
{ makeChannel: makeChannel, spawn: spawn, send: send, recv: recv } := import('thread')

fn workerPool(numWorkers, jobs) {
    jobCh := makeChannel()
    resultCh := makeChannel()
    
    // Start workers
    each(range(numWorkers), fn(id) {
        spawn(fn {
            with std.loop() fn(again) {
                job := recv(jobCh)
                if job != ? -> {
                    result := processJob(job)
                    send(resultCh, result)
                    again()
                }
            }
        })
    })
    
    // Send jobs
    each(jobs, fn(job) {
        send(jobCh, job)
    })
    
    // Send termination signals
    each(range(numWorkers), fn {
        send(jobCh, ?)
    })
    
    // Collect results
    map(jobs, fn { recv(resultCh) })
}

results := workerPool(4, ['task1', 'task2', 'task3', 'task4'])
```

### Protected Counter

```oak
{ Mutex: Mutex, spawn: spawn } := import('thread')

fn Counter {
    count := 0
    mu := Mutex()
    
    {
        inc: fn {
            mu.lock()
            count <- count + 1
            mu.unlock()
        }
        dec: fn {
            mu.lock()
            count <- count - 1
            mu.unlock()
        }
        get: fn {
            mu.lock()
            value := count
            mu.unlock()
            value
        }
    }
}

counter := Counter()

// Safe concurrent access
spawn(fn { each(range(1000), fn { counter.inc() }) })
spawn(fn { each(range(500), fn { counter.dec() }) })

// ... wait for completion ...
println(counter.get()) // => 500
```

### Rate Limiting with Semaphore

```oak
{ Semaphore: Semaphore, spawn: spawn } := import('thread')

// Allow max 5 concurrent API calls
rateLimiter := Semaphore(5)

each(urls, fn(url) {
    spawn(fn {
        rateLimiter.acquire()
        
        response := http.get(url)
        process(response)
        
        rateLimiter.release()
    })
})
```

### Thread-Safe Queue

```oak
{ Mutex: Mutex, makeChannel: makeChannel, send: send, recv: recv } := import('thread')

fn Queue {
    items := []
    mu := Mutex()
    notEmpty := makeChannel()
    
    {
        enqueue: fn(item) {
            mu.lock()
            items <- append(items, item)
            mu.unlock()
            send(notEmpty, true)
        }
        dequeue: fn {
            recv(notEmpty)
            mu.lock()
            item := items.0
            items <- slice(items, 1, len(items))
            mu.unlock()
            item
        }
        size: fn {
            mu.lock()
            sz := len(items)
            mu.unlock()
            sz
        }
    }
}

queue := Queue()

// Producer
spawn(fn {
    each(range(100), fn(i) {
        queue.enqueue(i)
    })
})

// Consumer
spawn(fn {
    each(range(100), fn {
        item := queue.dequeue()
        println('Dequeued: ' + string(item))
    })
})
```

### Pipeline Pattern

```oak
{ makeChannel: makeChannel, spawn: spawn, send: send, recv: recv } := import('thread')

fn pipeline(stages) {
    channels := map(range(len(stages) + 1), fn { makeChannel() })
    
    each(stages, fn(stage, i) {
        input := channels.(i)
        output := channels.(i + 1)
        
        spawn(fn {
            with std.loop() fn(again) {
                value := recv(input)
                if value != ? -> {
                    result := stage(value)
                    send(output, result)
                    again()
                }
            }
            send(output, ?) // Propagate completion
        })
    })
    
    { input: channels.0, output: channels.(len(stages)) }
}

// Create 3-stage pipeline
pipe := pipeline([
    fn(x) { x * 2 }        // Stage 1: Double
    fn(x) { x + 10 }       // Stage 2: Add 10
    fn(x) { string(x) }    // Stage 3: Convert to string
])

// Send inputs
spawn(fn {
    each(range(5), fn(i) {
        send(pipe.input, i)
    })
    send(pipe.input, ?) // Signal end
})

// Receive outputs
with std.loop() fn(again) {
    result := recv(pipe.output)
    if result != ? -> {
        println('Result: ' + result)
        again()
    }
}
```

### Fan-Out, Fan-In

```oak
{ makeChannel: makeChannel, spawn: spawn, send: send, recv: recv } := import('thread')

fn fanOut(input, workers) {
    outputs := map(range(workers), fn { makeChannel() })
    
    spawn(fn {
        i := 0
        with std.loop() fn(again) {
            value := recv(input)
            if value != ? -> {
                send(outputs.(i % workers), value)
                i <- i + 1
                again()
            }
        }
        // Signal all workers to stop
        each(outputs, fn(ch) { send(ch, ?) })
    })
    
    outputs
}

fn fanIn(inputs) {
    output := makeChannel()
    
    each(inputs, fn(ch) {
        spawn(fn {
            with std.loop() fn(again) {
                value := recv(ch)
                if value != ? -> {
                    send(output, value)
                    again()
                }
            }
        })
    })
    
    output
}

// Usage
jobCh := makeChannel()
workerOutputs := fanOut(jobCh, 4)
resultCh := fanIn(workerOutputs)
```

## Implementation Notes

- Channels are **blocking**: `send()` blocks until `recv()`, and vice versa
- Mutex operations are **blocking**: `lock()` blocks until lock is available
- Semaphore `acquire()` blocks if count is 0
- Threads do not return values directly (use channels)
- No thread joining/waiting primitive (use channels for synchronization)

## Best Practices

### Use Channels for Communication

```oak
// Good: Channels for data sharing
ch := makeChannel()
spawn(fn { send(ch, computeResult()) })
result := recv(ch)

// Avoid: Shared mutable state without locks
sharedData := []
spawn(fn { sharedData <- append(sharedData, 1) }) // Race condition!
```

### Always Unlock Mutexes

```oak
// Good: Always unlock
mu.lock()
// ... work ...
mu.unlock()

// Better: Ensure unlock on all paths
fn withLock(mutex, fn) {
    mutex.lock()
    result := fn()
    mutex.unlock()
    result
}
```

### Signal Completion

```oak
// Use sentinel value (?) to signal completion
spawn(fn {
    each(items, fn(item) {
        send(ch, item)
    })
    send(ch, ?) // Done
})
```

## Limitations

- No thread joining/waiting (use channels instead)
- No thread IDs or introspection
- No thread-local storage
- No priority scheduling
- No thread cancellation
- Channels are unbuffered (synchronous)
- No timeout for channel operations
- No select/multichannel receive
- Mutexes are not recursive
- No read-write locks

## Concurrency Patterns

- **Producer-Consumer**: Use channels
- **Worker Pool**: Use job/result channels
- **Pipeline**: Chain channels through stages
- **Fan-Out**: Distribute work to multiple workers
- **Fan-In**: Collect results from multiple workers
- **Rate Limiting**: Use semaphores
- **Critical Sections**: Use mutexes

## See Also

- Oak built-in `loop()` - For infinite loops
- `std` library - For iteration helpers
- Oak concurrency model documentation
