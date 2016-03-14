# Lint [![Build Status](https://travis-ci.org/daveym/lint.svg?branch=master)](https://travis-ci.org/daveym/lint)

## Synopsis

Lint is a terminal based application that pulls back your Pocket entries and allows you to update them, remove them or retag them. Lint is currently in develeopment and not yet at a releasable stage, but feel free to have a look. At the moment, I'm using this as a vehicle for learning Go as well, so any advice or pointers are greatly appreciated.

## Intro

To use this app, you need to have a [Pocket account](https://getpocket.com) and have access to your 'Consumer Key'. A consumer key created for each application that wants to access Pocket. As I'm open sourcing this application, feel free to go and set up your own app on Pocket to get a consumer key

## Motivation

I frequently find myself saving stuff to Pocket and never coming back to it again, or forgetting what I've yet to read. I also frequently save things without applying any tags whatsoever. This makes it tricky to find what I'm looking for later. I'm also attempting to learn Go, so I figured a good way to learn would be to create a simple test app that connects to the [Pocket API](https://getpocket.com/developer/docs/overview) and pulls back my Pocket items. I can then retag, delete etc as appropriate.

## Installation

If you want to have a go, download the source and get the latest [Go binary](https://golang.org/dl/). You'll need to compile, using 'go build lint.go' and get setup by executing './lint auth'.

Pocket requires that you authenticate using your COnsumer Token, and then you have to permit lint to access the pocket data via an authorisation webpage. When you do that, you get an access token that can be used for adding,. retrieving or modifying your pocket data. You'll only need to authenticate once with Lint, as it persists your access code to a local lint config (yaml) file.

## API Reference

Lint provides a Pocket Client, which is basically an endpoint that can be used to make requests into Pockets API. At the moment, authentication is present. I'll then be working on retrieving, adding and modifying documents.


## License

Licenses under MIT.

