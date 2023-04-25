package main

import "server/server"

func main() {
	srv := server.New(":8060")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}