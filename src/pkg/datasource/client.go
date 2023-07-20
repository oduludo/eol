package datasource

type ClientInterface interface {
	Resources() ([]string, error)
	All(args ...string) ([]ListedCycleDetail, error, bool)
	Get(args ...string) (CycleDetail, error, bool)
	GetCustom(source, resource, version, decryptionKey string) (ListedCycleDetail, error, bool)
}
