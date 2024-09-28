package routers

import (
	"net/http"
)

func MountStaticRouter(prefix string, path string) {
	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(path))))
}
