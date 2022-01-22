package factory

import "fmt"

type ISystemConfigParser interface {
	ParseSystemConfig(data []byte)
}

// JSONSystemConfigParser implements ISystemConfigParser
type JSONSystemConfigParser struct{}

func (JSONSystemConfigParser) ParseSystemConfig(data []byte) {
	fmt.Printf("parsing JSON system config: %v", data)
}

// YAMLSystemConfigParser implements ISystemConfigParser
type YAMLSystemConfigParser struct{}

func (YAMLSystemConfigParser) ParseSystemConfig(data []byte) {
	fmt.Printf("parsing YAML system config: %v", data)
}

var _ ISystemConfigParser = JSONSystemConfigParser{}
var _ ISystemConfigParser = YAMLSystemConfigParser{}

// IConfigParserFamilyFactory is the abstract factory
type IConfigParserFamilyFactory interface {
	CreateParser() IConfigParser
	CreateSystemParser() ISystemConfigParser
}

// JSONConfigParserFamilyFactory implements IConfigParserFamilyFactory
type JSONConfigParserFamilyFactory struct{}

func (JSONConfigParserFamilyFactory) CreateParser() IConfigParser {
	return JSONConfigParser{}
}

func (JSONConfigParserFamilyFactory) CreateSystemParser() ISystemConfigParser {
	return JSONSystemConfigParser{}
}

// YAMLConfigParserFamilyFactory implements IConfigParserFamilyFactory
type YAMLConfigParserFamilyFactory struct{}

func (YAMLConfigParserFamilyFactory) CreateParser() IConfigParser {
	return YAMLConfigParser{}
}

func (YAMLConfigParserFamilyFactory) CreateSystemParser() ISystemConfigParser {
	return YAMLSystemConfigParser{}
}
