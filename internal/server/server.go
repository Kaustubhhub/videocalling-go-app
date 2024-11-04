package server

import (
	"flag"
	"os"
)

var (
	addr = flag.String("addr", os.Getenv("PORT"), "Server address")
	cert = flag.String("cert", "", "Path to SSL certificate")
	key  = flag.String("key", "", "Path to SSL key")
)

func Run() {
	flag.Parse()

	if *addr == ":" {
		*addr = ":8080"
	}

	app.Get("/", handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocket")
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket")
	app.Get("/stream/:ssuid/chat/websocket")
	app.Get("/stream/:ssuid/viewer/websocket")
}
