# BibTeX-YAML

Write your bibliography in YAML and convert it to BibTeX. CLI built using [Cobra](https://github.com/spf13/cobra).

### Features

- [x] (Templating) Compile YAML bibliography into BibTeX
- [ ] (Reverse-templating) Convert BibTeX entries into YAML

### Installation and Usage

First, [install Go](https://go.dev/doc/install).

Next, install the current version of BibTeX-YAML

`go install github.com/Anthony-Gambale/BibTeX-YAML@v1.0.0`

You'll also need to add your go binaries to the PATH

`export PATH=$PATH:$HOME/go/bin`

Don't forget to add this to `~/.bashrc` to run it every time you open a new shell.

Now, simply run

`BibTeX-YAML template <file-name>.yaml`

to generate (or overwrite) `<file-name>.bib` with your bibliography.

### Examples

The following YAML file is converted into the BibTeX below.

```yaml
entries:
  - id: MyPaper
    type: inproceedings
    fields:
      author: Me Myself and Someone Else
      booktitle: "Book Title"
      title: "Title"
#     year: 2024
#     pages: C137--$\infty$
#     publisher: Unsure
```

```BibTeX
@inproceedings{MyPaper,
    title = "Title",
    author = "Me Myself and Someone Else",
    booktitle = "Book Title",
}
```

See `test.yaml` and `test.bib` for a more comprehensive example.
