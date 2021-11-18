# Make Simple Table from structure

Stable create Simple Table string from struct

## Usage

```go
package main

import (
	"fmt" 
	"github.com/kirill-scherba/stable"
)

func main() {

	var a = []struct {
		ID     string
		Name   string
		Num    int
		Status string
		IPv4   string
	}{
		{"ququ-1", "ququ-1_name", 1, "ready", "1.1.1.1"},
		{"ququ-2", "ququ-2_name", 2, "ready", "2.2.2.2"},
	}

	st := new(stable.Stable)
	fmt.Print(st.StructToTable(a) + "\n\n")
}
```

Run in [The Go Playground](https://play.golang.org/p/D7csZ6AWHgw)

## License

[BSD](LICENSE)