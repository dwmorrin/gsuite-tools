// upload is a general command line tool for Google Drive
// currently a work in progress
// current goal: make a general wrapper for the service.Files.Create call.
// TODO add target folder option
// TODO add team drive option

package main

import (
    "flag"
    "fmt"
    "log"
    "github.com/dwmorrin/gsuite-tools/auth"
    "google.golang.org/api/drive/v3"
    "os"
    "os/user"
    "path/filepath"
)

// global default TODO move to config file
var defaultSecretPath = filepath.Join(".credentials", "secret", "upload.json")

func main() {
    var secretPath, folderID string
    flag.StringVar(&secretPath, "s", getDefaultSecretPath(),
        "google cloud crenditials file path",
    )
    flag.StringVar(&folderID, "f", "", "Google Drive folder ID")
    flag.Parse()
    if flag.Arg(0) == "" {
        fmt.Println("Usage:", os.Args[0], "[options] file")
        flag.PrintDefaults()
        os.Exit(1)
    }

    // Get Drive Service
    client := auth.GetClient(secretPath, drive.DriveScope)
    srv, err := drive.New(client)
    if err != nil {
        log.Fatalf("Unable to retrieve drive Client %v", err)
    }

    newData, err := os.Open(flag.Arg(0))
    if err != nil {
        log.Fatalf("Unable to open data file %v", err)
    }
    defer newData.Close()

    newFile := drive.File{
        Name: filepath.Base(newData.Name()),
    }

    if folderID != "" {
        newFile.Parents = []string{folderID}
    }

    // make the update call
    _, err = srv.Files.Create(&newFile).
        Media(newData).
        //SupportsTeamDrives(true). TODO add team drive as option
        Do()
    if err != nil {
        log.Fatalf("Unable to upload data %v", err)
    }
}

func getDefaultSecretPath() string {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    path := filepath.Join(user.HomeDir, defaultSecretPath)
    return  path
}
