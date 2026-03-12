package task

type Options struct {
	Client string `arg:"" help:"Client that new task belongs to"`
	Task   string `arg:"" help:"Name of task"`
}
