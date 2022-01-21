package factory

import "fmt"

type IConfigParser interface {
	Parse(data []byte)
}

// JSONConfigParser implements IConfigParser
type JSONConfigParser struct{}

func (JSONConfigParser) Parse(data []byte) {
	fmt.Printf("parsing JSON: %v", data)
}

// YAMLConfigParser implements IConfigParser
type YAMLConfigParser struct{}

func (YAMLConfigParser) Parse(data []byte) {
	fmt.Printf("parsing YAML: %v", data)
}

var _ IConfigParser = JSONConfigParser{}
var _ IConfigParser = YAMLConfigParser{}

func NewConfigParser(t string) IConfigParser {
	switch t {
	case "JSON":
		return JSONConfigParser{}
	case "YAML":
		return YAMLConfigParser{}
	}
	return nil
}
