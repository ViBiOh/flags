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
//   -n,        --name            string    [db] Database name {MY_CLI_NAME}
//   -p,        --port            uint      [db] Database port {MY_CLI_PORT}
//   -replicaN, --replicaName     string    [replica] Database name {MY_CLI_REPLICA_NAME}
//   -replicaP, --replicaPort     uint      [replica] Database port {MY_CLI_REPLICA_PORT}
//              --replicaTimeout  duration  [replica] Request timeout {MY_CLI_REPLICA_TIMEOUT}
//   -replicaU, --replicaUrl      string    [replica] Database url {MY_CLI_REPLICA_URL}
//              --timeout         duration  [db] Request timeout {MY_CLI_TIMEOUT}
//   -u,        --url             string    [db] Database url {MY_CLI_URL}

func main() {
	fs := flag.NewFlagSet("my-cli", flag.ExitOnError)

	mainConfig := databaseFlags(fs, "")
	replicaConfig := databaseFlags(fs, "replica", flags.NewOverride("Name", "user-replica"))

	fs.Usage = flags.Usage(fs)

	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Main name=`%s`\n", *mainConfig.name)
	fmt.Printf("Replica name=`%s`\n", *replicaConfig.name)
	fmt.Printf("Timeout name=`%s`\n", *replicaConfig.timeout)
}
