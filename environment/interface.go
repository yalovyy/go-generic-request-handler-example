package environment

import (
	"ghandler/api/bar"
	"ghandler/api/foo"
)

type Storage interface {
	bar.Reader
	foo.Writer
}

type Environment interface {
	GetStorage() Storage
}
