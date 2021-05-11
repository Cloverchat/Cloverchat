package main

func main() {
	cloverchat := &Server{
		Address: ":8080",
	}
	cloverchat.Run()
}

