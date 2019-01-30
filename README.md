# ghp
A simple web server for serving static GitHub Pages locally, to test before
deploying.

This can be useful compared to browsing local HTML files from your browser when
you use absolute paths in links, such as `/about`, `/js/app.js`,
`/css/style.css`, etc., which won't resolve correctly in the context of your
filesystem.

It is also handy compared to something like `python -m http.server` which
doesn't support dropping the file extension, e.g. `/about` rather than
`/about.html`.

When requesting any path (`$path`), `ghp` will do the following (all file
operations are relative to the `root` commandline flag):
1. Check whether `$path` points to a file, if so serve that file
1. Check whether `$path` points to a directory, if so serve `$path/index.html`
2. Check whether `$path.html` points to a file, if so serve that file
3. Check whether `404.html` is a file, if so serve that file as a 404
4. Serve a 404

## Usage
```
$ go get github.com/CurtisLusmore/ghp
$ ghp -help
Usage of ghp:
  -port int
        The port to serve over (default 8080)
  -root string
        The root directory to serve files from (your GitHub Pages repo) (default ".")
$ ghp -root MyGitHubPages
```

## Notes
As this tool exposes your filesystem to your network, you should be careful
using this on untrusted networks.

## Todo
* Add support for serving rendered Markdown files (`.md`)