package conf

import (
	"os"
)

func LoadConfig(){
	os.Getenv("GO_ENV")
}