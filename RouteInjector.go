package routeinjector

type RouteInjector struct {
	Models []Model
}

type Model struct {
	Name      string  // Name of the model
	Plural    string  // Plural of the model
	ID        string  // Identifier of the model in the db
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

// Route provides the definition of a URL route
type Route struct {
	Path    string            // URL path where the route is map
	Method  string            // HTTP method the route uses
	Handler httprouter.Handle // Handler for the route
}