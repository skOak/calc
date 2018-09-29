package variables

import "strings"

type ValueSource interface {
	GetValue(str string) float64
}
type ValueSourceFunc func(string) float64

func (vs ValueSourceFunc) GetValue(str string) float64 {
	return vs(str)
}

// str:{{.variable_name}}
func IsVariable(str string) bool {
	if !strings.HasPrefix(str, "__") {
		return false
	}
	return true
}

func GetValue(str string, source ValueSource) (v float64) {
	if !IsVariable(str) {
		return 0
	}
	if source == nil {
		return 0
	}
	variableName := strings.TrimPrefix(str, "__")
	return source.GetValue(variableName)
}
