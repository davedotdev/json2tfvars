# json2tfvars

A simple CLI that parses arbritrary JSON to the HCL format of tfvars files.

## Installation

`go get github.com/imjoshholloway/json2tfvars`

## Usage

Pipe some `json` in:

```bash
echo '{ "string": "value", "object": { "key": { "subkey": "value" }, "key2": { "subkey2": "value" } }}' | json2tfvars

string = "value"
object = { key = { subkey = "value" }, key2 = { subkey2 = "value" } }
```

Alternatively pass in a `json` file:

```bash
json2tfvars -source=path/to/file.json

```

## Why is this needed?

In versions of [terraform][#terraform] below the currently unreleased `0.12`
nested maps in `json` variable files are not parsed correctly. This utility
converts `json` payloads to the `tfvars` format in order to work around this
issue.

See: https://github.com/hashicorp/terraform/issues/15549 for more information

