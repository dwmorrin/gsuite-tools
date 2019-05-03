package main

import (
    "flag"
    "fmt"
    "log"
    "github.com/dwmorrin/gsuite-tools/auth"
    "github.com/dwmorrin/gsuite-tools/driveutil"
    "google.golang.org/api/drive/v3"
    "os"
    "os/user"
    "path/filepath"
)

// global default TODO move to config file
var defaultSecretPath = filepath.Join(".credentials", "secret", "upload.json")

func getDefaultSecretPath() string {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    path := filepath.Join(user.HomeDir, defaultSecretPath)
    return  path
}

func usage() {
    fmt.Println("Usage:", os.Args[0], "download|search|upload [options]")
    flag.PrintDefaults()
    os.Exit(1)
}

func main() {
    var (
        action, fileID, parentID, outpath, secretPath string
        update bool
    )
    flag.StringVar(&action, "a", "", "action: download|search|upload")
    flag.StringVar(&fileID, "f", "", "Google Drive file ID")
    flag.StringVar(&outpath, "o", "", "output file path")
    flag.StringVar(&parentID, "p", "", "Google Drive parent folder ID")
    flag.StringVar(&secretPath, "s", getDefaultSecretPath(),
        "google cloud crenditials file path",
    )
    flag.BoolVar(&update, "u", false, "update Google Drive file by ID")
    flag.Parse()

    // Get Drive Service
    client := auth.GetClient(secretPath, drive.DriveScope)
    srv, err := drive.New(client)
    if err != nil {
        log.Fatalf("Unable to retrieve drive Client %v", err)
    }
    if flag.Arg(0) == "" || action == "" {
        usage()
    }
    file := flag.Arg(0)
    if action == "upload" {
        driveutil.Upload(srv, fileID, parentID, file)
    }
    if action == "download" {
        if outpath == "" {
            outpath, err = os.Getwd()
            if err != nil {
                log.Fatalf("Unable to determine current directory: %v", err)
            }
        }
        driveutil.Download(srv, file, parentID, outpath)
    }
    if action == "search" {
        //q := `'` + parentID + `' in parents and name contains '` + file + `'`
        q := `name contains '` + file + `'`
        result, err := driveutil.Query(srv, q)
        if err != nil {
            log.Fatalf("Unable to complete query %v", err)
        }
        fmt.Printf("Result: %v\n", result.Name)
    }
    os.Exit(0)
}
