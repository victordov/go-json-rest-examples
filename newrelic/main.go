package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
	"log"
	"net/http"
	"time"
)

type NewRelicMiddleware struct {
	License string
	Name    string
	Verbose bool
	agent   *gorelic.Agent
}

func (mw *NewRelicMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {

	mw.agent = gorelic.NewAgent()
	mw.agent.NewrelicLicense = mw.License
	mw.agent.HTTPTimer = metrics.NewTimer()
	mw.agent.Verbose = mw.Verbose
	mw.agent.NewrelicName = mw.Name
	mw.agent.CollectHTTPStat = true
	mw.agent.Run()

	return func(writer rest.ResponseWriter, request *rest.Request) {

		handler(writer, request)

		// the timer middleware keeps track of the time
		startTime := request.Env["START_TIME"].(*time.Time)
		mw.agent.HTTPTimer.UpdateSince(*startTime)
	}
}

func main() {
	handler := rest.ResourceHandler{
		OuterMiddlewares: []rest.Middleware{
			&NewRelicMiddleware{
				License: "<REPLACE WITH THE LICENSE KEY>",
				Name:    "<REPLACE WITH THE APP NAME>",
				Verbose: true,
			},
		},
	}
	err := handler.SetRoutes(
		&rest.Route{"GET", "/message", func(w rest.ResponseWriter, req *rest.Request) {
			w.WriteJson(map[string]string{"Body": "Hello World!"})
		}},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", &handler))
}
