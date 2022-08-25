/* Alta3 Research | RZFeeser
   exec.Command().Output() - Capturing output from commands   */


package main

import (
    "fmt"
    "log"
    "os/exec"
    "os"
)

func main() {


    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

//    out, err := exec.Command("ls", "-l", "*.xmll").Output()
    out, err := exec.Command("ls").Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
}

