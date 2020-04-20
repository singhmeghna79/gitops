package controller

import (
	"github.com/storage-provisiong-poc/gittrack/pkg/controller/gittrack"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, gittrack.Add)
}
