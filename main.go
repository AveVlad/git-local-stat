package main

import (
	"os"
	"github.com/jessevdk/go-flags"
	"github.com/avevlad/git-local-stat/server"
	"fmt"
	"time"
)

var options server.CommandLineOptions
var parser = flags.NewParser(&options, flags.Default)

func main() {
	start := time.Now()

	if _, err := parser.Parse(); err != nil {
		os.Exit(1)
	}
	server.Init(options)
	bodyCommits := server.GetBodyCommits()
	//_ = server.SetBody(&server.Commits[0], bodyCommits[server.Commits[0].CommitId])

	hh := server.SetBodies(server.Commits, bodyCommits)
	xd := time.Now()
	fmt.Println("BEFORE X", time.Since(start))
	rex := hh.DateRangeFilter("", "", true, []server.Predicate{
		func(commit server.GitCommitResponse) bool {
			return commit.Author.Name == "osxfcn"
		},
		func(commit server.GitCommitResponse) bool {
			return commit.CommitID == "b22cda08bdc28963a249bd4b4f1a6da246c067e2"
		}}...)
	fmt.Println(len(rex))
	fmt.Println(hh[0].Author.Date)
	fmt.Println("AFTER X", time.Since(xd))
	fmt.Println("//////////\\///313/23////d/f/vc/c/vc/xv/cx/vd/fas/df/ds/fa")
}