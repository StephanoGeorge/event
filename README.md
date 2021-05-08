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
	for i := 0; i < 3; i++ {
		go func() {
			e.Set()
		}()
	}

	time.Sleep(time.Second)
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			e.Wait()
			c <- i
		}(i)
	}

	time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		go func() {
			e.Clear()
		}()
	}

	time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		select {
		case <-time.After(time.Second):
			panic("Fail")
		case v := <-c:
			fmt.Println(v)
		}
	}
}
```
