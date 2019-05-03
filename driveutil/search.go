/* custom wrappers for Drive search functions */
package driveutil

import (
    "google.golang.org/api/drive/v3"
)

func Query(srv *drive.Service, q string) (foundFile *drive.File, err error) {
    var pageToken string

    for searching := true; searching; {
        pageOfFiles, err := srv.Files.List().
            SupportsTeamDrives(true).Q(q).PageToken(pageToken).Do()
        if err != nil {
            return nil, err
        }

        pageToken = pageOfFiles.NextPageToken
        for _, file := range pageOfFiles.Files {
            if file.Trashed {
                continue
            }
            foundFile = file
            searching = false
            break
        }
        if pageToken == "" {
            searching = false
        }
    }
    return foundFile, nil
}

