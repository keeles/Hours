package main

import (
	"github.com/keeles/hours/cli/add"
	"github.com/keeles/hours/cli/complete"
	"github.com/keeles/hours/cli/config"
	"github.com/keeles/hours/cli/delete"
	"github.com/keeles/hours/cli/get"
	"github.com/keeles/hours/cli/list"
	"github.com/keeles/hours/cli/new"
	"github.com/keeles/hours/cli/remove"
	"github.com/keeles/hours/cli/start"
	"github.com/keeles/hours/cli/stop"
	"github.com/keeles/hours/cli/task"
	"github.com/keeles/hours/cli/time"
	"github.com/keeles/hours/cli/version"
)

type Hours struct {
	Version  version.Options  `cmd:"" aliases:"v" help:"Print Version Number"`
	Config   config.Options   `cmd:"" aliases:"conf" help:"Commands for configurations"`
	New      new.Options      `cmd:"" aliases:"n" help:"Add a new client for time tracking"`
	Start    start.Options    `cmd:"" aliases:"s" help:"Start a timer for working task for client - hours start <client-name> <task-name>"`
	Stop     stop.Options     `cmd:"" aliases:"s" help:"Stop the timer for working task - hours stop <client-name> <task-name>"`
	Task     task.Options     `cmd:"" aliases:"t" help:"Add new task to client - hours task <client-name> <task-name>"`
	Time     time.Options     `cmd:"" aliases:"t" help:"Show the active timer"`
	Get      get.Options      `cmd:"" aliases:"g" help:"Get the tasks for a given client name"`
	Add      add.Options      `cmd:"" aliases:"a" help:"Add time to an existing project, default is minutes | Use flag --hours to record hours"`
	Remove   remove.Options   `cmd:"" aliases:"rm" help:"Remove time from an existing project, default is minutes | Use flag --hours to record hours"`
	List     list.Options     `cmd:"" aliases:"ls" help:"List all projects with hour tracking"`
	Delete   delete.Options   `cmd:"" aliases:"d" help:"Delete a client and ALL tasks - hours delete <client-name>"`
	Complete complete.Options `cmd:"" help:"Complete task and delete it from the database - hours complete <client-name> <task>"`
}
