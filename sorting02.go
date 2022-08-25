/* Alta3 Research | RZFeeser
   Custom Sorting - Sorting by custom functions   */

package main

import (
    "fmt"
    "sort"
    "os"
    "os/exec"
)

// In order to sort by a custom function in Go we need a corresponding type
// byLength type is just an alias for the builtin []string type
type byLength []string

// the sort.Interface - Len, Less, and Swap
func (s byLength) Len() int {
    return len(s)
}
func (s byLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}


// Implement our custom sort by converting the original fruits slice to byLength, 
// and then use "sort.Sort" on that typed slice
func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    fruits := []string{"aaaaaaaaaa", "bb", "peach", "banana", "kiwi"}
    fmt.Println(fruits)
    sort.Sort(byLength(fruits))
    fmt.Println(fruits)
}

