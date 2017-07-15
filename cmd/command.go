package cmd

type Command interface {
	Execute(args []string) error
}
