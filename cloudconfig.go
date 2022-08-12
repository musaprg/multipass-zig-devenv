package main

import (
    "io"
	"fmt"
    "text/template"
)

// cloudConfig represents configurations used for generating cloud-config.yaml
type cloudConfig struct {
    AuthorizedKey string // AuthroizedKey is the authorized SSH public key used for the VM user
    ZLSVersion string // ZLSVersion is the semvar-formatted version string for ZLS installed in the VM
}

func (cc *cloudConfig) printAsYAML(w io.Writer) error {
    tmpl := template.Must(template.ParseFiles("templates/cloud-config.yaml.tmpl"))
    if err := tmpl.Execute(w, cc); err != nil {
        return fmt.Errorf("unable to generate cloud config as yaml: %w+", err)
    }
    return nil

}

