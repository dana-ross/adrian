# Adrian: A platform for secure, performant web font hosting.

Adrian ([Frutiger](https://en.wikipedia.org/wiki/Adrian_Frutiger), not Pennino) is a server for hosting web fonts. It scans all the fonts in a directory and automatically generates CSS with @font-face declarations for individual fonts or families. Caching headers are added to all responses so browsers know not to request the same files repeatedly, while CORS and filename obfuscation helps comply with security restrictions in some fonts' licenses.

## Getting Started

### Prerequisites

Adrian is written in [Node.js](https://nodejs.org/en/), and you will need Node.js installed to run it.

### Installing

1. Clone the repository and `cd` into the directory
1. Edit `adrian.yaml` to configure
1. `node app/src/server.js` to start the server

Test it by loading a font CSS file, such as http://example.com/font/Arial.css (replace `example.com` with your server's hostname and `Arial.css` with the name of a font available to Adrian.

#### Configuring
##### domains
A whitelist of domains allowed to use fonts hosted by this instance

##### directories
A list of directories where Adrian should look for font files. On Linux, system-wide fonts are stored in `/usr/share/fonts`.

Supported font formats:
* otf
* ttf
* woff
* woff2

## Built With

* [Express](https://expressjs.com/) - The web framework used
* [Fontkit](https://github.com/devongovett/fontkit) - Font file parsing

## Contributing

Please consider opening a [Pull Request](https://github.com/daveross/adrian/pulls) to submit changes to this project.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/daveross/adrian/tags). 

## Authors

* **Dave Ross** - *Initial work* - [daveross](https://github.com/daveross)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
