// gmail is a command line interface to the gmail api
// currently just sends a message
package main

import (
    "encoding/base64"
    "flag"
    "fmt"
    "log"
    "github.com/dwmorrin/gsuite-tools/auth"
    "google.golang.org/api/gmail/v1"
    "os"
    "os/user"
    "path/filepath"
)

// global default TODO move to config file
var defaultSecretPath = filepath.Join(".credentials", "secret", "gmail.json")

func getDefaultSecretPath() string {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    path := filepath.Join(user.HomeDir, defaultSecretPath)
    return path
}

func usage() {
    fmt.Println("Usage:", os.Args[0], "-f from -t to -s subject -b body")
    flag.PrintDefaults()
    os.Exit(1)
}

func main() {
    var (
        body, from, secretPath, subject, to string
    )
    flag.StringVar(&body, "b", "", "body of message")
    flag.StringVar(&from, "f", "me", "@gmail.com address")
    flag.StringVar(&subject, "s", "", "subject line")
    flag.StringVar(&to, "t", "", "recipients")
    flag.StringVar(&secretPath, "p", getDefaultSecretPath(),
        "google cloud crenditials file path",
    )
    flag.Parse()

    if body == "" || to == "" || subject == "" {
        usage()
    }

    message := []byte("From: " + from + "\r\n" +
             "To: " + to + "\r\n" +
             "Subject: " + subject + "\r\n\r\n" +
             body)

    var email gmail.Message
    email.Raw = base64.StdEncoding.EncodeToString(message)

    client := auth.GetClient(secretPath, gmail.MailGoogleComScope)
    srv, err := gmail.New(client)
    _, err = srv.Users.Messages.Send("me", &email).Do()
    if err != nil {
        log.Fatalf("Message failed to send: %v", err)
    }
}
