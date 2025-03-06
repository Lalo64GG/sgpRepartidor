package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lalo64/sgp/src/client/infraestructure/http/routes"
	supplier "github.com/lalo64/sgp/src/supplier/infraestructure/http/routes"
	delivery "github.com/lalo64/sgp/src/delivery/infraestructure/http/routes"
	driver "github.com/lalo64/sgp/src/deliverydriver/infraestructure/http/routes"
	"github.com/lalo64/sgp/src/database"
)

type Server struct {
	engine 		*gin.Engine
	http 		string
	port    	string
	httpAddr 	string
}


func NewServer(http, port string) Server {
	gin.SetMode(gin.ReleaseMode)

	srv := Server{
		engine: gin.Default(),
		http: http,
		port: port,
		httpAddr: http + ":" + port,
	}

	database.Connect()
	srv.engine.RedirectTrailingSlash = true
	srv.registerRoutes()

	return srv
}

func (s *Server) registerRoutes() {
	userRoutes := s.engine.Group("/v1/client")
	supplierRoutes := s.engine.Group("/v1/supplier")
	deliveryRoutes := s.engine.Group("/v1/delivery")
	driverRoutes := s.engine.Group("/v1/driver")

	supplier.SupplierRoutes(supplierRoutes)
	routes.ClientRoutes(userRoutes)
	delivery.DeliveryRoutes(deliveryRoutes)
	driver.DriverRoutes(driverRoutes)

}

func (s *Server) Run() error {
	log.Println("Server running on " + s.httpAddr)
	return s.engine.Run(s.httpAddr)
}