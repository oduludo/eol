# Getting started
EOL is designed to be as simple and straightforward in use as possible.
Advanced features are available, but in its simplest form, usage works like this:

```shell
eol check <resource> <version>
```

For example: `eol check ruby 2.7` will print `true`, indicating `ruby@2.7` has passed its end of life.
Releases which have not passed their end of life, `false` will be printed.
