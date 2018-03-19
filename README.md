# setTimeout and setInterval in golang

## setInterval

```go

import "github.com/giuseppesalvo/tm-go"

timeout := tm.SetTimeout(func () {
    
    fmt.Println("fired!")

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
    
    fmt.Println("tick!")

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

Remember to be always thread safe.
