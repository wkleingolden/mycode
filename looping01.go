/* Alta3 Research | RZFeeser
   Looping - basic for-loop    */

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

    total := 0  // init total

    // init;      condition; post statement
    for alta := 0; alta < 4; alta++ {
        fmt.Println("The value of alta -> ", alta)

        total += alta   // use total to keep track of alta
    }

fmt.Println("\nThe value of total is", total)

}

