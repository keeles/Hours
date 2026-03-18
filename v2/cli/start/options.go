package start

type Options struct {
	Client string `arg:"" optional:"" help:"Client that timer belongs to"`
	Task   string `arg:"" optional:"" help:"Name of task"`
}
