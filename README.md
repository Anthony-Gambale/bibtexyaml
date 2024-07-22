# BibTeX-YAML

Write your bibliography in YAML and convert it to BibTeX. CLI built using [Cobra](https://github.com/spf13/cobra).

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
  - id: Oetiker2021LatexIntroduction
    type: misc
    fields:
      author: Tobias Oetiker
      year: 2021
      title: A (Not So) Short Introduction to LaTeX 2_ε
#     url: https://www.ctan.org/tex-archive/info/lshort/english/
```

```BibTeX
@misc{Oetiker2021LatexIntroduction,
    year = "2021",
    title = "A (Not So) Short Introduction to LaTeX 2_ε",
    author = "Tobias Oetiker",
}
```

See `test.yaml` and `test.bib` for a more comprehensive example.

### Todo (contribution wanted)

- [x] Compile YAML bibliography into BibTeX
- [ ] Quality of life changes
  - [ ] Better error messages
  - [ ] Stop printing templating output to terminal
- [ ] (Stretch goal) Reverse templating: convert BibTeX entries into YAML
