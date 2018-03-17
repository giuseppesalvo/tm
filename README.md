# setTimeout and setInterval in golang

Remember to be always thread safe.

## setInterval

```go

import "github.com/giuseppesalvo/tm-go"

timeout := tm.SetTimeout(func () {
    
    fmt.Print("fired!")

}, 500) // Milliseconds

tm.ClearTimeout(timeout)

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

interval := tm.SetInterval(func () {
    
    fmt.Print("tick!")

}, 500) // Milliseconds

tm.ClearInterval(interval)

```

**Struct**

```go

type Interval struct {
    Ticker *time.Ticker
    Running bool
}

```