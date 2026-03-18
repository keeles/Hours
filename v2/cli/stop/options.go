package stop

type Options struct {
	Client string `arg:"" help:"Client that timer belongs to"`
	Task   string `arg:"" optional:"" help:"Name of task"`
}
