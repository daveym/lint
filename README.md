# Lint [![Build Status](https://travis-ci.org/daveym/lint.svg?branch=master)](https://travis-ci.org/daveym/lint)

## Synopsis

Lint is a terminal based application that pulls back your Pocket entries and allows you to add, update, remove or retag these entries. Lint is currently in development and not yet at a releasable stage, but feel free to have a look. At the moment, I'm using this as a vehicle for learning Go as well, so any advice or pointers are greatly appreciated.

## Intro

To use this app, you need to have a [Pocket account](https://getpocket.com) and have access to your 'Consumer Key'. A consumer key is created by Pocket for each application that wants to access Pocket data via the Pocket API. As I'm open sourcing this application, feel free to go and set up your own copy on Pocket to get a consumer key

## Motivation

I frequently find myself saving stuff to Pocket and never coming back to it again, or forgetting what I've yet to read. I also frequently save things without applying any tags whatsoever. This makes it tricky to find what I'm looking for later. 

I'm also attempting to learn Go, so I figured a good way to learn would be to create a simple app that connects to the [Pocket API](https://getpocket.com/developer/docs/overview) and pulls back my Pocket items. I can then retag, delete etc as appropriate.

## Installation

If you want to have a go, download the source and get the latest [Go binary](https://golang.org/dl/). You'll need to compile, using 'go build lint.go' and get setup by executing './lint auth'.

Pocket requires that you authenticate using your Consumer Token, and then you have to permit Lint to access the pocket data via an authorisation webpage. When you do that, you get an access token returned back from Pocket that can is used in subsequent API requests to add, retrieve or modify your pocket data. You'll only need to authenticate once with Lint, as it persists your access code to a local lint config (yaml) file.

## API Reference

Lint provides a Pocket Client, which is basically an endpoint that can be used to make requests into Pockets API. At the moment, authentication is present. I'll then be working on retrieving, adding and modifying items via the Pocket API.

## License

Licensed under MIT.

