package main

import (
	"flag"
	"fmt"
	"os"
    "log"
)

var (
	name   string
	cpus   int
	mem    string
	disk   string
	image  string
	dryrun bool
)

func printHelp() {
	fmt.Fprintf(os.Stderr, `Generating multipass-based development environment for Zig project

Usage: %s

Available Commands:
  launch        launch multipass VM
  gen           generate cloud-config.yaml based on passed value

Options:
`, os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.StringVar(&name, "name", "multipass", "name of devenv, which is used for determining multipass VM name")
	flag.IntVar(&cpus, "cpus", 2, "number of vCPU for multipass VM")
	flag.StringVar(&mem, "mem", "4G", "amount of memory for multipass VM")
	flag.StringVar(&disk, "disk", "20G", "amount of disk for multipass VM")
	flag.StringVar(&image, "image", "latest", "ubuntu image used for launching multipass VM")

	flag.Usage = printHelp
}


func launchVM() error {
    // TODO(musaprg): generate ssh key
    // TODO(musaprg): add to cloud config authorized key setting
    return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}

    cc := cloudConfig{
        AuthorizedKey: "hogehoge",
        ZLSVersion: "0.9.0",
    }

	switch args[0] {
	case "launch":
        err := launchVM()
        if err != nil {
            log.Fatalln(err)
        }
    case "gen":
        err := cc.printAsYAML(os.Stdout)
        if err != nil {
            log.Fatalln(err)
        }
	default:
		flag.Usage()
	}
}
