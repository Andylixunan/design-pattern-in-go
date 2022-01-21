package factory

type IConfigParserFactory interface {
	CreateParser() IConfigParser
}

// JSONConfigParserFactory implements IConfigParserFactory
type JSONConfigParserFactory struct{}

func (JSONConfigParserFactory) CreateParser() IConfigParser {
	return JSONConfigParser{}
}

// YAMLConfigParserFactory implements IConfigParserFactory
type YAMLConfigParserFactory struct{}

func (YAMLConfigParserFactory) CreateParser() IConfigParser {
	return YAMLConfigParser{}
}

var _ IConfigParserFactory = JSONConfigParserFactory{}
var _ IConfigParserFactory = YAMLConfigParserFactory{}
