# eldiario

[![first-timers-only](https://img.shields.io/badge/first--timers--only-friendly-blue.svg?style=flat-square)](https://www.firsttimersonly.com/)

<!-- Add build and coverage tags -->

**eldiario** is a command-line utility to write down notes, in the form of a journal.
eldiario is based on the client-server model. eldiario's backend is written in go
which handles API requests. A web interface which operates on the API.

<!-- INSERT SCREENSHOT/GIF HERE -->

<!-- ## Features

As a development point of view, this project features...

- Unit tests responsible in CI/CD pipeline for deployment. -->

## Getting Started

### Installation

    git clone github.com/santosh/eldiario

To run a local version of the server.

    docker-compose up

This will spin up the eldiario server with mongodb dependency.

If you have done the modification to the code, you can force rebuild of the image:

    docker-compose up --build

### Usage

Two frontends are planned for development.

- A web version. <https://github.com/santosh/eldiario-web>
- A command-line version. <https://github.com/santosh/eldiario-cli>

For now, REST endpoints are the only way to access the API. See <https://github.com/santosh/eldiario/wiki/API-Tests>.

<!-- Hook it to listing website, something known like awesome-go -->

## Contributing

To get started with the development, head over <https://github.com/santosh/eldiario/issues>, look for something you are familiar with, fork the repo.

While I formulate the text to write here, head over to [CONTRIBUTING.md](https://github.com/santosh/eldiario/blob/master/CONTRIBUTING.md) to get an orientation of the project.

## TODO

[ ] - Authentication
[ ] - Autherization for endpoint access

## Contributors

Add yourself here:

- Santosh Kumar (twitter, github)

## License

See [LICENSE](https://github.com/santosh/eldiario/blob/master/LICENSE).
