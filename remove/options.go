package remove

type Options struct {
	Name string `arg:"" help:"Name of project to remove hours from"`
	HoursToRemove int `arg:"" help:"Number of hours to remove from project, must be integer value"`
}