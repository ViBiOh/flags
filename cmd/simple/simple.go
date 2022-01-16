package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ViBiOh/flags"
)

// Usage of my-cli:
//   -exempleAddress string
//         [exemple] Listen address {MY_CLI_EXEMPLE_ADDRESS}
//   -exemplePort uint
//         [exemple] Listen port (0 to disable) {MY_CLI_EXEMPLE_PORT} (default 1080)

func main() {
	fs := flag.NewFlagSet("my-cli", flag.ExitOnError)

	address := flags.New("exemple", "server", "Address").Default("", nil).Label("Listen address").ToString(fs)
	port := flags.New("exemple", "server", "Port").Default(1080, nil).Label("Listen port (0 to disable)").ToUint(fs)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("address=`%s`\n", *address)
	fmt.Printf("port=%d\n", *port)
}
