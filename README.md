# Project template

Simple web server template with pre-configured MySQL and Redis using stdlib.

This code was derived, modified from (https://github.com/fiorix)

It has the basics for http and logging thru a middleware, config file, Redis and MySQL.

Copy and modify to suit your project. 

Handlers and http middleware are setup inside http.go and handlers.go as a struct with methods.


## Preparing the environment

Prerequisites:

- Git
- GNU Make
- [Go](http://golang.org) 1.6 or newer

First, you should make a copy of this directory, and prepare the new project:

	cp -r golang_http_api_skel foobar
	cd foobar
	./bootstrap.sh

Your project is now called **foobar** and is ready to use.

Make sure the Go compiler, `$GOPATH` is set and the repo is on your src dir.

Install dependencies, and compile:

	make deps
	make clean
	make all

Optional: Generate a self-signed SSL certificate (optional):

	cd ssl
	make

Optional: Start Redis and set up MySQL:

	sudo mysql < assets/files/database.sql

Edit the config file and run the server

	vi foobar.conf
	./foobar

Optional: Install, uninstall. Edit Makefile and set PREFIX to the target directory:

	sudo make install
	sudo make uninstall

Optional: Allow non-root process to listen on low ports:

	/sbin/setcap 'cap_net_bind_service=+ep' /opt/foobar/server

