import {createStore} from "redux";
import update from 'immutability-helper';


const LS_KEY = "devour-store";
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

const SET_ALL = "setall";
// localStorage.removeItem(LS_KEY)
const default_case = {user: "", friends: [], upcomingEvents: [], pastEvents: [], recipeBook: [], diets: [], allergies: [], favoriteFriends: [], pendingEvents: []};
var savedState = JSON.parse(localStorage.getItem(LS_KEY));
function reducer(state, action) {
    switch(action.type) {
        case SET_ALL:
            var newUpc = action.obj.upcomingEvents
            if (action.obj.pendingEvents) {
                for (var i=0;i<action.obj.pendingEvents.length;i++) {
                    newUpc = newUpc.filter((evt) => {
                        return evt.id !== action.obj.pendingEvents[i].id
                    })
                }
            }
            return Object.assign({}, state, {friends: action.obj.friends, upcomingEvents: newUpc, pastEvents: action.obj.pastEvents, recipeBook: action.obj.recipeBook, diets: action.obj.diets, allergies: action.obj.allergies, favoriteFriends: action.obj.favoriteFriends, pendingEvents: action.obj.pendingEvents, user: action.obj.user})
        case NewEv:
            //creation of new event, should add to upcoming
            var newUpcoming = state.upcomingEvents
            newUpcoming.push(action.mess)
            return Object.assign({}, state, {upcomingEvents: newUpcoming});
            break;
        case DeleteEv:
            //creation of new event, should add to upcoming
            var nu = state.upcomingEvents
            nu = nu.filter((evt) => {
                return evt.id !== action.mess.id
            })
            return Object.assign({}, state, {upcomingEvents: nu});
        case UpdateAttendance: //adding to upcoming
            var newPending = state.pendingEvents
            newPending = newPending.filter((evt) => {
                return evt.id !== action.mess.id
            })
            var newUpc = state.upcomingEvents
            newUpc.push(action.mess)
            console.log(newUpc)
            return Object.assign({}, state, {pendingEvents: newPending, upcomingEvents: newUpc});
        case RejectEv:
            var newPending = state.pendingEvents
            newPending = newPending.filter((evt) => {
                return evt.id !== action.mess.id
            })
            return Object.assign({}, state, {pendingEvents: newPending});
        case NewBook:
            var upBook = state.recipeBook
            upBook.push(action.mess)
            return Object.assign({}, state, {recipeBook: upBook});
        case RemoveBook:
            upBook = state.recipeBook
            upBook = upBook.filter((recipe) => {
                return recipe !== action.mess
            })
            return Object.assign({}, state, {recipeBook: upBook});
        case AddFriend:
            var newFriend = state.friends
            newFriend.push(action.mess)
            return Object.assign({}, state, {friends: newFriend});
        case AddFavFriend:
            var newFav = state.favoriteFriends
            newFav.push(action.mess)
            return Object.assign({}, state, {favoriteFriends: newFav});
        case RemoveFriend:
            var rmf = state.friends
            rmf = rmf.filter((usr) => {
                return usr.id !== action.mess.id
            })
            return Object.assign({}, state, {friends: rmf});
        case RemoveFavFriend:
            var rmFav = state.favoriteFriends
            rmFav = rmFav.filter((usr) => {
                return usr.id !== action.mess.id
            })
            return Object.assign({}, state, {favoriteFriends: rmFav});
        case AddRecipe:
            //really dont need to do anything here
            return state;
        case RemoveRecipe:
            //no need for a rerender
            return state;
        case InviteEv:
            //no need
            return state;
        default:
            //do nothing
            return state;
    }
}

export function handleResponse(event) {
    var mess;
    if (event.event) {
        mess = event.event
    }else if (event.user) {
        mess = event.user
    }else if (event.recipe) {
        mess = event.recipe
    }else if (event.diet) {
        mess = event.diet
    }else if (event.allergy) {
        mess = event.allergy
    }
    return {
        type: event.eventType,
        mess: mess
    }
}

export function setAll(obj) {
    return {
        type: SET_ALL,
        obj: obj
    }
}



export var store = createStore(reducer, savedState || default_case);
store.subscribe(() => localStorage.setItem(LS_KEY, JSON.stringify(store.getState())));