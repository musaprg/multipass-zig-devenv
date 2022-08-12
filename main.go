package main

import (
	"flag"
	"fmt"
	"os"
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
	flag.BoolVar(&dryrun, "dry-run", false, "only generating cloud-init.yaml without launching actual multipass VM")

	flag.Usage = printHelp
}

func launchVM() {

}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}

	switch args[0] {
	case "launch":
		launchVM()
	default:
		flag.Usage()
	}
}
