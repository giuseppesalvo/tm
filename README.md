# setTimeout and setInterval in golang

Remember to be always thread safe.

## setInterval

```go

tm := setTimeout(func () {
	
	fmt.Print("fired!")

}, 500) // Milliseconds

clearTimeout(tm)

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

tm := setInterval(func () {
	
	fmt.Print("tick!")

}, 500) // Milliseconds

clearInterval(tm)

```

**Struct**

```go

type Interval struct {
	Ticker *time.Ticker
	Running bool
}

```