package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	sshPublicKeyPath  = "~/.ssh/multipass.pub"
	sshPrivateKeyPath = "~/.ssh/multipass"
	zlsVersion        = "0.9.0"
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
	// TODO(musaprg): Use crypto package instead of directly executing OpenSSH command
	if _, err := exec.Command("ssh-keygen", "-t", "rsa", "-b", "4096", "-f", sshPublicKeyPath).Output(); err != nil {
		return err
	}
	pubKeyContent, err := os.ReadFile(sshPublicKeyPath)
	if err != nil {
		return err
	}
	cc := cloudConfig{
		AuthorizedKey: string(pubKeyContent),
		ZLSVersion:    zlsVersion,
	}
	buf := bytes.NewBufferString("")
	cc.printAsYAML(buf)
	cmd := exec.Command("multipass", "--name", name, "--cpus", fmt.Sprintf("%d", cpus), "--mem", mem, "--disk", disk, "--cloud-config", "-", image)
	cmd.Stdin = buf
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_, err = cmd.Output()
	return err
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}

	cc := cloudConfig{
		AuthorizedKey: "<modify here with your public key value>",
		ZLSVersion:    zlsVersion,
	}

	switch args[0] {
	case "launch":
		log.Printf("name: %s\n", name)
		log.Printf("cpus: %d\n", cpus)
		log.Printf("mem: %s\n", mem)
		log.Printf("disk: %s\n", disk)
		log.Printf("image: %s\n", image)
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
