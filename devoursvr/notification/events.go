package notification

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/leedann/devour/devoursvr/models/events"
	"github.com/leedann/devour/devoursvr/models/users"
)

//NewUser keyword for a new user event
const NewUser = "NewUser"

//NewEvent keyword for a new event
const NewEv = "NewEvent"

//UpdateEv keyword is for when an event is updated
const UpdateEv = "UpdateEvent"

//DeleteEv keyword for deleting an event
const DeleteEv = "DeleteEvent"

//InviteEv keyword for inviting a user to an invite
const InviteEv = "InviteEvent"

//RejectEv keyword for removing a user from an event
const RejectEv = "RejectEvent"

//UpdateAttendance keyword for updating the attendance to an event
const UpdateAttendance = "UpdateAttendance"

//AddRecipe keyword for adding recipe to event
const AddRecipe = "AddRecipe"

//RemoveRecipe keyword for removing a recipe from an event
const RemoveRecipe = "RemoveRecipe"

//NewDiet keyword for user adding new diet
const NewDiet = "NewDiet"

//RemoveDiet keyword for user removing a diet
const RemoveDiet = "RemoveRecipe"

//NewAllergy keyword for user adding a new allergy
const NewAllergy = "NewAllergy"

//RemoveAllergy keyword for user removing an allergy
const RemoveAllergy = "RemoveAllergy"

//NewBook keyword for user adding a recipe to recipe book
const NewBook = "NewBook"

//RemoveBook keyword for removing a recipe from recipe book
const RemoveBook = "RemoveBook"

//AddFriend keyword for adding a new friend
const AddFriend = "AddFriend"

//RemoveFriend keyword for removing a friend
const RemoveFriend = "RemoveFriend"

//AddFavFriend keyword for adding a friend as favorite
const AddFavFriend = "AddFavFriend"

//RemoveFavFriend keyword for removing a friend as favorite
const RemoveFavFriend = "RemoveFavFriend"

//UserEvent struct defines all of the events that have to do with users
type UserEvent struct {
	EventType string      `json:"eventType"`
	Message   *users.User `json:"user"`
}

//RecipesEvent struct defines all of the events that have to do with recipes
type RecipesEvent struct {
	EventType string `json:"eventType"`
	Message   string `json:"recipe"`
}

//DietEvent struct defines all of events that have to do with adding or removing diets
type DietEvent struct {
	EventType string   `json:"eventType"`
	Message   []string `json:"diet"`
}

//AllergyEvent struct defines all of events that have to do with adding or removing allergies
type AllergyEvent struct {
	EventType string   `json:"eventType"`
	Message   []string `json:"allergy"`
}

//EvtEvent are all the events that have to do with irl events
// Message is the data and you select it in the client via the json name
type EvtEvent struct {
	EventType string        `json:"eventType"`
	Message   *events.Event `json:"event"`
}

//FmtEvtEvent are all the events that have to do with irl events
// Message is the data and you select it in the client via the json name
type FmtEvtEvent struct {
	EventType string           `json:"eventType"`
	Message   *events.FmtEvent `json:"event"`
}

//Notifier represents a web sockets notifier
type Notifier struct {
	eventq  chan interface{}
	clients map[*websocket.Conn]bool
	mu      sync.RWMutex
}

//NewNotifier constructs a new Notifer.
func NewNotifier() *Notifier {
	return &Notifier{
		eventq:  make(chan interface{}, 100),
		clients: make(map[*websocket.Conn]bool),
		mu:      sync.RWMutex{},
	}
}

//Start begins a loop that checks for new events
//and broadcasts them to all web socket clients.
//This function should be called on a new goroutine
//e.g., `go mynotifer.Start()`
func (n *Notifier) Start() {
	for {
		select {
		case evnt := <-n.eventq:
			n.broadcast(evnt)
		}
	}
}

//AddClient adds a new web socket client to the Notifer
func (n *Notifier) AddClient(client *websocket.Conn) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.clients[client] = true
	go n.readPump(client)
}

//Notify will add a new event to the event queue
func (n *Notifier) Notify(event interface{}) {
	//TODO: add the `event` to the `eventq`
	n.eventq <- event
}

//Control messages
func (n *Notifier) readPump(client *websocket.Conn) {
	for {
		if _, _, err := client.NextReader(); err != nil {
			client.Close()
			break
		}
	}

}

//broadcast sends the event to all client as a JSON-encoded object
func (n *Notifier) broadcast(event interface{}) {
	n.mu.Lock()
	defer n.mu.Unlock()
	evt, err := json.Marshal(event)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	pm, err := websocket.NewPreparedMessage(websocket.TextMessage, evt)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	for key := range n.clients {
		err := key.WritePreparedMessage(pm)
		fmt.Println(err)
		if err != nil {
			delete(n.clients, key)
			key.Close()
		}
	}
}
