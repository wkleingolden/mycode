/* RZFeeser | Alta3 Research
Exploring arrays              */
package main

import (
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
)

type Notes struct {
    To      string `xml:"to"`
    From    string `xml:"from"`
    Subject string `xml:"subject"`
    Body    string `xml:"body"`
}

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    data, _ := ioutil.ReadFile("avengers.xml")    // open the file "avengers.xml"

    note := &Notes{}

    _ = xml.Unmarshal([]byte(data), &note)

    fmt.Println(note.To)                    // display value of "to"
    fmt.Println(note.From)                  // display value from "from"
    fmt.Println(note.Subject)               // display value from "subject"
    fmt.Println(note.Body)                  // display value from "body"
}
