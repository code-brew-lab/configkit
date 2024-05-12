package configkit

import (
	"fmt"
	"reflect"
	"strings"
)

type settings[TData any] struct {
	path          string
	structReader  StructReader[TData]
	envReader     EnvReader
	convertorFunc ConvertorFunc
}

func NewSettings[TData any](path string, convertor ...ConvertorFunc) Settings[TData] {
	settings := &settings[TData]{
		path:         path,
		structReader: newJsonRW[TData](),
		envReader:    newOsEnvRW(),
	}

	if len(convertor) == 0 {
		settings.convertorFunc = strings.ToUpper
	} else {
		settings.convertorFunc = convertor[0]
	}

	return settings
}

func (s *settings[TData]) Load() (*TData, error) {
	conf, err := s.structReader.Read(s.path)
	if err != nil {
		return nil, err
	}

	confValue := reflect.ValueOf(conf)

	if err := s.applyEnvOverrides(confValue.Elem(), ""); err != nil {
		return nil, err
	}

	return conf, nil
}

func (s *settings[TData]) applyEnvOverrides(v reflect.Value, prefix string) error {
	fields := reflect.VisibleFields(v.Type())

	for i, field := range fields {
		value := v.Field(i)
		p := s.mergePrefix(prefix, field.Name)

		if value.Kind() == reflect.Pointer {
			value = value.Elem()
		}
		if value.Kind() == reflect.Struct {
			s.applyEnvOverrides(value, p)
			continue
		}
		if !s.isEligibleForEnv(value.Type()) {
			continue
		}

		fmt.Println(p)
		envVal, err := s.envReader.ReadSafe(s.convertorFunc(p))
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

func (s *settings[TData]) mergePrefix(prefix, field string) string {
	if strings.Compare(prefix, "") == 0 {
		return field
	}

	return fmt.Sprintf("%s_%s", prefix, field)
}
