package cmd

import "github.com/blang/semver"

func SetSemverMake(f func(s string) (semver.Version, error)) {
	semverMake = f
}

func ResetSemverMake() {
	semverMake = semver.Make
}

func SetSemverCompare(f func(v1, v2 semver.Version) int) {
	semverCompare = f
}

func ResetSemverCompare() {
	semverCompare = func(v1, v2 semver.Version) int {
		return v1.Compare(v2)
	}
}
