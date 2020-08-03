package httpd

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/miiniper/loges"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type Service struct {
	addr string
	ln   net.Listener

	router *httprouter.Router
}

func New(listen string) (*Service, error) {
	return &Service{
		addr:   listen,
		router: httprouter.New(),
	}, nil
}

func (s *Service) Start() error {
	s.initHandler()

	server := http.Server{}
	server.Handler = s.router

	server.Handler = s.accessLog(cors(server.Handler))

	// Open listener.
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	s.ln = ln

	go func() {
		err := server.Serve(s.ln)
		if err != nil {
			loges.Loges.Error("httpd serve error", zap.Error(err))
		}
	}()
	loges.Loges.Info("httpd service started", zap.String("listen", s.addr))

	return nil
}

func (s *Service) Close() error {
	s.ln.Close()
	loges.Loges.Sync()
	return nil
}

func (s *Service) initHandler() {
	s.router.GET("/httpcheck", s.GetOk)
	s.router.POST("/sendmsg", SendMsg)

}

func cors(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set(`Access-Control-Allow-Origin`, origin)
			w.Header().Set(`Access-Control-Allow-Methods`, strings.Join([]string{
				`DELETE`,
				`GET`,
				`OPTIONS`,
				`POST`,
				`PUT`,
			}, ", "))

			w.Header().Set(`Access-Control-Allow-Headers`, strings.Join([]string{
				`Accept`,
				`Accept-Encoding`,
				`Authorization`,
				`Content-Length`,
				`Content-Type`,
				`X-CSRF-Token`,
				`X-HTTP-Method-Override`,
				`Authtoken`,
				`X-Requested-With`,
				`NS`,
				`Resource`,
			}, ", "))
		}

		if r.Method == "OPTIONS" {
			return
		}

		inner.ServeHTTP(w, r)
	})
}

func (s *Service) accessLog(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			inner.ServeHTTP(w, r)
			return
		}
		stime := time.Now().UnixNano() / 1e3
		inner.ServeHTTP(w, r)
		dur := time.Now().UnixNano()/1e3 - stime
		remoteIP := r.Header.Get("RemoteClentIP")
		if dur <= 1e3 {
			loges.Loges.Info("http access", zap.String("method", r.Method), zap.String("uri", r.RequestURI), zap.Int64("time(us)", dur), zap.String("RemoteClentIP", remoteIP))
		} else {
			loges.Loges.Info("http access", zap.String("method", r.Method), zap.String("uri", r.RequestURI), zap.Int64("time(ms)", dur/1e3), zap.String("RemoteClentIP", remoteIP))
		}
	})
}

func (s *Service) GetOk(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("ok"))
	return
}
