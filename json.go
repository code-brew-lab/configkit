package configkit

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type jsonRW[TData any] struct {
}

func newJsonRW[TData any]() StructReaderWriter[TData] {
	return &jsonRW[TData]{}
}

func (jrw *jsonRW[TData]) Read(path string) (*TData, error) {
	var (
		bytes []byte
		model = new(TData)
		err   error
	)

	if !filepath.IsLocal(path) {
		return nil, fmt.Errorf("given path is out of project root '%s'", path)
	}

	path = filepath.Join("./", filepath.Clean(path))

	bytes, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (jrw *jsonRW[TData]) Write(path string, configStruct TData) error {
	var (
		bytes []byte
		err   error
	)

	bytes, err = json.Marshal(configStruct)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, bytes, 0600)
	if err != nil {
		return err
	}
	return nil
}
