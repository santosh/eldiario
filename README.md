# eldiario

[![first-timers-only](https://img.shields.io/badge/first--timers--only-friendly-blue.svg?style=flat-square)](https://www.firsttimersonly.com/)

<!-- Add build and coverage tags -->

**eldiario** is a command-line utility to write down notes, in the form of a journal.
eldiario is based on the client-server model. eldiario's backend is written in go
which handles API requests. A web interface which operates on the API.

<!-- INSERT SCREENSHOT/GIF HERE -->

## Features

As a development point of view, this project features...

- Unit tests responsible in CI/CD pipeline for deployment.

## Getting Started

### Installation

    go get -u github.com/santosh/eldiario

To run a local version of the server.

    docker-compose up

This will spin up the eldiario server with mongodb dependency.

If you have done the modification to the code, you can force rebuild of the image:

    docker-compose up --build

### Usage

Two frontends are planned for development.

- A web version. Head over to <http://127.0.0.1:8080/>
- A command-line version.

<!-- TODO: When you launch eldiario, you are put inside your choice of editor (eldiario reads `$EDITOR`). -->

## Architecture

<!-- Document the design here -->

### Endpoints

You want to use the server using APIs and create your own frontends. Like the one we have made [eldiario-cli](https://github.com/santosh/eldiario-cli).

<!-- Hook it to listing website, something known like awesome-go -->

## Contributing

To get started with the development, fork the repo, look if you can do something with.

### Documentation

Check <> or run godoc server locally.

### Testing

While I formulate the text to write here, head over to [CONTRIBUTING.md](https://github.com/santosh/eldiario/blob/master/CONTRIBUTING.md) to get an orientation of the project.

## FAQ

## Contributors

Add yourself here:

- Santosh Kumar (twitter, github)

## License

See [LICENSE](https://github.com/santosh/eldiario/blob/master/LICENSE).
