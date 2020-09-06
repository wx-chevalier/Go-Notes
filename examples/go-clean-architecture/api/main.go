package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/L04DB4L4NC3R/jobs-mhrd/api/handler"
	"github.com/L04DB4L4NC3R/jobs-mhrd/api/middleware"
	"github.com/L04DB4L4NC3R/jobs-mhrd/pkg/admin"
	"github.com/L04DB4L4NC3R/jobs-mhrd/pkg/user"
	"github.com/gorilla/handlers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err.Error())
	}
	os.Setenv("secret", viper.GetString("jwt_secret"))
}

func dbConnect(host, port, user, dbname, password, sslmode string) (*gorm.DB, error) {

	// In the case of heroku
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	}
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode))

	return db, err
}

func main() {
	mode := viper.GetString("mode")

	// DB binding
	dbprefix := "database_" + mode
	dbhost := viper.GetString(dbprefix + ".host")
	dbport := viper.GetString(dbprefix + ".port")
	dbuser := viper.GetString(dbprefix + ".user")
	dbname := viper.GetString(dbprefix + ".dbname")
	dbpassword := viper.GetString(dbprefix + ".password")
	dbsslmode := viper.GetString(dbprefix + ".sslmode")

	db, err := dbConnect(dbhost, dbport, dbuser, dbname, dbpassword, dbsslmode)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	defer db.Close()
	log.Println("Connected to the database")

	// migrations
	db.AutoMigrate(&user.User{}, &admin.Admin{})

	// initializing repos and services
	userRepo := user.NewPostgresRepo(db)
	adminRepo := admin.NewPostgresRepo(db)

	userSvc := user.NewService(userRepo)
	adminSvc := admin.NewService(adminRepo)

	// Initializing handlers
	r := http.NewServeMux()

	handler.MakeUserHandler(r, userSvc)
	handler.MakeAdminHandler(r, adminSvc)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		return
	})

	// HTTP(s) binding
	serverprefix := "server_" + mode
	host := viper.GetString(serverprefix + ".host")
	port := os.Getenv("PORT")
	timeout := time.Duration(viper.GetInt("timeout"))

	if port == "" {
		port = viper.GetString(serverprefix + ".port")
	}

	conn := host + ":" + port

	// middlewares
	mwCors := middleware.CorsEveryWhere(r)
	mwLogs := handlers.LoggingHandler(os.Stdout, mwCors)

	srv := &http.Server{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Addr:         conn,
		Handler:      mwLogs,
	}

	log.Printf("Starting in %s mode", mode)
	log.Printf("Server running on %s", conn)
	log.Fatal(srv.ListenAndServe())
}
