package conf

type IConfigLoader interface {
	SetFinder(finder IConfigFinder)
	Load() (ServerConfig, error)
}
