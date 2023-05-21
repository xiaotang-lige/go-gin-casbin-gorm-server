package server

type server struct {
	UserConfig
}

var ServerApi = new(server)

func Main() {
	go ServerApi.handel()
	go ServerApi.messageHandle()
}
