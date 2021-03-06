<img align="right" src="logo.svg" alt="Adrian" height="100" />

# Adrian: A platform for hosting web fonts</h1>

[![CodeFactor](https://www.codefactor.io/repository/github/dana-ross/adrian/badge)](https://www.codefactor.io/repository/github/dana-ross/adrian) ![GitHub](https://img.shields.io/github/license/dana-ross/adrian) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/dana-ross/adrian)

Meet my friend Adrian! 

Adrian is a web server, like Apache or Nginx, but just for fonts. Really, just fonts. Point Adrian toward a directory of font files; it'll generate CSS to use all styles and weights you have, and serve everything up for your visitors. If you’ve used Google Fonts, you should find Adrian quite familiar.

Adrian supports these font formats:

* otf
* ttf
* woff
* woff2

Do you have a project that you’ve open-sourced but you can't put the commercial fonts you bought in your repo? Put your code up on GitHub and load your fonts off your private Adrian server. 

Did you actually read the license for that font you bought, and your head is spinning after reading things like “adequate technical protection measures that restrict the Use of and/or access to the Licensed Web Fonts, for instance by utilizing JavaScript or access control mechanisms for cross-origin resource sharing”? Yeah, Adrian's got your back. 

## Getting Started

### Installing

1. Grab the most recent [Adrian release](https://github.com/dana-ross/adrian/releases) for your operating system. Extract the files somewhere on your server or local development environment.
1. Copy `adrian.yaml.example` to a new file named `adrian.yaml`
1. Edit `adrian.yaml` to configure Adrian
1. Run `adrian` to begin serving

To use a YAML config file in a different location, specify it with the `--config` parameter when starting Adrian. For example: `./adrian --config /etc/adrian/adrian.yaml`

Test it by loading a font CSS file, such as http://example.com/font/Arial.css (replace `example.com` with your server's hostname and `Arial` with the name of a font available to Adrian.

#### Configuring Adrian with `adrian.yaml`

```yaml
global:

  # Port number Adrian responds to
  port: 80
  
  # Adrian will only allow fonts to be used on these URLs (CORS functionality)
  domains:
    - example.com
    
  # Directories where Adrian should look for fonts
  directories:
    - /usr/share/fonts
    
  # If true, replace font filenames with hashes so they can't be guessed as easily
  obfuscate filenames: false

  # Used to set the cache-control header in responses
  cache-control lifetime: 2628000

  # Paths for writing logs to disk
  logs:
    access: "/var/log/adrian/access.log"
```

##### port: &lt;integer&gt;

The TCP/IP port Adrian will listen to. Defaults to port 80.

##### domains

A whitelist of domains allowed to use fonts hosted by this instance

##### directories

A list of directories where Adrian should look for font files. On Linux, system-wide fonts are usually found in `/usr/share/fonts`.

##### obfuscate filenames: &lt;boolean&gt;

If true, the filenames of font files are replaced with hashes so they can't be guessed as easily.

##### cache-control lifetime: &lt;seconds&gt;

Used to set the cache-control header sent to browsers and CDNs. This header instructs everyone downstream to cache Adrian's CSS and font files for this amount of time.

##### logs

###### access: &lt;string&gt;
Path where Adrian should write an access log. Access logs use Common Log Format for easy parsing.

## Usage

### CSS Import

In your site's CSS, import Adrian's CSS for the fonts you want to use:
```
@import "https://example.com/css?family=Arian|Times+New+Roman;
```

### URL formats

#### http://example.com/css/?family=Arial

Generates CSS for including the Arial font and all of its variants in a web project.

Get one CSS file for muliple fonts by separating the names with pipe characters. For example: http://example.com/css?family=Arial|Courier%20New

#### http://example.com/css/?family=Arial&display=swap

The `display` query parameter allows you to set the [`font-display`](https://developer.mozilla.org/en-US/docs/Web/CSS/@font-face/font-display) style for all of the requested fonts. For example, the `display=swap` value tells browsers to render text with fallback fonts until custom ones are downloaded.

## Built With

* [Go](https://golang.org/)
* [Echo](https://echo.labstack.com/) - The web framework used
* [Fastcache](https://github.com/VictoriaMetrics/fastcache) - In-memory caching library
* [sfnt](https://github.com/ConradIrwin/font/tree/master/sfnt) - Font file parsing

## Contributing

Please consider opening a [Pull Request](https://github.com/dana-ross/adrian/pulls) to submit changes to this project.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/dana-ross/adrian/tags). 

## Authors

* **Dana Ross** - *Initial work* - [dana-ross](https://github.com/dana-ross)

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/dana-ross/adrian/blob/main/LICENSE) file for details
