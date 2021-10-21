package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-postgres/pkg/db"
	"go-postgres/pkg/middleware"
	"go-postgres/pkg/router"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("configs"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Не удалось считать конфиг файл: %s \n", err)
	}

	handler := middleware.Handler{Repo: db.PostgresRepo{ConnString: viper.GetString("postgres.connString")}}
	r := router.Router(handler)

	hostPort := fmt.Sprintf("%v:%v", viper.GetString("webServer.host"), viper.GetString("webServer.port"))
	fmt.Printf("Starting server on the port %v...", viper.GetString("webServer.port"))
	log.Fatal(http.ListenAndServe(hostPort, r))
}
