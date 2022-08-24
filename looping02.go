/* Alta3 Research | RZFeeser
   Looping - for loop behaving like a while loop    */

package main

import "fmt"
import "math/rand"
import "os"
import "os/exec"

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    drive := 0 // drive is defined at the function level
    fmt.Print("Get a mulligan on any drive under 60 yards.\n")

    // typically after "for" is the init var, which is omitted
    // no post statement appears after the last semi-colon
    for ; drive <= 60; {           // same as: for drive <= 60
        drive = rand.Intn(70)
        fmt.Print("SWING!\n")
    }
    fmt.Println("Your longest drive was", drive, "yards")
}

