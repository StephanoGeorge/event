# Event in Go

Implementation of Python threading.Event in Go, add channel support

# Features

- Can be `Set()` and `Clear()` multiple times
- Add Channel support, can receive events from then channel every time `Clear()` is called

# Example

```go
import (
	"github.com/StephanoGeorge/event"
)
```

```go
e := event.Event()

go func() {
    e.Wait()
}()

go func() {
    e.Clear()
}()

go func() {
    e.Set()
}()
```

With channel

```go
e := event.Event(true)

go func() {
    for {
        select {
        case <-e.WaitChan:
            fmt.Println("Cleared")
        case <-time.After(time.Minute):
            fmt.Println("Timed out")
        }
    }
}()

go func() {
	e.Wait()
}()

go func() {
    e.Clear()
}()

go func() {
	e.Set()
}()

go func() {
	e.Close()
}()
```
