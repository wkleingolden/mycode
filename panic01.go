/* Alta3 Research | RZFeeser
   Calling the panic() function directly */

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

    // panic produces a quick exit
    panic("Jim, we have a problem.") // this line will be displayed

    fmt.Println("You will not even see this line. The panic creates a fast fail.")
}

