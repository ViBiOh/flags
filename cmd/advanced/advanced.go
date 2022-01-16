package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ViBiOh/flags"
)

type databaseConfig struct {
	url  *string
	port *uint
	name *string
}

func databaseFlags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) databaseConfig {
	return databaseConfig{
		url:  flags.New(prefix, "db", "URL").Default("", overrides).Label("Database url").ToString(fs),
		port: flags.New(prefix, "db", "Port").Default(5432, overrides).Label("Database port").ToUint(fs),
		name: flags.New(prefix, "db", "Name").Default("user", overrides).Label("Database name").ToString(fs),
	}
}

// Usage of my-cli:
//   -mainName string
//         [main] Database name {MY_CLI_MAIN_NAME} (default "user")
//   -mainPort uint
//         [main] Database port {MY_CLI_MAIN_PORT} (default 5432)
//   -mainURL string
//         [main] Database url {MY_CLI_MAIN_URL}
//   -replicaName string
//         [replica] Database name {MY_CLI_REPLICA_NAME} (default "user-replica")
//   -replicaPort uint
//         [replica] Database port {MY_CLI_REPLICA_PORT} (default 5432)
//   -replicaURL string
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
}
