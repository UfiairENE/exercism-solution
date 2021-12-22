package hamming

import (
    "fmt"
    "errors"
)
func Distance(a, b string) (int, error) {
    var distance int
    if len(a) != len(b) {
        return distance, errors.New(fmt.Sprintf("length of a (%d) != length of b (%d)", len(a), len(b)))
    }
    for i :=0; i < len(a); i++ {
        if a[i] != b[i] {
            distance++
        }
    }
    return distance, nil
	
}
