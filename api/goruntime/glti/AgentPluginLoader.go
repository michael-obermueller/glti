package glti

const gltiVersion = "<GLTI version>"

// InitGLTI loads and initializes the agent shared library. Agent file path is given by environment variable GO_GLTI_AGENT:
// 1. Native plugins: GO_GLTI_AGENT is set to file location of a non-Go shared object which mimics a Go plugin.
// 2. Go based plugins: GO_GLTI_AGENT is set to a folder containing versioned Go plugins, e.g. "/opt/glti/".
//    The Go runtime loads the correct Go plugin based on its Go version string which is contained in
//    the plugin's file name (e.g. plugin_go1.12.8.so).
//
// Special case - Go application is statically linked:
//    GO_GLTI_LOADER is set to file location of a vendor tool responsible for starting the Go application with
//    monitoring applied.
//
// An agent plugin exports the following symbols:
// - String variable "runtime.buildVersion": represents the Go build version of the agent plugin. It may have the
//   special value "native" for non-Go shared object plugins.
// - func InitGltiAgent(*glti.GLTIEnv): called to initialize the agent.
// - (optional) func ShutdownGltiAgent(int): called before the Go application exits.
func InitGLTI() {
	panic("not implemented")
}
