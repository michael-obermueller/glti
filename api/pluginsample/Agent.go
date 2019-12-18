package pluginsample

import "GLTI/goruntime/glti"

// InitGltiAgent is called by Go runtime to initialize the agent plugin.
func InitGltiAgent(env *glti.GLTIEnv) {

}

// ShutdownGltiAgent is called by Go runtime to gracefully shut down the agent plugin.
func ShutdownGltiAgent(exitcode int) {

}
