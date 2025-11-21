package delete

type Options struct {
	Name string `arg:"" aliases:"d" help:"Delete a project from the database"`
	Force bool `short:"f" help:"Force deletion of project without confirmation"`
}