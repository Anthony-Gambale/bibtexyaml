# BibTeX-YAML

Write your bibliography in YAML and convert it to BibTeX. CLI built using [Cobra](https://github.com/spf13/cobra).

### Installation

First, [install Go](https://go.dev/doc/install).

Next, install the current version of BibTeX-YAML

`go install github.com/Anthony-Gambale/BibTeX-YAML@v1.1.0`

You'll also need to add your go binaries to the PATH

`export PATH=$PATH:$HOME/go/bin`

Don't forget to add this to `~/.bashrc` to run it every time you open a new shell.

### Usage

Start by putting your new BibTeX references into a file called `new.bib`. For example

```BibTeX
@conference{davis2018,
  booktitle = {Proceedings of the Example Conference},
  year      = {2018},
  author    = {Davis, Bob},
  title     = {A Conference Paper}
}
```

Then use `BibTeX-YAML reverse <name>.yaml` to convert the entries to YAML and *append* them to the end of your current bibliography in `<name>.yaml`. If you don't have this file yet, it will be created. It may look something like the following

```yaml
entries:
  - id: davis2018
    type: conference
    fields:
#     booktitle: Proceedings of the Example Conference
      year: 2018
      author: Davis, Bob
      title: A Conference Paper
```

Note that once you do this, the entries in `new.bib` will be cleared. Next, you can comment out fields you aren't sure if you want to use, as has been done with `booktitle` above. Use `BibTeX-YAML template <name>.yaml` to convert your entire YAML bibliography back into BibTeX. This will overwrite your `<name>.bib` file, or create it if it does not exist. This may look something like the following

```BibTeX
@conference{davis2018,
  author = {Davis, Bob},
  title  = {A Conference Paper},
  year   = {2018}
}
```

This is the general workflow of `BibTeX-YAML`.

See `new.bib`, `test.yaml` and `test.bib` for a more comprehensive example.

### Todo (contribution wanted)

- [x] Templating: Compile YAML bibliography into BibTeX
- [x] Reverse templating: convert BibTeX entries into YAML
- [ ] Quality of life changes
  - [x] Stop printing templating output to terminal
  - [x] Add curly braces to every field
  - [ ] Better error handling (some errors are being ignored)
  - [ ] Better YAML parsing errors (some information is missing e.g. column number)
  - [ ] Breaking code into smaller pieces, readability
