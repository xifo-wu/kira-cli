package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"xifo.in/kira/app"
)

var CONFIG_PATH = "$HOME/.kira"

// var CONFIG_PATH = "/Users/xifo/Code/Projects/kira-cli"

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(CONFIG_PATH)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	var filename string
	var path string
	autoCmd := flag.NewFlagSet("auto", flag.ExitOnError)
	autoCmd.StringVar(&filename, "filename", "", "filename")
	autoCmd.StringVar(&path, "path", "", "path")

	// fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// fooEnable := fooCmd.Bool("enable", false, "enable")
	// fooName := fooCmd.String("name", "", "name")

	// barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	// barLevel := barCmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'auto' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "auto":
		autoCmd.Parse(os.Args[2:])

		app.Auto(path, filename)
	default:
		fmt.Println("expected 'auto' subcommands")
		os.Exit(1)
	}
}
