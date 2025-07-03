package plugin

type InputPlugin interface {
	Init(config map[string]interface{}) error
	Read() (map[string]interface{}, error)
	Close() error
}

type TransformPlugin interface {
	Execute(record map[string]interface{}) (map[string]interface{}, error)
}

type OutputPlugin interface {
	Write(record map[string]interface{}) error
	Close() error
}
