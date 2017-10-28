package HTTPHandler

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"strconv"
	"net/url"
)

// HTTPHandlerUtil implement interface
type HTTPHandlerUtil struct {
	adapter *PredictionAdapter
}

// HTTPHandlerFactory creates a interface for HTTPHandlerUtil
func HTTPHandlerFactory() *HTTPHandlerUtil {
	thisHandler := new(HTTPHandlerUtil)
	thisHandler.adapter = new(PredictionAdapter)
	thisHandler.adapter.InitParser()
	return thisHandler
}

func (handler *HTTPHandlerUtil) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	if r.Method == "GET" {
		if r.URL.Path == "/Demands" {
			m, _ := url.ParseQuery(r.URL.RawQuery)
			log.Println(m)
			log.Println(m["k"][0])
			fragment, _ := strconv.Atoi(m["k"][0])
		//	log.Println("Fragment: ", fragment)
			handler.adapter.HandleDemandsRequest(w, fragment)
			return
		}
	}
	path := r.URL.Path
	if path == "/" {
		path = "/StartPage.html"
	}

	modifiedPath, contentType := handler.parseRequest(&path)
	data, err := ioutil.ReadFile(string(modifiedPath))
	log.Println(modifiedPath)
	if err == nil {
		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 not found"))
	}
}

func (handler *HTTPHandlerUtil) parseRequest(path *string) (string, string) {
	var contentType string
	modifiedSourcePath := "../../Webview" + *(path)
	if strings.HasSuffix(*path, ".html") {
		contentType = "text/html"
	} else if strings.HasSuffix(*path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(*path, ".js") {
		modifiedSourcePath = "../../ViewModel" + *(path)
		contentType = "application/javascript"
	} else if strings.HasSuffix(*path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	return modifiedSourcePath, contentType
}