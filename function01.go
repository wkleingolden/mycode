/* Alta3 Research | RZFeeser
   Functions - single function   */ 
   
package main

import (
    "fmt"
    "os"
    "os/exec"
)

  //name of funct   //input                      //output
func rollchar(firstName string, lastName string) (string) {
  fullname := firstName + " the " + lastName 
  return fullname
}

func main() {
    
    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    fmt.Println("Welcome to the Character Generator")

    playerChar := rollchar("Gandalf", "Grey")
    fmt.Println(playerChar)
}

