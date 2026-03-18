package remove

type Options struct {
	Name            string  `arg:"" help:"Name of client"`
	Task            string  `arg:"" help:"Name of task to remove hours from"`
	MinutesToRemove float32 `arg:"" help:"Number of minutes to remove from project, must be integer value"`
}
