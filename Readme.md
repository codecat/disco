# Disco
**Dis**posable **Co**ntainers for development using Docker containers.

# Install
If you have Go installed, simply run this:
```
go install github.com/codecat/disco
```

# Usage
Simply run `disco` from the command line like this:
```
$ disco
```
This will create a new temporary Alpine Docker container with the current working directory under `/src`, and a zsh shell.

From this shell you could install tools as well using `apk`:
```
$ sudo apk add npm nodejs python3
```

Of course, any changes to the container will disappear when you close out of the shell. To account for this, you can specify an alternative image type when starting `disco`.

```
$ disco js
$ npm install
```

# Image types
The following image types are currently available out of the box:

* `base`: the base image, all other images are based on this. It has some basic utilities such as `git`, `vim`, and `tmux`.
* `js`, `javascript`: image for Javascript development. This adds the `nodejs` and `npm` packages.
* `vite`: the same as `js`, but with port 5173 automatically exposed to 127.0.0.1.
* `py`, `python`: image for Python3 development. This adds the `python3` and `py3-pip` packages.
* `php`: image for PHP development. This adds the `php81` and `composer` packages.

# `disco.toml`
You may also set up a `disco.toml` file with a predefined configuration, so that you can run `disco` and automatically use the image type you need.

```toml
type = "js"
```

You can also create this file by running `disco -s` or `disco --setup`. For example, the following command will create a `disco.toml` file with the `js` image type:

```
$ disco -s js
```

# Commands
You can also pass additional commands to `disco` to run commands instead of opening a shell. For example, building a Vite project with a `disco.toml` file:

```
$ disco npx vite build
```

Or without a `disco.toml` file:

```
$ disco js npx vite build
```

# Options
There are a couple extra options you can specify either in the config file or on the command line.

* `ssh`, `--ssh`: mount a read-only volume from your own `~/.ssh` to the container's `~/.ssh`, this allows you to run `git push/pull` or `ssh` commands using your own keys.
* `zshrc`, `--zshrc`: mount a read-only volume from your own `~/.zshrc` to the container's `~/.zshrc` in case you want to use your own zsh configuration.

Passed options may also be combined with `--setup` to create a config file with those options.
