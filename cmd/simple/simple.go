package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ViBiOh/flags"
)

// Usage of my-cli:
//       --address  string        [server] Listen address {MY_CLI_ADDRESS}
//   -h, --header   string slice  [server] Header to add {MY_CLI_HEADER}, as a string slice, environment variable separated by `,` (default [x-user, x-auth])
//   -p, --port     uint          [server] Listen port (0 to disable) {MY_CLI_PORT} (default 1080)

func main() {
	fs := flag.NewFlagSet("my-cli", flag.ExitOnError)

	address := flags.New("address", "Listen address").DocPrefix("server").String(fs, "", nil)
	port := flags.New("port", "Listen port (0 to disable)").Shorthand("p").DocPrefix("server").Uint(fs, 1080, nil)
	headers := flags.New("header", "Header to add").Shorthand("h").DocPrefix("server").StringSlice(fs, []string{"x-user", "x-auth"}, nil)

	fs.Usage = flags.Usage(fs)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("address=`%s`\n", *address)
	fmt.Printf("port=%d\n", *port)
	fmt.Printf("header=%s\n", *headers)
}
