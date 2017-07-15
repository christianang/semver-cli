# semver-cli
Small cli to compare semvers written in go

# Install

```
go get github.com/christianang/semver-cli/semver
```

# Usage

Comparing two semvers:

```
$ semver compare 1.2.3 2.3.4
-1
```

`semver compare`:
- returns a `0` if the two semvers are equal
- returns a `1` if the first semver is greater than the second semver
- returns a `-1` if the first semver is less than the second semver
