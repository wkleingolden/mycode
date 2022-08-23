/* RZFeeser | Alta3 Research
   Create a struct type named Userdata
   to track user information */


package main

import (
    "fmt"
    "os/user"
    "os"
    "os/exec"
)

type Userdata struct {
    na string    // user id
    un string    // username
    hd string    // home directory
    em string    // email
}

func main() {

    cmd := exec.Command("clear") // Clear the term screen before program run
    cmd.Stdout = os.Stdout
    cmd.Run()

    // user.Current() is from import os/user - get info about "current system user"
    user, err := user.Current()      // returns user and error
    if err != nil {
         panic(err)
    }

    // create a struct named currentUser
    currentUser := Userdata{user.Uid, user.Username, user.HomeDir, "rzfeeser@example.com"}

    // recall currentUser
    fmt.Println("User id:", currentUser.na)        // recall current user id
    fmt.Println("Username:", currentUser.un)       // recall username
    fmt.Println("Home Directory:", currentUser.hd) // recall homedirectory
    fmt.Println("Email:", currentUser.em)          // recall email
}
