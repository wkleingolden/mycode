/* Alta3 Research | RZFeeser
   Raising a panic() for an error we are not prepared to handle ourselves. */

package main

import "os"
import "os/exec"

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()


    _, err := os.Create("/carolDanvers/msmarvel")

    // if we have an unexpected error
    if err != nil {
        panic(err)  // fast fail and dump the error to the screen
    }
}

