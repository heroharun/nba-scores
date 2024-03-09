package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"nba-scores/internal/api"
	"net/http"
	"sync"
	"time"
)

type router struct {
	lock            *sync.Mutex
	serverStartTime time.Time
	api             api.API
}

func (router *router) Start() {
	router.serverStartTime = time.Now()
	fmt.Println("Server is running on port 8080")
	http.Handle("/ws", simulation(router.api))
	http.HandleFunc("/", SampleUsage)
	http.ListenAndServe(":8080", nil)
}

func simulation(api api.API) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP request to WebSocket
		conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
			return
		}
		defer conn.Close()

		// Simulate matches and send updates
		simulateMatch(conn, api)
	})
}

func simulateMatch(conn *websocket.Conn, api api.API) {
	for i := 1; i < 49; i++ { // Simulate 48 minutes
		match := api.Simulator.GetOneMinuteSimulation(i)
		matchJSON, _ := json.Marshal(match)

		conn.WriteMessage(websocket.TextMessage, matchJSON)
		time.Sleep(5 * time.Second) // Wait for 5 seconds for every minute
	}

}

func SampleUsage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Sample Usage")
	fmt.Fprintln(w, "/(ws)play-basketball")
}

var (
	myRouter   *router
	routerOnce sync.Once
)

func Router(api api.API) router {
	routerOnce.Do(func() {
		myRouter = &router{
			lock:            &sync.Mutex{},
			serverStartTime: time.Time{},
			api:             api,
		}
	})
	return *myRouter
}
