/* RZFeeser | Alta3 Research
   The pointers of a type are initialized using the address-of(&)
   operator on a subject of that specific type.  */
   
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

    var q int = 42
    var p *int      // declare the pointer, "p is a pointer to a var that will be an int type"
    p = &q          // initialize the pointer
    fmt.Println(p)  // memory address (0x20e010)
}

