package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ksyoon0321/gomicro/registry"
	"github.com/ksyoon0321/gomicro/registry/adapter"
)

func main() {
	fmt.Println("init service registry")

	env := flag.String("env", "local", "current env")
	flag.Parse()
	godotenv.Load(".env." + *env)

	monid := os.Getenv("CH_MONITOR")
	fetcherid := os.Getenv("CH_FETCH")
	regserver := registry.NewServiceRegistry(registry.NewRegistryChannel(monid, fetcherid))

	httpAdap := adapter.NewHttpAdapter(os.Getenv("HTTP_ADDR"))
	redisAdap := adapter.NewRedisAdapter(os.Getenv("REDIS_CHAN"), os.Getenv("REDIS_ADDR"))

	regserver.RegistAdapter(httpAdap)
	regserver.RegistAdapter(redisAdap)
	regserver.Listen()
}
