/* RZFeeser | Alta3 Research
   Create a struct type named coord to contain
   x and y coordinates */

package main

import (
    "fmt"
    "os"
    "os/exec"
)

// define coord
type coord struct {
    x int
    y int
}

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    fmt.Println(coord{1, 2}) // this coord type can be interacted with

    zulu := coord{42, 100}   // create the struct zulu
    zulu.y = 180             // change the value of zulu to {42, 180}
    fmt.Println(zulu)        // display zulu
    fmt.Println("The x coordinate is:", zulu.x)  // display x int
    fmt.Println("The y coordinate is:", zulu.y)  // display y int
}

