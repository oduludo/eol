# Usage
This page covers all of EOL's available functionalities. For a short guide on getting started, see the [getting started guide](./getting_started.md).

Multiple commands are available. `eol --help` and `eol help` will output the program's help message.

Available commands are:
- `eol check`
- `eol help`
- `eol list-resources`
- `eol list-versions`
- `eol source`
  - `eol source decrypt`
  - `eol source encrypt`
  - `eol source verify`

## Check
The core function of EOL is to check if a resource's version has passed its end of life. This is done using the `check` command.

```
Usage:
  eol check RESOURCE VERSION [flags]

Flags:
      --e                set to true to give a non-zero exit code on EOL
  -h, --help             help for check
      --key string       one or more keys to decrypt custom sources, delimiting keys using ','
      --source string    combine the root source with one or more URLs pointing to custom datasource resources, delimiting URLs using ','
      --xsource string   exclusively use one or more URLs pointing to custom datasource resources, delimiting URLs using ',', not using any of the root data
```

The resource is the software product you want to check for, e.g. Ruby (`ruby`) or the Numpy Python package (`numpy`).
Resources can be checked if their data is present in the underlying (root) datasource.
If your resource is not listed, whether it's simply not included or a custom/internal product, custom datasources can be configured (see below).

The version is the specific version to check for end of life status. Be aware that the formatting of versions is not consistent.
Ruby may use incrementing version numbers, while Amazon's Linux versioning uses a `2018.03`-like formatting.
The format being used lies with the party that defines the versions, not with the root datasource.

### Simple EOL check
Running `eol check ruby 2.7` will print `true`, indicating `ruby@2.7` has reached end of life.
If the resource version has not yet reached end of life, `false` will be printed.

### Enabling non-zero exit code
There are situations where it is desirable the program indicates a 'failure' using an exit code.
Think of a CI job which needs a command to fail with an exit code to be considered a failed job in the CI.

EOL offers this functionality, by adding the `--e` flag to your command.
If the version has reached end of life, the program exits with a non-zero exit code.
If the version has not reached end of life yet, the program will exit with exit code `0`.

### Using custom datasources
The root datasource does not contain all public software versions for the entire internet, let alone covering internal software or small open source projects.
EOL allows you to configure custom datasources in two modes:
- On top of the root source
- Exclusively using the custom source(s)

If you want to add custom resources and versions on top of the root source, use the `--source` flag.
If you want to exclusively use configured custom sources instead of the root source, use the `--xsource` flag.
Apart from the inclusion/exclusion of the root source, both flags work the same way.

#### Specifying sources
One or more sources can be specified under the `--(x)source` flag.
Use `,` for a delimiter between URLs when specifying multiple sources.

For `--source`, the root source is checked last.
So if your custom source would say `x@y` has reached end of life, but the root source say it hasn't, the program will consider `x@y` having reached end of life.

For both `--source` and `--xsource`, the order of URLs is followed when determining what end of life information to use.
If `x@y` is listed in two datasources, the information from the first datasource is used.

#### Working with encrypted data
EOL gives you the option to encrypt your datasource.
This comes in handy if you are forced to put the datasource out on the public internet, but don't want others to read the information. 
The details for encrypting and decrypting datasources are listed further below.

To use an encrypted datasource, pass the `--key` flag to the command.
It is important to list as many keys as sources, to ensure the correct key is mapped to the correct source.
`_` can be used as a placeholder when mixing encrypted and unencrypted sources, e.g. `--key=_,KEY_FOR_URL2,_` if only the second out of three URLs serves encrypted data.

The `--key` flag can be omitted if you are only using unencrypted sources.

## Listing resources
Besides checking end of life statuses, you may want to see what resources can be checked.
This is done using the `list-resources` command.

```
Usage:
  eol list-resources [flags]

Flags:
      --contains string   filtering on resources that contain the provided substring
  -h, --help              help for list-resources
```

> `list-resources` currently does not accept custom datasources, but this feature is on the project's roadmap.

### Filtering resources
The returned list can be quite long, so a `--contains` flag is available for filtering.
`--contains` does a simple substring check to limit the list's size.

Let's say you only want to see resources related to Amazon.
Run `eol list-resources --contains=amazon` to get a limited list of resources that all have a name with 'amazon' in it.

## Listing versions
If you know the resource you want to check, but aren't sure the particular version is included in the root source, run the `list-versions` command.

```
Usage:
  eol list-versions RESOURCE [flags]

Flags:
  -h, --help   help for list-versions
```

> `list-versions` currently does not accept custom datasources, but this feature is on the project's roadmap.

The output will be a table with the columns 'Cycle name' and 'Release name', or an error message if the requested resource is unknown.

## Encrypting and decrypting datasources
As mentioned earlier, EOL allows you to encrypt your datasources as to obfuscate them.
Tooling for this is in subcommands under the `source` command.

### Encrypting a datasource
To encrypt a datasource, use `source encrypt`.

```
Usage:
  eol source encrypt FILE [flags]

Flags:
  -h, --help         help for encrypt
      --key string   optionally configure a key to use for encryption
      --to string    location to write encrypted data to
```

The command will behave differently depending on the flags set.

#### In-place encryption versus writing to a different file
By default, `source encrypt` overwrites the contents of the input file with the encrypted data.
This can be prevented by configuring the `--to` flag. `--to` can be set to the desired output file path.
If set, the encrypted data won't overwrite the original data, but will be written to the output file location.

#### Specifying an encryption key
If you already have an encryption key to use, you can set it via the `--key` flag.
Encryption keys must have a length of 16 characters.

> EOL uses CFB for encryption.

If you don't provide an encryption key, the program will randomly generate one.
The key is logged to the console after encryption.
Store it well, as you won't be able to retrieve it from the encrypted data afterwards.

### Decrypting a datasource
To decrypt a datasource, use `source decrypt`. This requires the encryption key.

```
Usage:
  eol source decrypt FILE [flags]

Flags:
  -h, --help         help for decrypt
      --key string   configure a key to use for decryption
      --to string    location to write decrypted data to
```

The command will behave differently depending on the flags set.
It takes the input file and attempts to decrypt it using the key passed in using `--key`.

#### In-place decryption versus writing to a different file
The default behaviour for this command is to overwrite the encrypted file data with the decrypted data.
If you want the decrypted data to be written to a different file, specify an output file path in the `--to` flag.

### Validating a datasource
A datasource can be validated on its structure using `source verify`.

```
Usage:
  eol source verify FILE [flags]

Flags:
  -h, --help         help for verify
      --key string   optionally configure a key to use for encryption
```

Validation works for both unencrypted and encrypted datasources.
When running on an encrypted datasource, pass in the appropriate encryption key via the `--key` flag.
