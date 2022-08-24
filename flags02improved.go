/* Alta3 Research | RZFeeser
   Flags - Using the init() function   */
package main

import (
    "flag"
    "fmt"; "os"; "os/exec"
)

var (
    env  *string
    port *int
)

// the 'init()' function is often used to initialize state variables
func init() {
    env = flag.String("env", "development", "current environment")
    port = flag.Int("port", 3000, "port number")
}

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    flag.Parse()

    fmt.Println("env:", *env)
    fmt.Println("port:", *port)
}
