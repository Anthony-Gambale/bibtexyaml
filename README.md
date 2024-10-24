# bibtexyaml

Write your bibliography in YAML and convert it to BibTeX.

## Usage

Put one or more BibTeX references into a file called `new.bib`. For example

```BibTeX
@conference{davis2018,
  booktitle = {Proceedings of the Example Conference},
  year      = {2018},
  author    = {Davis, Bob},
  title     = {A Conference Paper}
}
```

Then use `bibtexyaml reverse <name>.yaml` to convert the entries to YAML and *append* them to the end of your current bibliography in `<name>.yaml`. If it doesn't exist, it will be created automatically. It may look something like this

```yaml
entries:
  - id: davis2018
    type: conference
    fields:
      # booktitle: Proceedings of the Example Conference
      year: 2018
      author: Davis, Bob
      title: A Conference Paper
```

Note that once you do this, `new.bib` will be cleared. Next, you can comment out fields as has been done with `booktitle` above.

Use `bibtexyaml template <name>.yaml` to convert your YAML bibliography back into BibTeX. This will overwrite your `<name>.bib` file, or create it if it does not exist. In this case it will produce the following

```BibTeX
@conference{davis2018,
  author = {Davis, Bob},
  title  = {A Conference Paper},
  year   = {2018}
}
```

The unwanted field `booktitle` is now removed, and this can easily be undone by uncommenting the line in `<name>.yaml`. This is the general workflow of `bibtexyaml`.

See `new.bib`, `test.yaml` and `test.bib` for a more comprehensive example.

## Installation

First, [install Go](https://go.dev/doc/install).

Next, install the current version of bibtexyaml

`go install github.com/anthonygam/bibtexyaml@v1.2.4`

You'll also need to add your go binaries to the PATH

`export PATH=$PATH:$HOME/go/bin`

Don't forget to add this to `~/.bashrc` to run it every time you open a new shell.

Built using [Cobra](https://github.com/spf13/cobra).
