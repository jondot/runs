# Runs

A simple implementation of clustering to groups with shared distance properties, or "runs".

Use this if you want to:

* Group files by [their modification date](examples/filesbydate.go). For example, you can group photos
  into shooting sessions, by specifying that two photos belong in a session if
  they are not more than one hour apart.

* Group files by [their size](examples/filesbysize.go). For example, group same-length TV shows and movies by
  defining that two videos belong in a group if they are not more than 50MB
  apart, without caring too much about the various different length
  ahead-of-time.

## Quick Start

```
$ git clone https://github.com/jondot/runs
$ cd runs
$ make examples
```

Look at [examples](examples/). You can now run the examples on any folder you
like (they will date current "." folder always).



# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2014 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.



