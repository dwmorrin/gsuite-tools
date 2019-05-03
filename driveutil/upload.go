// upload is a general command line tool for Google Drive
// currently a work in progress
// current goal: make a general wrapper for the service.Files.Create call.
// TODO add team drive option

package driveutil

import (
    "log"
    "google.golang.org/api/drive/v3"
    "os"
    "path/filepath"
)

func Upload(srv *drive.Service, fileID, folderID, file string) {
    newData, err := os.Open(file)
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

    if fileID == "" {
        _, err = srv.Files.Create(&newFile).
            Media(newData).
            //SupportsTeamDrives(true). TODO add team drive as option
            Do()
        if err != nil {
            log.Fatalf("Unable to upload data %v", err)
        }
    } else {
        _, err = srv.Files.Update(fileID, &newFile).
            Media(newData).
            Do()
        if err != nil {
            log.Fatalf("Unable to update file %v", err)
        }
    }
}
