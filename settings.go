package configkit

import (
	"reflect"
)

type settings[TData any] struct {
	path          string
	structReader  StructReader[TData]
	envReader     EnvReader
	convertorFunc ConvertorFunc
}

func NewSettings[TData any](path string) Settings[TData] {
	return &settings[TData]{
		path:          path,
		structReader:  NewJsonRW[TData](),
		envReader:     NewOsEnvRW(),
		convertorFunc: camelToSnakeCaseConvertor,
	}
}

func (s *settings[TData]) Load() (*TData, error) {
	conf, err := s.structReader.Read(s.path)
	if err != nil {
		return nil, err
	}

	confValue := reflect.ValueOf(conf)

	if err := s.applyEnvOverrides(confValue.Elem()); err != nil {
		return nil, err
	}

	return conf, nil
}

func (s *settings[TData]) applyEnvOverrides(v reflect.Value) error {
	fields := reflect.VisibleFields(v.Type())

	for i, field := range fields {
		value := v.Field(i)

		if value.Kind() == reflect.Pointer {
			value = value.Elem()
		}
		if value.Kind() == reflect.Struct {
			s.applyEnvOverrides(value)
			continue
		}
		if !s.isEligibleForEnv(value.Type()) {
			continue
		}

		envVal, err := s.envReader.ReadSafe(field.Name)
		if err != nil {
			continue
		}

		if !value.CanSet() {
			continue
		}

		convertedValue := reflect.ValueOf(envVal).Convert(value.Type())
		value.Set(convertedValue)
	}

	return nil
}

func (s *settings[TData]) isEligibleForEnv(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Chan, reflect.Map, reflect.Func, reflect.Interface, reflect.Struct:
		return false
	default:
		return true
	}
}
