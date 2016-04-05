# Lint [![Build Status](https://travis-ci.org/daveym/lint.svg?branch=master)](https://travis-ci.org/daveym/lint) <a href="http://goreportcard.com/report/daveym/lint"><img src="https://camo.githubusercontent.com/897657567000bd843315798076b1b9f097bc5bb4/687474703a2f2f676f7265706f7274636172642e636f6d2f62616467652f626f627a697563686b6f76736b692f637565" alt="Report Card" data-canonical-src="http://goreportcard.com/badge/daveym/lint" style="max-width:100%;"></a> 

## Synopsis

Lint is a terminal based application that pulls back your Pocket entries and allows you to archive, favourite, unfavourite readd (from archived state) or delete these entries. Lint is currently in development and has not yet reached a releasable milestone, but feel free to have a look. At the moment, I'm using this as a vehicle for learning Go as well, so any advice or guidance is greatly appreciated as well.

## Intro

To use this app, you need to have a [Pocket account](https://getpocket.com) and have access to your 'Consumer Key'. A consumer key is created by Pocket for each application that wants to access Pocket data via the Pocket API. As I'm open sourcing this application, feel free to go and set up your own copy on Pocket to get a consumer key - you'll need one to run this app! 


## Installation

If you want to have a go, download the source and get the latest [Go binary](https://golang.org/dl/). You'll need to compile, using 'go build lint.go' and get setup by executing './lint auth'.

Pocket requires that you authenticate using your Consumer Token, and then you have to permit Lint to access the pocket data via an authorisation webpage. When you do that, you get an access token returned back from Pocket that can is used in subsequent API requests to add, retrieve or modify your pocket data. You'll only need to authenticate once with Lint, as it persists your access code to a local lint config (yaml) file. The following gives you an idea on how to get lint setup:

``` 
1) ./lint auth
Please make sure that lint.yaml exists in the working directory.

2) touch lint.yaml
vim lint.yaml

3) paste in the following:
consumerKey: your-consumer-key

4) re-run ./lint auth - you'll then be presented with an oAuth page for authentication.

5) At this point you ca retrieve articles. They'll come back in the form:
ID, Title, URL 

If you want, you can pipe these to a file using ./lint retrieve -s all >> myfile. From there, see what articles you want to keep or archive or favourite and use their ID's as inputs to lint using:

./lint modify -a item1 item2 item3 etc.
```

## Motivation

I frequently find myself saving stuff to Pocket and never coming back to it again, or forgetting what I've yet to read. I also frequently save things without applying any tags whatsoever. This makes it tricky to find what I'm looking for later. I want an app that can help surface what I've stored in Pocket, but also something that I can use to manage items I've saved.

I'm also attempting to learn Go, so I figured a good way to learn would be to create a simple app that connects to the [Pocket API](https://getpocket.com/developer/docs/overview) and pulls back my Pocket items. I can then retag, delete etc as appropriate.

## API Reference

The app is constructed with the following components:

1. [Cobra](https://github.com/spf13/cobra) - Command line processing
2. [Viper](https://github.com/spf13/viper) - Configuration file access
3. Pocket Package - Wrapper aroound the [Pocket API](https://getpocket.com/developer/docs/overview)

## Getting Involved
If you want to be involved in this, drop me a note on twitter, at @djmcglade


## License

Licensed under MIT.

