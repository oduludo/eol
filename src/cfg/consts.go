package cfg

// TestAssets points to assets for testing in Docker container, per absolute path.
const (
	TestAssets              = "/app/test_assets"
	IsIntegrationTestEnvKey = "IS_INTEGRATION_TEST"
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
        },
        {
          "type": "object",
          "properties": {
            "eol": {
              "type": "string"
            },
            "cycle": {
              "type": "string"
            },
            "releaseName": {
              "type": "string"
            }
          },
          "required": [
            "eol",
            "cycle",
            "releaseName"
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
)
