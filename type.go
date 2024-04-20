package configkit

type StructReader[TData any] interface {
	Read(path string) (*TData, error)
}

type StructWriter[TData any] interface {
	Write(path string, configStruct TData) error
}

type StructReaderWriter[TData any] interface {
	StructReader[TData]
	StructWriter[TData]
}

type EnvReader interface {
	Read(key string) string
	ReadSafe(key string) (string, error)
}

type EnvWriter interface {
	Write(key, value string) error
}

type EnvReaderWriter interface {
	EnvReader
	EnvWriter
}

type ConvertorFunc func(s string) string

type Settings[TData any] interface {
	Load() (*TData, error)
}
