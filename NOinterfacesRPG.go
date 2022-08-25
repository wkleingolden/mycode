/* Alta3 Research | rzfeeser@alta3.com
   Interfaces - Simple RPG example with NO interfaces */

package main

import (
    "fmt"
    "os"
    "os/exec"
//    "reflect"
)

// Create Wizard struct
type Wizard struct{}

// Create Barbarian struct
type Barbarian struct{}

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    gandalf := Wizard{}
    fmt.Println("Wizard defends:", gandalf.Defend())

    // repeat the same pattern for the Barbarian
    conan := Barbarian{}
    fmt.Println("Barbarian defends:", conan.Defend())

    // continue the pattern for any other players


    fmt.Println("Wizard makes us all forget:", gandalf.Forget())

}

// Wizard Receiver Function (method) - This is how we want the Wizard to Defend()
func (w Wizard) Defend() string {
    return "Expelliarmus"
}

// Wizard Receiver Function (method) - A Wizard has the unique ability to cast a Forget() spell
func (w Wizard) Forget() string {
    return "Obliviate"
}

// Barbarian Receiver Function (method) - This is how we want the Barbarian to Defend()
func (b Barbarian) Defend() string {
    return "Dodge"
}
