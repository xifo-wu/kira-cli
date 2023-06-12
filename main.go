package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var CONFIG_PATH = "$HOME/.kira"

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
	savePath := viper.GetString("save_path")
	rclonePath := viper.GetString("rclone_path")
	log.Println(savePath)
	log.Println(rclonePath)

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

		l := filepath.Join(path, filename)
		newRclonePath := strings.Replace(l, savePath, rclonePath, 1)
		fmt.Println(path, savePath, rclonePath)

		fmt.Println(newRclonePath, "newRclonePath")

		fmt.Println(l)
		cmd := exec.Command("rclone", "moveto", "-v", l, newRclonePath)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

	// case "foo":
	// 	fooCmd.Parse(os.Args[2:])
	// 	fmt.Println("subcommand 'foo'")
	// 	fmt.Println("  enable:", *fooEnable)
	// 	fmt.Println("  name:", *fooName)
	// 	fmt.Println("  tail:", fooCmd.Args())
	// case "bar":
	// 	barCmd.Parse(os.Args[2:])
	// 	fmt.Println("subcommand 'bar'")
	// 	fmt.Println("  level:", *barLevel)
	// 	fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'auto' subcommands")
		os.Exit(1)
	}
}
