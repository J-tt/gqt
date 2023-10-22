package server

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/centrifugal/centrifuge"
)

const creatureSubscription = "creatures"
const settingsSubscription = "settings"

type WSServer struct {
	node *centrifuge.Node
	db   *gorm.DB
}

func InitWeb() *WSServer {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}

	cfg := centrifuge.Config{}
	cfg.LogLevel = centrifuge.LogLevelDebug

	node, err := centrifuge.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	server := WSServer{
		node,
		db,
	}

	node.OnConnect(func(client *centrifuge.Client) {
		transportName := client.Transport().Name()
		transportProto := client.Transport().Protocol()
		log.Printf("client connected via %s (%s)", transportName, transportProto)

		client.OnSubscribe(func(e centrifuge.SubscribeEvent, cb centrifuge.SubscribeCallback) {
			log.Printf("client subscribes on channel %s", e.Channel)
			cb(centrifuge.SubscribeReply{}, nil)
			switch e.Channel {
			case creatureSubscription:
				state, err := server.GetState()
				if err != nil {
					log.Println(err)
				}
				stateBytes, err := json.Marshal(state)
				if err != nil {
					log.Println(err)
				}
				node.Publish(creatureSubscription, stateBytes)
			case settingsSubscription:
				// TODO
			default:
				// TODO
			}
		})
		client.OnDisconnect(func(e centrifuge.DisconnectEvent) {
			log.Printf("client disconnected")
		})

	})

	if err := node.Run(); err != nil {
		log.Fatal(err)
	}

	wsHandler := centrifuge.NewWebsocketHandler(node, centrifuge.WebsocketConfig{})
	http.Handle("/connection/websocket", auth(wsHandler))

	http.Handle("/", http.FileServer(http.Dir("./client")))
	go runWeb()
	return &server
}

func (wss *WSServer) Publish(creatures []Creature) {
	stateBytes, err := json.Marshal(creatures)
	if err != nil {
		log.Println(err)
	}
	wss.node.Publish(creatureSubscription, stateBytes)
}

func runWeb() {
	log.Printf("Starting server, visit http://localhost:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cred := &centrifuge.Credentials{
			UserID: "",
		}
		newCtx := centrifuge.SetCredentials(ctx, cred)
		r = r.WithContext(newCtx)
		h.ServeHTTP(w, r)
	})
}
