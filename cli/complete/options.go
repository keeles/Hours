package complete

type Options struct {
	Client string `arg:"" help:"Name of client"`
	Task   string `arg:"" help:"Name of task to complete (Removes task from database)"`
}
