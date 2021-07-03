# Event in Go

Implementation of Python threading.Event in Go, add channel support

# Example

```go
import (
	"github.com/StephanoGeorge/event"
)
```

```go
e := event.Event()
go func() {
    e.Set()
}()

go func() {
    e.Wait()
}()

go func() {
    e.Clear()
}()
```

With channel

```go
e := event.Event(true)
go func() {
    e.Set()
}()

go func() {
    select {
    case <-e.WaitChan:
        if !e.IsSet() {
            fmt.Println("Cleared")
        }
    case <-time.After(time.Minute):
        fmt.Println("Timed out")
    }
}()

go func() {
    e.Clear()
}()
```
