package cmd

import (
	"io"
)

type Help struct {
	stdout io.Writer
	stderr io.Writer
}

func NewHelp(stdout, stderr io.Writer) Help {
	return Help{
		stdout: stdout,
		stderr: stderr,
	}
}

func (c Help) Execute(args []string) error {
	c.stderr.Write([]byte("Usage: semver [COMMAND]\n"))
	return nil
}
