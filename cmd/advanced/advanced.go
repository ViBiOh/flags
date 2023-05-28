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
		url:     flags.New("url", "Database url").Shorthand("u").Prefix(prefix).DocPrefix("db").String(fs, "", overrides),
		port:    flags.New("port", "Database port").Shorthand("p").Prefix(prefix).DocPrefix("db").Uint(fs, 5432, overrides),
		name:    flags.New("name", "Database name").Shorthand("n").Prefix(prefix).DocPrefix("db").String(fs, "user", overrides),
		timeout: flags.New("yimeout", "Request timeout").Prefix(prefix).DocPrefix("db").Duration(fs, time.Second, overrides),
	}
}

// Usage of my-cli:
//   -n,        --name            string    [db] Database name ${MY_CLI_NAME} (default "user")
//   -p,        --port            uint      [db] Database port ${MY_CLI_PORT} (default 5432)
//   -replicaN, --replicaName     string    [replica] Database name ${MY_CLI_REPLICA_NAME} (default "user-replica")
//   -replicaP, --replicaPort     uint      [replica] Database port ${MY_CLI_REPLICA_PORT} (default 5432)
//   -replicaU, --replicaUrl      string    [replica] Database url ${MY_CLI_REPLICA_URL}
//              --replicaYimeout  duration  [replica] Request timeout ${MY_CLI_REPLICA_YIMEOUT} (default 1s)
//   -u,        --url             string    [db] Database url ${MY_CLI_URL}
//              --yimeout         duration  [db] Request timeout ${MY_CLI_YIMEOUT} (default 1s)

func main() {
	fs := flag.NewFlagSet("my-cli", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

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
