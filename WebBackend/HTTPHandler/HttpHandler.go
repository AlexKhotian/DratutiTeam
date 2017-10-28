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
			m, err := url.ParseQuery(r.URL.RawQuery)
			if err != nil {
				log.Println("Failed to parse url: ", m["h"][0])
				return
			} else {
				log.Println("Requested day type: ", m["h"][0])
				fragment, err := strconv.Atoi(m["h"][0])
				if err != nil {
					log.Println("Failed to convert fragment")
				}
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				handler.adapter.HandleDemandsRequest(w, fragment)
			}
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
