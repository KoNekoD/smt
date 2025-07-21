# Go Slice & Map Tools

Go helper library for extra slice and map operations

## Usage

```go
package main

import (
	"fmt"
	"github.com/KoNekoD/smt/pkg/smt"
	"slices"
)

func main() {
    // IterToSlice
    iter1 := func(yield func(int) bool) {
        for i := 0; i < 20; i++ {
            if !yield(i) {
                return
            }
        }
    }
    
    slice1 := smt.IterToSlice(iter1)
    
    fmt.Println(slice1) // [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]
    
    // IterToMap
    iter2 := func(yield func(int, int) bool) {
        for i := 0; i < 10; i++ {
            if !yield(i*10, i) {
                return
            }
        }
    }
    
    map2 := smt.IterToMap(iter2)
    
    fmt.Println(map2) // map[0:0 10:1 20:2 30:3 40:4 50:5 60:6 70:7 80:8 90:9]
    
    // MapSlice
    slice3 := smt.MapSlice(slice1, func(i int) int {
        return i * 2
    })
    
    fmt.Println("MapSlice", slice3) // MapSlice [0 2 4 6 8 10 12 14 16 18 20 22 24 26 28 30 32 34 36 38]
    
    // Map
    map4 := smt.Map(map2, func(k int, v int) int {
        return k + v
    })
    
    fmt.Println("Map", map4) // Map map[0:0 10:11 20:22 30:33 40:44 50:55 60:66 70:77 80:88 90:99]
    
    // SliceIntersect
    slice5 := smt.SliceIntersect(slice1, slice3)
    
    slices.Sort(slice5) // sort for deterministic result
    
    fmt.Println("SliceIntersect", slice5) // SliceIntersect [0 2 4 6 8 10 12 14 16 18]
    
    // SliceFlip
    slice6 := smt.SliceFlip(slice5)
    
    fmt.Println("SliceFlip", slice6[0], slice6[2], slice6[4], slice6[6]) // SliceFlip 0 1 2 3
    
    // SliceShift
    slice7 := []int{99, 10, 2}
    fmt.Println("SliceShiftBefore", slice7) // SliceShiftBefore [99 10 2]
    itemFromSlice, slice7 := smt.SliceShift(slice7)
    fmt.Println("SliceShift", *itemFromSlice) // SliceShift 99
    fmt.Println("SliceShiftAfter", slice7)    // SliceShiftAfter [10 2]
    
    // SliceSearch
    index := smt.SliceSearch(10, slice7)
    fmt.Println("SliceSearch", index) // SliceSearch 0
    index2 := smt.SliceSearch(11, slice7)
    fmt.Println("SliceSearch", index2) // SliceSearch -1
    
    // SafeSlice
    slice8 := []int{1, 2, 3, 4, 5}
    fmt.Println(smt.SafeSlice(slice8, 1, 3)) // [2 3 4]
    
    // SliceReverse
    slice9 := []string{"a", "b", "c", "d", "e"}
    fmt.Println(smt.SliceReverse(slice9)) // [e d c b a]
    
    // SliceUniq
    slice10 := []string{"a", "b", "c", "d", "e", "e", "e", "e"}
    fmt.Println(smt.SliceUniq(slice10)) // [c d e a b]
}

```
