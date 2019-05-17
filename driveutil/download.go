package driveutil

import (
    "io/ioutil"
    "google.golang.org/api/drive/v3"
    "log"
    "path/filepath"
)

/*Download is a wrapper for service.Files.Get().Download()*/
func Download(srv *drive.Service, filename, folderID, driveID, outpath string) {
    // query string
    queryString := `name contains '` + filename + `'`
    if folderID != "" {
        queryString = `'` + folderID + `' in parents and ` + queryString
    }

    file, err := Query(srv, driveID, queryString)
    if err != nil {
        log.Fatalf("search error, check folder Id: %v", err)
    }
    if file == nil {
        log.Fatalf("file %v not found in %v", filename, folderID)
    }

    // Download the file
    res, err := srv.Files.Get(file.Id).Download()
    if err != nil {
        log.Fatalf("Error with download %v", err)
    }
    defer res.Body.Close()

    // Write the file to disk
    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        log.Fatalf("Error reading download %v", err)
    }

    outpath = filepath.Join(outpath, file.Name)

    err = ioutil.WriteFile(outpath, body, 0600)
    if err != nil {
        log.Fatalf("Error writing download %v", err)
    }
}
