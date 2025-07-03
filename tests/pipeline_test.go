package tests

import (
	"testing"

	"github.com/onyxnox/helixflow/internal/pipeline"
	"github.com/onyxnox/helixflow/internal/plugin"
)

type mockInput struct{}

func (input *mockInput) Init(cfg map[string]any) error { return nil }
func (input *mockInput) Read() (map[string]any, error) {
	return map[string]any{"message": "hello"}, nil
}
func (input *mockInput) Close() error { return nil }

type goodbyeTransform struct{}

func (transform *goodbyeTransform) Execute(stream map[string]any) (map[string]any, error) {
	stream["message"] = "GOODBYE"

	return stream, nil
}

type mockOutput struct {
	written []map[string]any
}

func (output *mockOutput) Write(stream map[string]any) error {
	output.written = append(output.written, stream)

	return nil
}

func (output *mockOutput) Close() error { return nil }

func TestPipelineExecuteOnce(test *testing.T) {
	output := &mockOutput{}
	pipeline := pipeline.Pipeline{
		Input:      &mockInput{},
		Transforms: []plugin.TransformPlugin{&goodbyeTransform{}},
		Output:     output,
	}

	if err := pipeline.ExecuteOnce(); err != nil {
		test.Fatalf("pipeline failed: %v", err)
	}

	if output.written[0]["message"] != "GOODBYE" {
		test.Errorf("expected GOODBYE, got %v", output.written[0]["message"])
	}
}
