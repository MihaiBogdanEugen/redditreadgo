redditreadgo
=========

The goal of this project is to provide an efficient way of retrieving Reddit submissions of a specified subreddit or of a redditor.

[![Build Status](https://travis-ci.org/MihaiBogdanEugen/redditreadgo.svg?branch=master)](https://travis-ci.org/MihaiBogdanEugen/redditreadgo) [![Go Report Card](https://goreportcard.com/badge/github.com/MihaiBogdanEugen/redditreadgo)](https://goreportcard.com/report/github.com/MihaiBogdanEugen/redditreadgo) [![GoDoc Widget]][GoDoc]

[GoDoc]: https://godoc.org/github.com/MihaiBogdanEugen/redditreadgo
[GoDoc Widget]: https://godoc.org/github.com/MihaiBogdanEugen/redditreadgo?status.svg

Installing
----------
Run

    go get github.com/MihaiBogdanEugen/redditreadgo

Include in your source:

    import "github.com/MihaiBogdanEugen/redditreadgo"
    
Using dep:

    dep ensure -add github.com/MihaiBogdanEugen/redditreadgo

Godoc
-----
See http://godoc.org/github.com/MihaiBogdanEugen/redditreadgo

Configuration
-----

One can configure a `ReadOnlyRedditClient` by providing the following:
- the Reddit API Consumer Key,
- the Reddit API Consumer Secret and
- an user agent used for preventing banning or API throttling
