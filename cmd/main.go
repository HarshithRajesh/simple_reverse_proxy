package main
import (
	"log"
	"github.com/HarshithRajesh/reverse_proxy_server/internal/server"
)

func main(){
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server: %v",err)
	}
}