module github.com/nitin0409sep/Go

go 1.25.6

// Dependencies -> Just Like - Package.json
require (
	github.com/fatih/color v1.18.0
	github.com/mattn/go-colorable v0.1.13 // indirect; indirect --> Indirect -> U are not using them in your project till now
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)

//! go mod tidy -> In case u r using any dependecy and its still in indirect -> This command will fix it
