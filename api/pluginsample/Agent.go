package pluginsample

import "GLTI/goruntime/glti"

// InitGltiAgent is called by Go runtime to initialize the agent plugin.
func InitGltiAgent(env *glti.GLTIEnv) {

}

// ShutdownGltiAgent is called by Go runtime to gracefully shutdown the agent plugin.
func ShutdownGltiAgent(exitcode int) {

}
