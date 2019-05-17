// upload is a general command line tool for Google Drive
// currently a work in progress
// current goal: make a general wrapper for the service.Files.Create call.

package driveutil

import (
    "log"
    "google.golang.org/api/drive/v3"
    "os"
    "path/filepath"
)

/*Upload is a wrapper for service.Files.{Create,Update}().Do()*/
func Upload(srv *drive.Service, fileID, folderID, file string) {
    newData, err := os.Open(file)
    if err != nil {
        log.Fatalf("Unable to open data file %v", err)
    }
    defer newData.Close()

    newFile := drive.File{}
    if folderID != "" {
        newFile.Parents = []string{folderID}
    }

    if fileID == "" {
        newFile.Name = filepath.Base(newData.Name())
        createCall := srv.Files.Create(&newFile).SupportsAllDrives(true)
        _, err = createCall.Media(newData).Do()
        if err != nil {
            log.Fatalf("Unable to upload data %v", err)
        }
    } else {
        updateCall := srv.Files.Update(fileID, &newFile).SupportsAllDrives(true)
        _, err = updateCall.Media(newData).Do()
        if err != nil {
            log.Fatalf("Unable to update file %v", err)
        }
    }
}
