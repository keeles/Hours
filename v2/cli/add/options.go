package add

type Options struct {
	Name   string  `arg:"" help:"Name of client"`
	Task   string  `arg:"" help:"Name of the task to add time to"`
	Amount float32 `arg:"" help:"Number of minutes to add to project, must be integer value"`

	Hours bool `short:"H" help:"Interpret amount as hours instead of minutes"`
}
