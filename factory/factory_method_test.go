package factory

import (
	"reflect"
	"testing"
)

func TestConfigParserFactory(t *testing.T) {
	JSONFactory := JSONConfigParserFactory{}
	YAMLFactory := YAMLConfigParserFactory{}
	if JSONParser := JSONFactory.CreateParser(); !reflect.DeepEqual(JSONParser, JSONConfigParser{}) {
		t.Errorf("Creating JSON Parser error ==> got: %v, want %v", JSONParser, JSONConfigParser{})
	}
	if YAMLParser := YAMLFactory.CreateParser(); !reflect.DeepEqual(YAMLParser, YAMLConfigParser{}) {
		t.Errorf("Creating YAML Parser error ==> got: %v, want %v", YAMLParser, YAMLConfigParser{})
	}
}
