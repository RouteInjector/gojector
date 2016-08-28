package model

import (
	"github.com/RouteInjector/gojector/route"
	"strings"
)

type Model struct {
	Schema    Schema        // Schema of the model
	Name      string        // Name of the model
	Plural    string        // Plural of the model
	ID        string        // Identifier of the model in the db
	Get       bool          // Is GET (retrieve one instance) method enabled
	Put       bool          // Is PUT (update a instance) method enabled
	Post      bool          // Is POST (add a new instance) method enabled
	Delete    bool          // Is DELETE (delete one instance) method enabled
	Search    bool          // Is SEARCH (retrieve a list of instances matching an expression) method enabled
	Validate  bool          // Is VALIDATE (check the consistency of all instances of the model) method enabled
	Import    bool          // Is IMPORT (import from csv,json) method enabled
	Export    bool          // Is EXPORT (export to csv,json) method enabled
	Aggregate bool          // Is AGGREGATE (retrive an aggragated list of instances) method enabled
	Routes    []route.Route // List of additional routes for this model
}

func (m *Model) GetPathName() string {
	return "/" + strings.ToLower(m.Name)
}

func (m *Model) GetPluralPathName() string {
	return "/" + strings.ToLower(m.Plural)
}

type Schema interface {

}