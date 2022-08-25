/* Alta3 Research | RZFeeser
   Errors - Creating basic errors with fmt.Errorf */

package main

import (
    "fmt"
    "time"
    "os"
    "os/exec"
)

func main() {


    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    // build a error message with the current time
    err := fmt.Errorf("error occurred at: %v", time.Now()) // use a formatting verb to add info to the error
    fmt.Println("An error happened:", err)
}

