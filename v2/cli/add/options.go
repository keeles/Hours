package add

type Options struct {
	Name       string  `arg:"" help:"Name of client"`
	Task       string  `arg:"" help:"Name of the task to add hours to"`
	NewMinutes float32 `arg:"" help:"Number of minutes to add to project, must be integer value"`
}
