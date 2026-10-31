package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ViBiOh/flags"
)

// Usage of my-cli:
//       --address  string        [server] Listen address {MY_CLI_ADDRESS}
//   -h, --header   string slice  [server] Header to add {MY_CLI_HEADER}, as a string slice, environment variable separated by `,` (default [x-user, x-auth])
//   -p, --port     uint          [server] Listen port (0 to disable) {MY_CLI_PORT} (default 1080)

func main() {
	flagSet := flag.NewFlagSet("my-cli", flag.ExitOnError)
	flagSet.Usage = flags.Usage(flagSet)

	address := flags.New("address", "Listen address").DocPrefix("server").String(flagSet, "", nil)
	port := flags.New("port", "Listen port (0 to disable)").Shorthand("p").DocPrefix("server").Uint(flagSet, 1080, nil)
	headers := flags.New("header", "Header to add").Shorthand("h").DocPrefix("server").StringSlice(flagSet, []string{"x-user", "x-auth"}, nil)

	flagSet.Usage = flags.Usage(flagSet)

	_ = flagSet.Parse(os.Args[1:])

	fmt.Printf("address=`%s`\n", *address)
	fmt.Printf("port=%d\n", *port)
	fmt.Printf("header=%s\n", *headers)
}
