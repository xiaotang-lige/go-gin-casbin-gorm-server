package server

type server struct {
	UserConfig
	Contacts *contacts
}

var Api = new(server)

func Main() {
	go Api.handel()
	go Api.messageHandle()
}
