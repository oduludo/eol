# EOL
Simple CLI application to check whether a tool's version has reached end-of-life (EOL).

EOL uses https://endoflife.date/ as a data source.

## Usage
Multiple commands are available. `eol --help` will output the program's help message.

### Check if a version is passed EOL
`eol check RESOURCE VERSION` will print `true` if the resource's version is passed its EOL and `false` if not.
Optionally the `--e` flag can be added to ensure a non-zero exit code is used if the EOL date has passed. This ensures
easy usage of EOL in pipelines.

If the resource's version isn't present in the datasource, a table with cycle and release names will be printed.
The cycle names are the values to use for the version and have proper formatting for the API.
The release names are optional and will be shown if a cycle has a release name that differs from the formatted version.

### List versions
Prints the table of available versions for a resource, described above.

### List resources
Prints a list of resources present in the datasource.
Optionally the `--contains=xxx` flag may be set to do some filtering. The output will then only contain resources which
names contain the provided substring.

## Docker
Images are available on Docker Hub: https://hub.docker.com/r/ludovanorden/eol
