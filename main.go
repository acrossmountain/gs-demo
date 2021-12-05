package main

import (
	"log"

	_ "learn/controllers"
	_ "learn/modules/minio"
	_ "learn/services"

	"github.com/go-spring/spring-core/gs"
	_ "github.com/go-spring/starter-gin"
)

func main() {
	//gs.Setenv("GS_SPRING_PROFILES_ACTIVE", "test")
	//gs.Property("GS_SPRING_CONFIG_EXTENSIONS", ".properties,.prop,.yaml,.yml,.toml,.tml,.ini")
	//gs.Property("GO_SPRING_CONFIG_LOCATIONS", "")
	log.Fatal(gs.Run())
}
