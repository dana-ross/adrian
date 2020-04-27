<img align="right" src="logo.svg" alt="Adrian" height="100" />

# Adrian: A platform for secure, performant web font hosting</h1>

[![CodeFactor](https://www.codefactor.io/repository/github/daveross/adrian/badge)](https://www.codefactor.io/repository/github/daveross/adrian)

Meet my friend Adrian! 

Adrian is a web server, like Apache or Nginx, but just for fonts. Really, just fonts. Point Adrian toward a directory of font files; it will generate CSS to use all styles and weights you have, and serve everything up for your visitors. If you’ve used Google Fonts, you should find Adrian quite familiar.

Adrian supports these font formats:

* otf
* ttf
* woff
* woff2

Do you have a project that you’ve open-sourced but you want to use commercial fonts in it? Put your code in a public repo and load your fonts off your Adrian server. 

Did you actually read the license for that font you bought, and your head is spinning after reading things like “adequate technical protection measures that restrict the Use of and/or access to the Licensed Web Fonts, for instance by utilizing JavaScript or access control mechanisms for cross-origin resource sharing”? Yeah, Adrian's got your back. 

## Getting Started

### Installing

1. Grab the most recent [Adrian release](https://github.com/daveross/adrian/releases) for your operating system. Extract the files somewhere on your server or local development environment.
1. Copy `adrian.yaml.example` to a new file named `adrian.yaml`
1. Edit `adrian.yaml` to configure Adrian
1. Run `adrian` to begin serving

To use a YAML config file in a different location, specify it with the `--config` parameter when starting Adrian. For example: `./adrian --config /etc/adrian/adrian.yaml`

Test it by loading a font CSS file, such as http://example.com/font/Arial.css (replace `example.com` with your server's hostname and `Arial.css` with the name of a font available to Adrian.

#### Adrian.yaml settings

##### domains

A whitelist of domains allowed to use fonts hosted by this instance

##### directories

A list of directories where Adrian should look for font files. On Linux, system-wide fonts are usually found in `/usr/share/fonts`.

## Usage

### http://example.com/css/?family=Arial%20Bold

Generates CSS for including the Arial Bold font file in a web project.

Get one CSS file for muliple fonts by separating the names with pipe characters. For example: http://example.com/css?family=Arial|Courier%20New

The `display` query parameter allows you to set the `font-display` style for all of the requested fonts.

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
