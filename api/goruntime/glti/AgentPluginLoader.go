package glti

const gltiVersion = "<GLTI version>"

// InitGLTI loads and initializes agent shared library. Agent file path is given by environment variable GO_GLTI_AGENT:
// 1. Native plugins
//    GO_GLTI_AGENT is set to file location of a C++ shared object which mimics a Go plugin.
// 2. Go based plugins
//    GO_GLTI_AGENT is set to a folder containing versioned Go plugins, e.g. "/opt/glti/agentplugin_<go build version>"
//    The Go runtime loads the correct Go plugin based on its Go version string which is specified to be contained in
//    the plugins file name (e.g. plugin_go1.12.8).
//
// Special case - Go application is statically linked:
//    GO_GLTI_LOADER is set to file location of a vendor tool responsible to start the Go application with applied
//    monitoring.
//
// An agent plugin exports the following symbols:
// - string variable "runtime.buildVersion" which represents the Go build version of the agent plugin. It may have the
//   special value "native" for C++ shared object plugins.
// - func InitGltiAgent(*glti.GLTIEnv): which is called to initialize the agent.
// - (optional) func ShutdownGltiAgent(int): which is called before Go application exits.
func InitGLTI() {
	panic("not implemented")
}
