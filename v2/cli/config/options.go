package config

type Options struct {
	AddDirectory    AddDirectoryOptions    `cmd:"" help:"Associate current directory with a client"`
	RemoveDirectory RemoveDirectoryOptions `cmd:"" help:"Remove current directory association"`
	List            ListOptions            `cmd:"" help:"List configured directories"`
}

type AddDirectoryOptions struct {
	Client string `arg:"" help:"Client name to associate with this directory"`
}

type RemoveDirectoryOptions struct {
	Client string `arg:"" optional:"" help:"Remove directory association with client"`
}

type ListOptions struct{}
