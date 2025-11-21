package main

import (
	"github.com/keeles/hours/add"
	"github.com/keeles/hours/get"
	"github.com/keeles/hours/list"
	"github.com/keeles/hours/new"
	"github.com/keeles/hours/remove"
	"github.com/keeles/hours/version"
)

type Hours struct {
	Version version.Options `cmd:"" short:"v" help:"Print Version Number"`
	New new.Options `cmd:"" short:"n" help:"Add a new project for hour tracking"`
	Get get.Options `cmd:"" short:"g" help:"Get the hours for a given project name"`
	Add add.Options `cmd:"" short:"a" help:"Add hours to an existing project - hours add <project> <hours>"`
	Remove remove.Options `cmd:"" aliases:"rm" help:"Remove hours from an existing project - hours remove <project> <hours>"`
	List list.Options `cmd:"" aliases:"ls" help:"List all projects with hour tracking"`
}
