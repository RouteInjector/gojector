package conf

type Configuration struct {
	Database *Database
	Bind     int
	Images   *Images
	Auth     bool
	Swagger  bool
}

type Database struct {
	Endpoint string
	Name     string
	debug    bool
	Logger   Logger
}

type Images struct {
	Path          string
	Cache         string
	GalleryFolder string
}

type Logger struct {
	Level string
}