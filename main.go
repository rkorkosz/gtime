package main

func main() {
	ac := NewAppContext()
	r := NewRouter(ac)
	RunServer(NewServer(r))
}
