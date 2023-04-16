package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ViBiOh/flags"
)

type databaseConfig struct {
	url     *string
	port    *uint
	name    *string
	timeout *time.Duration
}

func databaseFlags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) databaseConfig {
	return databaseConfig{
		url:     flags.String(fs, prefix, "db", "Url", "u", "Database url", "", overrides),
		port:    flags.Uint(fs, prefix, "db", "Port", "p", "Database port", 5432, overrides),
		name:    flags.String(fs, prefix, "db", "Name", "n", "Database name", "user", overrides),
		timeout: flags.Duration(fs, prefix, "db", "Timeout", "", "Request timeout", time.Second, overrides),
	}
}

// Usage of my-cli:
//   -mainN string
//         Shorthand for -mainName (default "user")
//   -mainName string
//         [main] Database name {MY_CLI_MAIN_NAME} (default "user")
//   -mainP uint
//         Shorthand for -mainPort (default 5432)
//   -mainPort uint
//         [main] Database port {MY_CLI_MAIN_PORT} (default 5432)
//   -mainTimeout duration
//         [main] Request timeout {MY_CLI_MAIN_TIMEOUT} (default 1s)
//   -mainU string
//         Shorthand for -mainUrl
//   -mainUrl string
//         [main] Database url {MY_CLI_MAIN_URL}
//   -replicaN string
//         Shorthand for -replicaName (default "user-replica")
//   -replicaName string
//         [replica] Database name {MY_CLI_REPLICA_NAME} (default "user-replica")
//   -replicaP uint
//         Shorthand for -replicaPort (default 5432)
//   -replicaPort uint
//         [replica] Database port {MY_CLI_REPLICA_PORT} (default 5432)
//   -replicaTimeout duration
//         [replica] Request timeout {MY_CLI_REPLICA_TIMEOUT} (default 1s)
//   -replicaU string
//         Shorthand for -replicaUrl
//   -replicaUrl string
//         [replica] Database url {MY_CLI_REPLICA_URL}

func main() {
	fs := flag.NewFlagSet("my-cli", flag.ExitOnError)

	mainConfig := databaseFlags(fs, "main")
	replicaConfig := databaseFlags(fs, "replica", flags.NewOverride("Name", "user-replica"))

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Main name=`%s`\n", *mainConfig.name)
	fmt.Printf("Replica name=`%s`\n", *replicaConfig.name)
	fmt.Printf("Timeout name=`%s`\n", *replicaConfig.timeout)
}
