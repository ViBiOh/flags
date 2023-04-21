package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ViBiOh/flags"
)

// Usage of my-cli:
//              --exempleAddress  string        [exemple] Listen address {MY_CLI_EXEMPLE_ADDRESS}
//   -exempleH, --exempleHeader   string slice  [exemple] Header to add {MY_CLI_EXEMPLE_HEADER}
//   -exempleP, --exemplePort     uint          [exemple] Listen port (0 to disable) {MY_CLI_EXEMPLE_PORT}

func main() {
	fs := flag.NewFlagSet("my-cli", flag.ExitOnError)

	address := flags.String(fs, "exemple", "server", "Address", "", "Listen address", "", nil)
	port := flags.Uint(fs, "exemple", "server", "Port", "p", "Listen port (0 to disable)", 1080, nil)
	headers := flags.StringSlice(fs, "exemple", "server", "Header", "h", "Header to add", []string{"x-user", "x-auth"}, nil)

	fs.Usage = flags.Usage(fs)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("address=`%s`\n", *address)
	fmt.Printf("port=%d\n", *port)
	fmt.Printf("header=%s\n", *headers)
}
