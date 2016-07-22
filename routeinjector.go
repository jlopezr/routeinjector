// Package routeinjector provides a system for webservice fast prototyping
//
// Author: Juan López <j.lopez.r@gmail.com>
package routeinjector

import (
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

// RouteInjector provides the root object of the system
type RouteInjector struct {
	Db     *mgo.Session //MongoDB Session
	Models []Model      //List of models of the system
}

// NewInjector creates a RouteInjector object and initializes the system
func NewInjector() *RouteInjector {
	r := RouteInjector{}
	return &r
}

// Start the system
func (ri *RouteInjector) Start() {
	ri.StartDatabase()
}

// Stop the system
func (ri *RouteInjector) Stop() {
	ri.StopDatabase()
}

// Model provides the definition of a model in the database
type Model struct {
	Name      string  // Name of the model
	Plural    string  // Plural of the model
	Get       bool    // Is GET (retrieve one instance) method enabled
	Put       bool    // Is PUT (update a instance) method enabled
	Post      bool    // Is POST (add a new instance) method enabled
	Delete    bool    // Is DELETE (delete one instance) method enabled
	Search    bool    // Is SEARCH (retrieve a list of instances matching an expression) method enabled
	Validate  bool    // Is VALIDATE (check the consistency of all instances of the model) method enabled
	Import    bool    // Is IMPORT (import from csv,json) method enabled
	Export    bool    // Is EXPORT (export to csv,json) method enabled
	Aggregate bool    // Is AGGREGATE (retrive an aggragated list of instances) method enabled
	Routes    []Route // List of additional routes for this model
}

//Schema keeps the list of fields and their attributes
type Schema struct {
	Fields []Field // List of fields
}

//Field keeps name, type and its attributes
type Field struct {
	Name string // Name of the Field
	Type string // Type of the field
}

// NewModel creates a new Model by setting Name and Plural
func NewModel(name string) *Model {
	r := Model{Name: name, Plural: name + "s"}
	return &r
}

// AddRoute adds a new route to the aditional routes for this model
func (m *Model) AddRoute(route Route) {
	m.Routes = append(m.Routes, route)
}

// ProcessRoutes injects the additional routes for this model in the router
func (m *Model) ProcessRoutes(router *httprouter.Router) {
	for _, r := range m.Routes {
		//fmt.Printf("** %v\n", r)
		switch r.Method {
		case "GET":
			router.GET(r.Path, r.Handler)
		case "POST":
			router.POST(r.Path, r.Handler)
		case "PUT":
			router.PUT(r.Path, r.Handler)
		case "DELETE":
			router.DELETE(r.Path, r.Handler)
		case "HEAD":
			router.HEAD(r.Path, r.Handler)
		case "PATCH":
			router.PATCH(r.Path, r.Handler)
		}
	}
}

// Route provides the definition of a URL route
type Route struct {
	Path    string            // URL path where the route is map
	Method  string            // HTTP method the route uses
	Handler httprouter.Handle // Handler for the route
}
