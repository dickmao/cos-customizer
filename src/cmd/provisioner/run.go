// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"
	"flag"
	"log"

	"github.com/google/subcommands"
)

// Run implements subcommands.Command for the "run" command.
// This command runs the provisioner from a provided configuration file.
type Run struct {
	configPath string
}

// Name implements subcommands.Command.Name.
func (r *Run) Name() string {
	return "run"
}

// Synopsis implements subcommands.Command.Synopsis.
func (r *Run) Synopsis() string {
	return "Provision a COS instance from the provided configuration file."
}

// Usage implements subcommands.Command.Usage.
func (r *Run) Usage() string {
	return `run [flags]
`
}

// SetFlags implements subcommands.Command.SetFlags.
func (r *Run) SetFlags(f *flag.FlagSet) {
	f.StringVar(&r.configPath, "config", "", "Path to a configuration file to use for provisioning.")
}

func (r *Run) validate() error {
	if r.configPath == "" {
		return errors.New("-config must be provided")
	}
	return nil
}

// Execute implements subcommands.Command.Execute.
func (r *Run) Execute(_ context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if err := r.validate(); err != nil {
		log.Printf("Error in flags: %v", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}