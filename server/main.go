package server

type server struct{}

var ServerApi = new(server)

func Main() {
	go ServerApi.handel()
	ServerApi.messageHandle()
}
