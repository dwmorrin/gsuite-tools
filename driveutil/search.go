/*Package driveutil contains custom wrappers for Drive search functions */
package driveutil

import (
    "google.golang.org/api/drive/v3"
)

/*Query is a wrapper for searching Google Drive */
func Query(srv *drive.Service, driveID string, q string) (foundFile *drive.File, err error) {
    var pageToken string

    for searching := true; searching; {
        filesCall := srv.Files.List();
        if driveID != "" {
            filesCall = filesCall.
              IncludeItemsFromAllDrives(true).
              SupportsAllDrives(true).
              Corpora("teamDrive").
              DriveId(driveID)
        }
        pageOfFiles, err := filesCall.Q(q).PageToken(pageToken).Do()
        if err != nil {
            return nil, err
        }

        pageToken = pageOfFiles.NextPageToken
        for _, file := range pageOfFiles.Files {
            if file.Trashed {
                continue
            }
            return file, nil
        }
        if pageToken == "" {
            searching = false
        }
    }
    return nil, nil
}

