package add

type Options struct {
	Name     string `arg:"" help:"Name of client"`
	Task     string `arg:"" help:"Name of the task to add hours to"`
	NewHours int    `arg:"" help:"Number of hours to add to project, must be integer value"`
}
