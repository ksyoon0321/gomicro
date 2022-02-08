# MSA 기능은 최대한 간단하게

##Service Registry
 - Client Side
 - Register Adapter ( Redis Pub/Sub, HttpServer )

#
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
#

##Gateway ( grpc gateway, http restapi 2가지 혼용 v1 -> http, v2 -> grpc gateway)
 - Load Balancer (Round Robin, Failover)

##Log Service

##transaction Cordinator

##Order Service

##Stock Service

##Message Service (email?)
 
