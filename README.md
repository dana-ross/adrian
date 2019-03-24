<img align="left" src="logo.svg" alt="Adrian" width="200">

# Adrian: A platform for secure, performant web font hosting</h1>
<div style="clear:both;"></div>
NOTE: Adrian is under active development & major architectural changes and is not ready for use.

Adrian ([Frutiger](https://en.wikipedia.org/wiki/Adrian_Frutiger), not Pennino) is a server for hosting web fonts. It scans all the fonts in a directory and automatically generates CSS with @font-face declarations for individual fonts or families. Caching headers are added to all responses so browsers know not to request the same files repeatedly, while CORS and filename obfuscation helps comply with security restrictions in some fonts' licenses.

Basically, it lets you serve fonts for your sites from a central location so they're not sitting in your project's repo, which is important if you have commercial fonts on an otherwise open source site. It also uses browser same-origin security and filename obfuscation to help prevent you from being a one-stop distribution site for the fonts you paid good money for.

## Getting Started

### Prerequisites

Adrian is written in [Go](https://golang.org/).

### Installing

1. Clone the repository and `cd` into the directory
1. `npm install`
1. Edit `adrian.yaml` to configure
1. `node app/src/server.js` to start the server

Test it by loading a font CSS file, such as http://example.com/font/Arial.css (replace `example.com` with your server's hostname and `Arial.css` with the name of a font available to Adrian.

### docker-compose

When running Adrian through `docker-compose`, don't forget to mount a `fonts` directory containing the fonts you want served, and your own configuration file. For example:

```yaml
version: '3'
services:
  adrian:
    build: 'https://github.com/daveross/adrian.git'
    volumes:
      - './fonts:/usr/share/fonts'
      - './adrian.yaml:/usr/src/app/adrian.yaml'
    ports:
      - '3000:3000'
```

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

## Usage

### http://example.com/css/?family=Arial%20Bold

Generates CSS for including the Arial Bold font file in a web project.

Get one CSS file for muliple fonts by separating the names with pipe characters. For example: http://example.com/css?family=Arial|Courier%20New

## Built With

* [Go](https://golang.org/)
* [Echo](https://echo.labstack.com/) - The web framework used
* [sfnt](https://github.com/ConradIrwin/font/tree/master/sfnt) - Font file parsing

## Contributing

Please consider opening a [Pull Request](https://github.com/daveross/adrian/pulls) to submit changes to this project.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/daveross/adrian/tags). 

## Authors

* **Dave Ross** - *Initial work* - [daveross](https://github.com/daveross)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
