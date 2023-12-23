package p

import (
	"github.com/go-gl/gl"
)

func CreateProgram() (program uint32, err error) {
	err = gl.Init()
	if err != nil {
		return
	}

	program = gl.CreateProgram()

	return
}
