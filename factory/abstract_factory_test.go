package factory

import (
	"reflect"
	"testing"
)

func TestConfigParserFamilyFactory(t *testing.T) {
	JSONFamilyFactory := JSONConfigParserFamilyFactory{}
	YAMLFamilyFactory := YAMLConfigParserFamilyFactory{}
	jp, jsp := JSONFamilyFactory.CreateParser(), JSONFamilyFactory.CreateSystemParser()
	if !reflect.DeepEqual(jp, JSONConfigParser{}) || !reflect.DeepEqual(jsp, JSONSystemConfigParser{}) {
		t.Errorf("got json parser: %v, json system parser: %v\n", jp, jsp)
	}
	yp, ysp := YAMLFamilyFactory.CreateParser(), YAMLFamilyFactory.CreateSystemParser()
	if !reflect.DeepEqual(yp, YAMLConfigParser{}) || !reflect.DeepEqual(ysp, YAMLSystemConfigParser{}) {
		t.Errorf("got yaml parser: %v, yaml system parser: %v\n", yp, ysp)
	}
}
