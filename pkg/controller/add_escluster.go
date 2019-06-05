package controller

import (
	"github.com/zyseeker/zy-es-operator/pkg/controller/escluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, escluster.Add)
}
