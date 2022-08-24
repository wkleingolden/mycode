/* RZFeeser | Alta3 Research
   switch - multiple expressions */

package main

import (
    "time"
    "fmt"
    "os"
    "os/exec"
)

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    switch time.Now().Weekday() {

    case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
        fmt.Println("Working for the weekend.")
    case time.Saturday, time.Sunday:
        fmt.Println("Party time!")
    }
}

