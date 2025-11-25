package main

import (
	"github.com/keeles/hours/add"
	"github.com/keeles/hours/complete"
	"github.com/keeles/hours/delete"
	"github.com/keeles/hours/get"
	"github.com/keeles/hours/list"
	"github.com/keeles/hours/new"
	"github.com/keeles/hours/remove"
	"github.com/keeles/hours/task"
	"github.com/keeles/hours/version"
)

type Hours struct {
	Version  version.Options  `cmd:"" aliases:"v" help:"Print Version Number"`
	New      new.Options      `cmd:"" aliases:"n" help:"Add a new client for hour tracking"`
	Task     task.Options     `cmd:"" aliases:"t" help:"Add new task to client - hours task <client-name> <task-name>"`
	Get      get.Options      `cmd:"" aliases:"g" help:"Get the tasks for a given client name"`
	Add      add.Options      `cmd:"" aliases:"a" help:"Add hours to an existing project - hours add <client-name> <task> <hours>"`
	Remove   remove.Options   `cmd:"" aliases:"rm" help:"Remove hours from an existing project - hours remove <client-name> <task> <hours>"`
	List     list.Options     `cmd:"" aliases:"ls" help:"List all projects with hour tracking"`
	Delete   delete.Options   `cmd:"" aliases:"d" help:"Delete a client and ALL tasks - hours delete <client-name>"`
	Complete complete.Options `cmd:"" help:"Complete task and delete it from the database - hours complete <client-name> <task>"`
}
