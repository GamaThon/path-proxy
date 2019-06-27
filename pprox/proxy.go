package pprox

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy(c *Config) {

	r := mux.NewRouter()

	for _, route := range c.Routes {
		u, err := url.Parse(route.Target)
		if err != nil {
			logrus.Errorln("Could not target", route.Target)
			panic(err)
		}
		rp := httputil.NewSingleHostReverseProxy(u)
		r.PathPrefix(route.Prefix).HandlerFunc(rp.ServeHTTP)
		logrus.Println("Mapping", route.Prefix, "to", route.Target)
	}

	l, err := net.Listen("tcp", c.Listen)
	if err != nil {
		logrus.Println(err)
	}
	panic(http.Serve(l, r))

}
