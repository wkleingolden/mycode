/* RZFeeser | Alta3 Research
   if and else statements     */

package main

import "fmt"
import "os"
import "os/exec"

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

var age int = 42

fmt.Println("given that age is", age) // show age

if age%2 == 0 {                       // if age mod 2 is 0
fmt.Println("age is even")
} else {                              // in all other cases
fmt.Println("age is odd")
    }
}

