/*
Package workflow provides utilities for building workflows for Alfred 2.
https://www.alfredapp.com/

You need Alfred's Powerpack to be able to use workflows.

NOTE: This software is very alpha and not even half-finished.

To read this documentation on godoc.org, see
http://godoc.org/gogs.deanishe.net/deanishe/awgo.git

This library provides an API for communicating with Alfred and several
convenience methods for common workflow tasks.


Features

The current main features are:

	- Easy access to Alfred context, such as data and cache directories.
	- Simple generation of Alfred XML feedback.
	- Fuzzy filtering.
	- Catches panics, logs stack trace and shows user an error message.
	- Log file for easier debugging.


Usage

Typically, you'd call your program's main entry point via Run(). This
way, the library will rescue any panic, log the stack trace and show
an error to the user.

program.go:

	package main

	import "gogs.deanishe.net/deanishe/awgo"

	func run() {
		// Your workflow starts here
	}

	func main() {
		// Package is called workflow
		workflow.Run(run)
	}

In the Script Filter's Script box (Language = /bin/bash):

	./program "{query}"


Fuzzy filtering

SortFuzzy() implements Alfred-like fuzzy search, e.g. "of" will match
"OmniFocus" and "got" will match "Game of Thrones".

To use SortFuzzy, your struct must implement the Fuzzy interface, which
is sort.Interface plus a Keywords() method that returns the string
the fuzzy filtering should be applied to.

*/
package workflow
