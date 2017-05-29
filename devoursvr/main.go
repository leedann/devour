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
	"github.com/leedann/devour/devoursvr/models/events"
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
	specific   = "/"
	diets      = "/diets"
	allergies  = "/allergies"
	friends    = "/friends"
	recipebook = "/recipebook"
	favorites  = "/favorites"
	event      = "events"
	recipes    = "/recipes"
	attendance = "attendance"
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
	datasrcName := fmt.Sprintf("user=pgstest dbname=devourpg sslmode=disable host=%s port=%s", pgAddr[0], pgAddr[1])
	pgstore, err := sql.Open("postgres", datasrcName)

	_, err = pgstore.Exec("DELETE FROM users")
	_, err = pgstore.Exec("DELETE FROM user_diet_type")
	_, err = pgstore.Exec("DELETE FROM user_allergy_type")
	_, err = pgstore.Exec("DELETE FROM grocery_list")
	_, err = pgstore.Exec("DELETE FROM user_like_list")
	_, err = pgstore.Exec("DELETE FROM friends_list")
	_, err = pgstore.Exec("DELETE FROM event_attendance")
	_, err = pgstore.Exec("DELETE FROM events")
	_, err = pgstore.Exec("DELETE FROM recipe_suggestions")
	_, err = pgstore.Exec("ALTER SEQUENCE users_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE user_diet_type_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE user_allergy_type_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE grocery_list_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE user_like_list_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE friends_list_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE event_attendance_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE events_id_seq RESTART")
	_, err = pgstore.Exec("ALTER SEQUENCE recipe_suggestions_id_seq RESTART")

	if err != nil {
		log.Fatalf("error starting db: %v", err)
	}

	usrStore := &users.PGStore{
		DB: pgstore,
	}
	evtStore := &events.PGStore{
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
		EventStore:   evtStore,
	}

	mux := http.NewServeMux()
	//users
	mux.HandleFunc(apiRoot+usr, ctx.UserHandler)
	//sessions
	mux.HandleFunc(apiRoot+sess, ctx.SessionsHandler)
	//my session
	mux.HandleFunc(apiRoot+sessme, ctx.SessionsMineHandler)
	//my user
	mux.HandleFunc(apiRoot+usrme, ctx.UsersMeHandler)
	//my diet
	mux.HandleFunc(apiRoot+usr+diets, ctx.UserDietHandler)
	//my allergies
	mux.HandleFunc(apiRoot+usr+allergies, ctx.UserAllergyHandler)
	//my friends
	mux.HandleFunc(apiRoot+usr+friends, ctx.UserFriendsHandler)
	//my recipes
	mux.HandleFunc(apiRoot+usr+recipebook, ctx.UserRecipesHandler)
	//my favorite friends
	mux.HandleFunc(apiRoot+usr+favorites, ctx.UserFavoritesHandler)
	//Specific diet
	mux.HandleFunc(apiRoot+usr+diets+specific, ctx.SpecificDietHandler)
	//Specific allergy
	mux.HandleFunc(apiRoot+usr+allergies+specific, ctx.SpecificAllergyHandler)
	//Specific friend
	mux.HandleFunc(apiRoot+usr+friends+specific, ctx.SpecificFriendHandler)
	//Specific favorite recipe
	mux.HandleFunc(apiRoot+usr+recipebook+specific, ctx.SpecificFavRecipeHandler)
	//Specific favorite friend
	mux.HandleFunc(apiRoot+usr+favorites+specific, ctx.SpecificFavFriendHandler)
	//events
	mux.HandleFunc(apiRoot+event, ctx.EventsHandler)
	//Specific events
	mux.HandleFunc(apiRoot+event+specific, ctx.SpecificEventsHandler)
	//my attendances
	mux.HandleFunc(apiRoot+attendance, ctx.EventAttendanceHandler)
	//my recipes (adding recipes to an event)
	mux.HandleFunc(apiRoot+event+recipes+specific, ctx.EventRecipesHandler)

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
