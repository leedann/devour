package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	redis "gopkg.in/redis.v5"

	"strings"

	"github.com/leedann/devour/devoursvr/handlers"
	"github.com/leedann/devour/devoursvr/middleware"
	"github.com/leedann/devour/devoursvr/models/users"
	"github.com/leedann/devour/devoursvr/sessions"
	_ "github.com/lib/pq"
)

const defaultPort = "443"

const (
	apiRoot    = "/v1/"
	apiSummary = apiRoot + "summary"
	pgPort     = 5432
	usr        = "users"
	sess       = "sessions"
	sessme     = "sessions/mine"
	usrme      = "users/me"
	channels   = "channels"
	specific   = "/"
	msgs       = "messages"
	diets      = "/diets"
	allergies  = "/allergies"
	friends    = "/friends"
	recipebook = "/recipebook"
	favorites  = "/favorites"
)

//main is the main entry point for this program
func main() {
	//read and use the following environment variables
	//when initializing and starting your web server
	// PORT - port number to listen on for HTTP requests (if not set, use defaultPort)
	// HOST - host address to respond to (if not set, leave empty, which means any host)

	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = defaultPort
	}
	HOST := os.Getenv("HOST")

	certPath := os.Getenv("TLSCERT")
	keyPath := os.Getenv("TLSKEY")

	SESSIONKEY := os.Getenv("SESSIONKEY")
	REDISADDR := os.Getenv("REDISADDR")
	DBADDR := os.Getenv("DBADDR")

	client := redis.NewClient(&redis.Options{
		Addr:     REDISADDR,
		Password: "",
		DB:       0,
	})
	pgAddr := strings.Split(DBADDR, ":")
	datasrcName := fmt.Sprintf("user=pgstest dbname=pg2 sslmode=disable host=%s port=%s", pgAddr[0], pgAddr[1])
	pgstore, err := sql.Open("postgres", datasrcName)

	_, err = pgstore.Exec("DELETE FROM users")
	_, err = pgstore.Exec("DELETE FROM user_diet_type")
	_, err = pgstore.Exec("DELETE FROM user_allergy_type")
	_, err = pgstore.Exec("DELETE FROM grocery_list")
	_, err = pgstore.Exec("DELETE FROM user_like_list")
	_, err = pgstore.Exec("DELETE FROM friends_list")
	_, err = pgstore.Exec("ALTER SEQUENCE users_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE user_diet_type_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE user_allergy_type_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE grocery_list_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE user_like_list_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE friends_list_id_seq RESTART")

	if err != nil {
		log.Fatalf("error starting db: %v", err)
	}

	usrStore := &users.PGStore{
		DB: pgstore,
	}
	//Pings the DB-- establishes a connection to the db
	err = pgstore.Ping()
	if err != nil {
		log.Fatalf("error pinging db %v", err)
	}
	redisStore := sessions.NewRedisStore(client, time.Hour*3600)

	//creating the starting table "general"
	ctx := &handlers.Context{
		SessionKey:   SESSIONKEY,
		SessionStore: redisStore,
		UserStore:    usrStore,
	}

	mux := http.NewServeMux()
	mux.HandleFunc(apiRoot+usr, ctx.UserHandler)
	mux.HandleFunc(apiRoot+sess, ctx.SessionsHandler)
	mux.HandleFunc(apiRoot+sessme, ctx.SessionsMineHandler)
	mux.HandleFunc(apiRoot+usrme, ctx.UsersMeHandler)
	mux.HandleFunc(apiRoot+usr+diets, ctx.UserDietHandler)
	mux.HandleFunc(apiRoot+usr+allergies, ctx.UserAllergyHandler)
	mux.HandleFunc(apiRoot+usr+friends, ctx.UserFriendsHandler)
	mux.HandleFunc(apiRoot+usr+recipebook, ctx.UserRecipesHandler)
	mux.HandleFunc(apiRoot+usr+favorites, ctx.UserFavoritesHandler)
	mux.HandleFunc(apiRoot+usr+diets+specific, ctx.SpecificDietHandler)
	mux.HandleFunc(apiRoot+usr+allergies+specific, ctx.SpecificAllergyHandler)
	mux.HandleFunc(apiRoot+usr+friends+specific, ctx.SpecificFriendHandler)
	mux.HandleFunc(apiRoot+usr+recipebook+specific, ctx.SpecificFavRecipeHandler)
	mux.HandleFunc(apiRoot+usr+favorites+specific, ctx.SpecificFavFriendHandler)
	mux.HandleFunc(apiSummary, handlers.SummaryHandler)
	http.Handle(apiRoot, middleware.Adapt(mux, middleware.CORS("", "", "", "")))

	//add your handlers.SummaryHandler function as a handler
	//for the apiSummary route
	//HINT: https://golang.org/pkg/net/http/#HandleFunc

	//start your web server and use log.Fatal() to log
	//any errors that occur if the server can't start
	//HINT: https://golang.org/pkg/net/http/#ListenAndServe
	addr := HOST + ":" + PORT
	fmt.Printf("listening at %s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, certPath, keyPath, nil))

}
