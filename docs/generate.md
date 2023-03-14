# How to generate a different format

## Using a template for markdown conversion

After successfully [installing](./install.md) go-bomd one can
use the following command to create a minimal markdown template:

```bash
bomd generate template --file my.tmpl
```

Which will create a new file called `my.tmpl` with a content
similar to:

```
# SBOM for {{ .Metadata.Component.Name }}
| Name | Version | Type |
| ---- | ------- | ---- |
{{ range .Components }}
| {{ .Name }} | {{ .Version }} | {{ .Type }} |
{{ end }}
```

If you are not familiar with Golang's template engine, you can
study the [documentation](https://pkg.go.dev/text/template).

Elements accessible, from within a template, are defined based on 
[cyclonedx-go](https://github.com/CycloneDX/cyclonedx-go). The
details can be found inside [cyclonedx.go](https://github.com/CycloneDX/cyclonedx-go/blob/master/cyclonedx.go).

## Generating a markdown file based on a SBOM

Once you have your template and a SBOM file as JSON, you can
generate the result via:

```bash
bomd generate result --file ./bom.json --file ./my.tmpl --file ./result.md 
```
As you can see, `--file` can be used multiple times and the order is not
important. It is assumed that a file ending with `.json` defines the 
SBOM, that a file ending with `.tmpl` defines the template and that any other
file represents the final output.

Using the template `my.tmpl`, which was generated in the first step, with the
SBOM for go-bomd would result in an output similar to:

```markdown
# SBOM for github.com/HaRo87/go-bomd
| Name | Version | Type |
| ---- | ------- | ---- |
| github.com/CycloneDX/cyclonedx-go | v0.7.0 | library |
| github.com/davecgh/go-spew | v1.1.1 | library |
```
