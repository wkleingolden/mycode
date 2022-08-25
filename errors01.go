/* Alta3 Research | RZFeeser
   Errors - Creating basic errors with errors.New*/

package main

import (
    "errors"
    "fmt"
    "os"
    "os/exec"
)

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    // this is the error we are creating
    err := errors.New("we don't have the power") // returns a new error with a static message
    fmt.Println("Scotty says:", err)
}

