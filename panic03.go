/* Alta3 Research | RZFeeser
   Illegal conditions that will result in a panic. */

package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    magic := make(map[string]string) // create a map that will store keys and values as type "string"

    magic["forests"] = "green"
    magic["mountains"] = "red"
    magic["plains"] = "white"
    magic["seas"] = "blue"
    magic["swamp"] = "black"

    fmt.Println(magic[100])  // an error will be raised when an "int"
                             // is used to access a key expecting a "string"
}

