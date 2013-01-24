package main 

import (
    "flag"
    "fmt"
)
import (
    "github.com/jessevdk/go-flags"
)
import (
    "github.com/dcoupal/examples-go/lib/goext"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	var opts struct {
		// Slice of bool will append 'true' each time the option
		// is encountered (can be set multiple times, like -vvv)
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

		// Example of automatic marshalling to desired type (uint)
		Offset uint `long:"offset" description:"Offset"`
	}

	// Parse flags
	flags.Parse(&opts)

    if len(opts.Verbose) > 0 {
        fmt.Println("Verbose is on")
    }
    goext.ArraysAreEqualStrings([]string{"Ruth"}, []string{"Ruth"})
}

