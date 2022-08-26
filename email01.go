/* RZFeeser | Alta3 Research
   Sending an Email (SMTP) message   */
   
package main

import (
    "log"
    "net/smtp"
)

func main() {

    // Configuration
    from := "surferdude789@gmail.com"            // update this to reflect your value
    password := "********************"     // update this to reflect your value
    to := []string{"surferdude789@gmail.com"}   // update this to reflect your value
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"

    // update this to be your message body 
    message := []byte("From: surferdude789@gmail.com\r\n" + "To: surferdude789@gmail.com\r\n" + "Subject: Testing email from go\r\n\n" + "My super secret message.")

    // Create authentication
    auth := smtp.PlainAuth("", from, password, smtpHost)

    // Send actual message
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
    if err != nil {
        log.Fatal(err)
    }
}

