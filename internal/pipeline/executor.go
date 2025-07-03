package pipeline

import (
	"github.com/onyxnox/helixflow/internal/plugin"
)

type Pipeline struct {
	Input      plugin.InputPlugin
	Transforms []plugin.TransformPlugin
	Output     plugin.OutputPlugin
}

func (pipeline *Pipeline) ExecuteOnce() error {
	record, err := pipeline.Input.Read()
	if err != nil {
		return err
	}

	for _, transform := range pipeline.Transforms {
		record, err = transform.Execute(record)
		if err != nil {
			return err
		}
	}

	return pipeline.Output.Write(record)
}
