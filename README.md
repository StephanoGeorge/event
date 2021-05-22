# Event in Go

Implementation of Python threading.Event in Go

# Example

```go
import (
	"fmt"
	"time"

	"github.com/StephanoGeorge/event"
)

func main() {
	e := event.Event()
    go func() {
        e.Set()
    }()

    go func(i int) {
        e.Wait()
    }(i)

    go func() {
        e.Clear()
    }()
}
```
