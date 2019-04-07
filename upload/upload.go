// upload is a general command line tool for Google Drive
// currently a work in progress
// current goal: make a very general wrapper for the service.Files.Create call.
// worked OK to upload a .zip in root personal Drive
// TODO first test uploaded a file named "untitled" - match current file name to Drive title
// TODO add ability to target folder (should allow team drive folders as well)
// TODO look for or write a helper function or helper package to automatically determine
//        the appropriate mime type based on the given local file path given. (currently hardcoded to .zip)
package main

import (
    "fmt"
    "log"
    "github.com/dwmorrin/gsuite-tools/auth"
    "google.golang.org/api/drive/v3"
    "os"
)

// command line arguments
const (
    programName = iota
    secretPath
    sourceData
)
const minArgLength = sourceData + 1

func main() {
    if len(os.Args) != minArgLength {
        fmt.Printf("Usage: %v secretPath sourceDataPath (not yet:) targetSheetId" +
            "\n  secretPath        google cloud crenditials file path" +
            "\n  sourceDataPath    file you want to upload (hardcoded for zip)" +
            "\n  (not implemented yet) targetFolderId    target parent folder\n",
            os.Args[programName])
        os.Exit(1)
    }

    // Get Drive Service
    client := auth.GetClient(os.Args[secretPath], drive.DriveScope)
    srv, err := drive.New(client)
    if err != nil {
        log.Fatalf("Unable to retrieve drive Client %v", err)
    }

    newData, err := os.Open(os.Args[sourceData])
    if err != nil {
        log.Fatalf("Unable to open data file %v", err)
    }
    defer newData.Close()

    newFile := drive.File{}
    newFile.MimeType = `application/zip` // TODO determine type based on user input

    // make the update call
    _, err = srv.Files.Create(&newFile).
        Media(newData).
        //SupportsTeamDrives(true). TODO add team drive as option
        Do()
    if err != nil {
        log.Fatalf("Unable to upload data %v", err)
    }
}
