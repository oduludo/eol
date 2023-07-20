package cfg

// TestAssets points to assets for testing in Docker container, per absolute path.
const (
	TestAssets               = "/app/test_assets"
	IsIntegrationTestEnvKey  = "IS_INTEGRATION_TEST"
	DecryptionKeyPlaceholder = "_"
)

const DatasourceSchema = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object",
  "patternProperties": {
    "^.*": {
      "type": "array",
      "items": [
        {
          "type": "object",
          "properties": {
            "eol": {
              "type": "string"
            },
            "cycle": {
              "type": "string"
            }
          },
          "required": [
            "eol",
            "cycle"
          ]
        }
      ]
    }
  }
}`

// Messages
const (
	DatasourceValidMsg   = "Datasource is valid"
	DatasourceInvalidMsg = "datasource schema is invalid"
	SourceXsourceXorMsg  = "only specify one of --source and --xsource, not both"
	InvalidKeysNumMsg    = "the number of keys does not match the number of configured datasources, use '_' per unencrypted datasource in the list"
	ZipLenMismatchMsg    = "mismatch in length for custom sources and keys"
)
