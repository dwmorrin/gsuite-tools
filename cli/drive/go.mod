module github.com/dwmorrin/gsuite-tools/cli/drive

go 1.16

replace github.com/dwmorrin/gsuite-tools/driveutil => ../../driveutil

replace github.com/dwmorrin/gsuite-tools/auth => ../../auth

require (
	github.com/dwmorrin/gsuite-tools/auth v0.0.0-00010101000000-000000000000
	github.com/dwmorrin/gsuite-tools/driveutil v0.0.0-00010101000000-000000000000
	google.golang.org/api v0.47.0
)
