package main

import kingpin "gopkg.in/alecthomas/kingpin.v2"

var (
	cli = kingpin.New("eos-actions", "firehose of eos actions")

	filterAccountNameExact = cli.Flag("account", "filter action by account name").String()
	filterActionNameExact  = cli.Flag("action", "filter action by action name").String()
)
