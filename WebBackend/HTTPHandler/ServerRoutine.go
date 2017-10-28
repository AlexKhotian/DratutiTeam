package HTTPHandler

import "net/http"

// IServerRoutine starts a web handler server
type IServerRoutine interface {
	RunServer()
}

// ServerRoutineImpl implementation for ServerRoutine interface
// contains handler
type ServerRoutineImpl struct {
	httpHandlerUtil *HTTPHandlerUtil
}

// ServerRoutineFactory creats intance of ServerRoutine
func ServerRoutineFactory() IServerRoutine {
	thisServer := new(ServerRoutineImpl)
	thisServer.httpHandlerUtil = HTTPHandlerFactory()
	return thisServer
}

// RunServer starts a web server
func (serverImpl *ServerRoutineImpl) RunServer() {
	http.Handle("/", serverImpl.httpHandlerUtil)
	http.ListenAndServe(":7777", nil)
}
