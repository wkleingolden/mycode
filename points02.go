/* RZFeeser | Alta3 Research
   Dereference operations  */
   
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
    var p *int      // declare the pointer
    p = &q          // initialize the pointer
    fmt.Println(p)  // memory address (0x20e010)        
    fmt.Println(*p) // 42   <-- dereference operation
}

