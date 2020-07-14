# eldiario

[![first-timers-only](https://img.shields.io/badge/first--timers--only-friendly-blue.svg?style=flat-square)](https://www.firsttimersonly.com/)

**eldiario** is a command-line utility to write down notes, in the form of a journal.
eldiario is based on the client-server model. eldiario's backend is written in go
which handles API requests. A web interface which operates on the API.

<!-- INSERT SCREENSHOT/GIF HERE -->

## Getting Started

### Installation

For local installation and testing, run.

To run a local instance of eldiario, run `docker-compose up`. This will spin up 
eldiario server with mongodb dependency.

### Usage

Two frontends are planned for development.

- A web version. Head over to <http://127.0.0.1:8080/>
- A command-line version.

<!-- TODO: When you launch eldiario, you are put inside your choice of editor (eldiario reads `$EDITOR`). -->

<!-- Hook it to listing website, something known like awesome-go -->

## Contributing

While I formulate the text to write here, head over to [CONTRIBUTING.md](https://github.com/santosh/eldiario/blob/master/CONTRIBUTING.md) to get an orientation of the project.

## License

See [LICENSE](https://github.com/santosh/eldiario/blob/master/LICENSE).
