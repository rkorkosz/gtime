package app

func Run() {
	ac := NewAppContext()
	r := NewRouter(ac)
	RunServer(NewServer(r))
}
