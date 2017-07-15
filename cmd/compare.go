package cmd

import (
	"fmt"
	"io"

	"github.com/blang/semver"
)

type Compare struct {
	stdout io.Writer
	stderr io.Writer
}

var semverMake = semver.Make
var semverCompare = func(v1, v2 semver.Version) int {
	return v1.Compare(v2)
}

func NewCompare(stdout, stderr io.Writer) Compare {
	return Compare{
		stdout: stdout,
		stderr: stderr,
	}
}

func (c Compare) Execute(args []string) error {
	v1, err := semverMake(args[0])
	if err != nil {
		panic(err)
	}

	v2, err := semverMake(args[1])
	if err != nil {
		panic(err)
	}

	c.stdout.Write([]byte(fmt.Sprintf("%d\n", semverCompare(v1, v2))))
	return nil
}
