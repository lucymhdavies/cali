# Cali - Cali Automation Layout Initialiser

![Cali logo](docs/cali.png)

Cali is a library which lets you create CLIs to run containers, without the hassle of having to alias long `docker run` commands.
It also has added bonus of being able to modify how the container is run programatically, so your users don't need to worry.

By default, it binds your current working directory when running a container. But you can also tell it to clone a git repository and run from that.

We're still figuring out how the open source version of Cali should work, and it needs quite a bit of refactoring before it's ready for a v1.

Feel free to play with it, but while we're still on v0.x.x, be sure to pin to specific versions in your dependency manager of choice.

## Docs

The API [docs](https://godoc.org/github.com/skybet/cali) are available on GoDoc

For a basic example, see [examples/basic](https://github.com/skybet/cali/blob/master/examples/basic)

Take a look (and help edit) [the wiki](https://github.com/skybet/cali/wiki) for more usage examples
