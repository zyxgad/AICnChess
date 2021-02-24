
package chess

import(
	runtime "runtime"

	// gl  "github.com/go-gl/gl/v2.1/gl"
	gl  "github.com/go-gl/gl/v4.1-core/gl"
	glfw "github.com/go-gl/glfw/v3.2/glfw"
)

var(
	window *glfw.Window = nil
	program uint32 = 0
)

func init(){
	runtime.LockOSThread()
}

func initOpengl() (_ uint32, err error) {
	err = gl.Init()
	if err != nil {
		return 0, err
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	INFO.Println("The opengl version is", version)
	program = gl.CreateProgram()
	gl.LinkProgram(program)

	vao = makeVao(triangle)
	return program, nil
}

func GetWindow() (_ *glfw.Window, err error){
	if window != nil {
		return window, nil
	}
	window, err = openGlfw()
	if err != nil {
		return nil, err
	}
	initOpengl()
	return window, nil
}

func openGlfw() (window *glfw.Window, err error){
	err = glfw.Init()
	if err != nil {
		ERROR.Println(err.Error())
		return nil, err
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err = glfw.CreateWindow(SCREEN_WIDTH, SCREEN_HEIGHT, SCREEN_TITLE, nil, nil)
	if err != nil {
		ERROR.Println(err.Error())
		return nil, err
	}
	window.MakeContextCurrent()
	return window, nil
}

func CloseWindow(){
	if window == nil {
		WARN.Println("You are trying to close a window that does not exist")
		return
	}
	glfw.Terminate()
	window = nil
}

func WindowShouldClose() bool {
	if window == nil {
		return true
	}
	return window.ShouldClose()
}

func clearWindow(){
	gl.Clear(gl.COLOR_BUFFER_BIT|gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)
}

var (
	vao uint32
	triangle []float32 = []float32{
		 0.0,  0.5, 0,
		-0.5, -0.5, 0,
		 0.5, -0.5, 0,
	}
)

func DrawWindow(){
	clearWindow()

    gl.BindVertexArray(vao)
    gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle) / 3))

	window.SwapBuffers()
}

func PollEvents(){
	glfw.PollEvents()
}

func makeVao(points []float32) (vao uint32) {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4 * len(points), gl.Ptr(points), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	return vao
}
