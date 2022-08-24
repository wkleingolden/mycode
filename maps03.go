/* Alta3 Research | RZFeeser
   Maps - Using empty interfaces as a "wildcard" to create
   composite types */

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

    // instead of declaring a type, we have supplied the "empty interface"
    // interface{} be in the place of the key, value, or both when making the map
    futurama := make(map[string]interface{})

    // now we can create a map of mixed types
    futurama["name"] = "Turanga Leela"  // string: string
    futurama["age"] = 30                // string: int
    futurama["height"] = 182.5          // string: float

    // display the "mixed" results
    fmt.Printf("%+v\n", futurama)
}
