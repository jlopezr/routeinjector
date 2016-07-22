package routeinjector

import (
	"fmt"
	"reflect"
	"user/routeinjector"

	"github.com/oleiade/reflections"
)

//RegisterSchema registers a new model from a JSON Schema
func (ri *RouteInjector) RegisterSchema() {
}

//RegisterModel registers a new schema from a Golang object in the system
func (ri *RouteInjector) RegisterModel(o interface{}) Schema {

	fields, err := reflections.Fields(o)
	if err != nil {
		panic(err)
	}
	name := reflect.TypeOf(o)
	fmt.Println("NAME", name)
	fmt.Printf("OBJECT %+v\n", o)
	fmt.Printf("FIELDS %+v\n", fields)

	schema := routeinjector.Schema{}

	structTags, _ := reflections.Tags(o, "description")
	fmt.Printf("%v+\n", structTags)

	return schema
}
