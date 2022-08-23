/* RZFeeser | Alta3 Research
   CHALLEGE 01 - Create a struct named Astro */

package main

import "fmt"

type Astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
    num1 int
    num2 int
}

func main() {

    ast1 := Astro{"Megan McArthur", 35, "ISS", false, 1234, 2789}
    ast2 := Astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true, 1, 1}
    ast3 := Astro{"Raja Chari", 39, "SpaceX Crew-3", true, 1, 1}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)

}

