/* RZFeeser | Alta3 Research
   Errors are "just" a type that can be returned by functions. This simple approach
   might be a bit different than other languages you've used.

   Errors can be constructed on the fly using Go’s built-in errors or fmt packages.
   In this example we look at using, errors.New().

   See documentation @ https://pkg.go.dev/errors
   errors.New constructs a basic error value with given message */

package main

import (
    "errors" // preview of using errors
    "fmt"
    "os"
    "os/exec"
)

//name of funct   //input                      //output
func rollchar(firstName string, lastName string) (string, error) {
    // generate an error
    if lastName == "Turnip" || lastName == "Radish" || lastName == "Potato" {
      return lastName, errors.New("Vegetables are not suitable last names for heroes.")
    }
    // desirable response
    return firstName + " the " + lastName, nil
}

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    fmt.Println("Welcome to the Character Generator")

    playerChar, err := rollchar("Gandalf", "Turnip")
    fmt.Println("\nHello?", err, "\n")

    if err != nil {
        fmt.Println("Error while spawning your requested character.")
    } else {
        fmt.Println(playerChar, "has been generated.")
    }
}

