package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/dwmorrin/gsuite-tools/auth"
	"github.com/dwmorrin/gsuite-tools/driveutil"
	"google.golang.org/api/drive/v3"
)

// global default TODO move to config file
var defaultSecretPath = filepath.Join(".credentials", "secret", "drive.json")

func getDefaultSecretPath() string {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    path := filepath.Join(user.HomeDir, defaultSecretPath)
    return  path
}

func usage() {
    fmt.Println("Usage:", os.Args[0], "-a download|search|upload [options] file")
    flag.PrintDefaults()
    os.Exit(1)
}

func main() {
    var (
        action, fileID, parentID, outpath, teamDriveID, secretPath string
        //update bool
    )
    flag.StringVar(&action, "a", "", "action: download|search|upload")
    flag.StringVar(&fileID, "f", "", "Google Drive file ID")
    flag.StringVar(&outpath, "o", "", "output file path")
    flag.StringVar(&parentID, "p", "", "Google Drive parent folder ID")
    flag.StringVar(&secretPath, "s", getDefaultSecretPath(),
        "google cloud credentials file path",
    )
    flag.StringVar(&teamDriveID, "t", "", "Team Drive ID")
    //flag.BoolVar(&update, "u", false, "update Google Drive file by ID")
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
        driveutil.Download(srv, file, parentID, teamDriveID, outpath)
    }
    if action == "search" {
        //q := `'` + parentID + `' in parents and name contains '` + file + `'`
        q := `name contains '` + file + `'`
        result, err := driveutil.Query(srv, teamDriveID, q)
        if err != nil {
            log.Fatalf("Unable to complete query %v", err)
        }
        if result != nil {
            fmt.Printf("Result: %v\n", result.Name)
        } else {
            fmt.Println("No results for", file)
        }
    }
    os.Exit(0)
}
