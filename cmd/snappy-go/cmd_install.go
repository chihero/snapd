package main

import (
	"os"

	"launchpad.net/snappy/helpers"
	"launchpad.net/snappy/snappy"
)

type cmdInstall struct {
}

func init() {
	var cmdInstallData cmdInstall
	cmd, _ := parser.AddCommand("install",
		"Install a snap package",
		"Install a snap package",
		&cmdInstallData)

	cmd.Aliases = append(cmd.Aliases, "in")
}

func (x *cmdInstall) Execute(args []string) (err error) {
	var lock *helpers.FileLock

	if lock, err = helpers.StartPrivileged(); err != nil {
		return err
	}
	defer func() { err = helpers.StopPrivileged(lock) }()

	err = snappy.Install(args)
	if err != nil {
		return err
	}
	// call show versions afterwards
	installed, err := snappy.ListInstalled()
	if err != nil {
		return err
	}

	showInstalledList(installed, os.Stdout)

	return err
}
