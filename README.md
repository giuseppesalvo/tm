# setTimeout and setInterval in golang

Remember to be always thread safe.

## setInterval

```go

import "github.com/giuseppesalvo/tm-go"

timeout := tm.setTimeout(func () {
    
    fmt.Print("fired!")

}, 500) // Milliseconds

tm.clearTimeout(timeout)

```

**Struct**

```go

type Timeout struct {
    Timer *time.Timer
    Fired bool
}

```


## setInterval

```go

import "github.com/giuseppesalvo/tm-go"

interval := tm.setInterval(func () {
    
    fmt.Print("tick!")

}, 500) // Milliseconds

tm.clearInterval(interval)

```

**Struct**

```go

type Interval struct {
    Ticker *time.Ticker
    Running bool
}

```