# go-bomd - Software Bill Of Materials to anything conversion

## About

**go-bomd** is a Golang CLI which allows you to take a
Software Bill Of Materials (SBOM) based on the [CycloneDX](https://cyclonedx.org)
standard and convert it into a different format by utilizing Golang's template
engine.

If you want to see an example, you can check out the [3rd Party](./3rd-party.md)
of this documentation.

## Why go-bomd?

One of the key use cases for creating go-bomd was to have a
tool which can take a SBOM file and transform relevant information
into a markdown file. This can easily be integrated into any
markdown based documentation framework like [MkDocs](https://www.mkdocs.org).
