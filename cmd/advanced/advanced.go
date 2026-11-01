package main

import (
	"flag"
	"fmt"
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
		url:     flags.New("url", "Database url").Shorthand("u").Prefix(prefix).DocPrefix("db").Flag(fs, "", overrides),
		port:    flags.New("port", "Database port").Shorthand("p").Prefix(prefix).DocPrefix("db").Flag(fs, uint(5432), overrides),
		name:    flags.New("name", "Database name").Shorthand("n").Prefix(prefix).DocPrefix("db").Flag(fs, "user", overrides),
		timeout: flags.New("timeout", "Request timeout").Prefix(prefix).DocPrefix("db").Flag(fs, time.Second, overrides),
	}
}

// Usage of my-cli:
//   -n,        --name            string    [db] Database name ${MY_CLI_NAME} (default "user")
//   -p,        --port            uint      [db] Database port ${MY_CLI_PORT} (default 5432)
//   -replicaN, --replicaName     string    [replica] Database name ${MY_CLI_REPLICA_NAME} (default "user-replica")
//   -replicaP, --replicaPort     uint      [replica] Database port ${MY_CLI_REPLICA_PORT} (default 5432)
//              --replicaTimeout  duration  [replica] Request timeout ${MY_CLI_REPLICA_TIMEOUT} (default 1s)
//   -replicaU, --replicaUrl      string    [replica] Database url ${MY_CLI_REPLICA_URL}
//              --timeout         duration  [db] Request timeout ${MY_CLI_TIMEOUT} (default 1s)
//   -u,        --url             string    [db] Database url ${MY_CLI_URL}

func main() {
	flagSet := flag.NewFlagSet("my-cli", flag.ExitOnError)

	mainConfig := databaseFlags(flagSet, "")
	replicaConfig := databaseFlags(flagSet, "replica", flags.NewOverride("Name", "user-replica"))

	flagSet.Usage = flags.Usage(flagSet)

	_ = flagSet.Parse(os.Args[1:])

	fmt.Printf("Main name=`%s`\n", *mainConfig.name)
	fmt.Printf("Replica name=`%s`\n", *replicaConfig.name)
	fmt.Printf("Replica timeout=`%s`\n", *replicaConfig.timeout)
}
