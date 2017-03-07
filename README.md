# Vector

Simple Vector library.

### Quick Example

```go
package main

import (
  "fmt"
  "github.com/atedja/go-vector"
)

func main() {
  v1 := vector.New(4)
  v2 := vector.NewWithValues([]float64{0.0, 1.0, 2.0, 3.0})

  result := vector.Add(v1, v2)
}
```

#### [Full API Documentation](https://godoc.org/github.com/atedja/go-vector)

### Basic Usage

#### Creating Vectors

    var v *vector.Vector
    v = vector.New(3)
    v = vector.NewWithValues([]float64{0.0, 1.0, 2.0})

#### Scale Vectors

    v.Scale(2.0)

#### Dot and Cross Products

    v1 := vector.NewWithValues([]float64{0.0, 1.0, 2.0})
    v2 := vector.NewWithValues([]float64{2.0, -1.0, 4.0})
    cross, _ := vector.Cross(v1, v2)
    dot, _ := vector.Dot(v1, v2)

#### Marshal and Unmarshal

Vector supports JSON/Text marshal and unmarshal. Marshalling a Vector returns a JSON array as a string, e.g. `[1,2,3]`.
Marshalling a Vector as Text returns its CSV format, e.g. `1,2,3\n` (with the CRLF character).
Binary marshal/unmarshal shall be supported in future version.
