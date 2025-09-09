# Disco
**Dis**posable **Co**ntainers for development using Docker containers.

# Install
If you have Go installed, simply run this:
```
$ go install github.com/codecat/disco@latest
```
You might also have to add `$HOME/go/bin` to your `$PATH` variable, for example:

```sh
export PATH=$HOME/go/bin:$PATH
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
* `js`, `javascript`: image for Javascript development. This adds the `nodejs`, `npm`, and `pnpm` packages.
* `vite`: the same as `js`, but with port 5173 automatically exposed to 127.0.0.1.
* `py`, `python`: image for Python3 development. This adds the `python3` and `py3-pip` packages.
* `php`: image for PHP development. This adds the `php81` and `composer` packages.
* `php-framework`: image for PHP development using frameworks, based on `php`. This adds a number of commonly required PHP modules.
* `php-rr`: image for PHP development using Roadrunner, based on `php-framework`. Has port 8080 automatically exposed to 127.0.0.1.

# `disco.toml`
You may also set up a `disco.toml` file with a predefined configuration, so that you can run `disco` and automatically use the image type you need.

```toml
type = "js"
```

You can also create this file by running `disco -i` or `disco --init`. For example, the following command will create a `disco.toml` file with the `js` image type:

```
$ disco -i js
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
* `fish`, `--fish`: mount a read-only volume from your own `~/.config/fish` to the container's `~/.config/fish` in case you want to use your own fish configuration.

Passed options may also be combined with `--setup` to create a config file with those options.

# Building
## Disco
To build Disco and make it available for execution:
```
$ go install
```
