package p

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	_ "github.com/go-gl/glfw/v3.2/glfw"
)

func CreateProgram() (program uint32, err error) {
	err = gl.Init()
	if err != nil {
		println(gl.GetError())
		return
	}

	program = gl.CreateProgram()

	return
}
