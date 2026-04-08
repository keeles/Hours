package config

type Options struct {
	AddDirectory    AddDirectoryOptions    `cmd:"" aliases:"add-dir" help:"Associate current directory with a client - hours config add-directory <client-name>"`
	RemoveDirectory RemoveDirectoryOptions `cmd:"" aliases:"rm-dir" help:"Remove current directory association - hours config remove-directory <client-name>"`
	List            ListOptions            `cmd:"" aliases:"ls" help:"List configured directories"`
	Completion      CompletionOptions      `cmd:"" aliases:"comp" help:"Get shell completion scripts"`
}

type AddDirectoryOptions struct {
	Client string `arg:"" help:"Client name to associate with this directory"`
}

type RemoveDirectoryOptions struct {
	Client string `arg:"" optional:"" help:"Remove directory association with client"`
}

type ListOptions struct{}

type CompletionOptions struct {
	Shell string `arg:"" help:"Shell to get completion script for - options: Bash, Zsh, fish"`
}
