package get

type Options struct {
	Name string `arg:"" help:"Name of the project"`

	Task bool `short:"t" help:"Return only task names not time"`
}

