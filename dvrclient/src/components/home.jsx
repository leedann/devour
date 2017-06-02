import React from "react"
import {store, setAll} from "./shared-state.js";
import Layout from "./Layout.jsx";
import { Link } from 'react-router-dom';
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));
const baseurl = 'https://dvrapi.leedann.me/v1/'

//The home page -- renders all of the tiles of the home page
export default class Home extends React.Component {
    constructor(props) {
        super(props);
        //not authorized
        if (!localStorage.getItem("Authorization")) {
            this.props.history.push('/')
        }
        this.state= store.getState()
    }
    componentDidMount() {
        this.unsub = store.subscribe(() => this.setState(store.getState()));
        var req = {
            method: 'GET',
            headers: myHeader
        }
        //need: upcoming, past, pending, recipe, diets, allergies, friends, favFriends
        fetch(baseurl+'events', req)
        .then((resp) => resp.json())
        .then((events) => {
            fetch(baseurl+'attendance', req)
            .then((resp) => resp.json())
            .then((pending) => {
                fetch(baseurl+'users/diets', req)
                .then((resp) => resp.json())
                .then((diet) => {
                    fetch(baseurl+'users/allergies', req)
                    .then((resp) => resp.json())
                    .then((allergies) => {
                        fetch(baseurl+'users/friends', req)
                        .then((resp) => resp.json())
                        .then((friends) => {
                            fetch(baseurl+'users/favorites', req)
                            .then((resp) => resp.json())
                            .then((favorites) => {
                                fetch(baseurl+'users/me',req)
                                .then((resp) => resp.json())
                                .then((user) => {
                                    console.log(events)
                                    store.dispatch(setAll({
                                        upcomingEvents: events.allEvents,
                                        pastEvents: events.pastEvents,
                                        recipeBook: [],
                                        diets: diet,
                                        allergies: allergies,
                                        friends: friends,
                                        favoriteFriends: favorites,
                                        pendingEvents: pending,
                                        user: user
                                    }))
                                })
                            })
                        })
                    })
                })
            })
        })
    }
    componentWillUnmount() {
        this.unsub();
    }

    render() {
        console.log(this.state)
        return(
            <Layout title="devour">
                <div className="mdl-grid">
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/recipes">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/spatula.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Explore
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/social">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/food.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Socialize
                                </div>
                            </div>
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/budget">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/restaurant-2.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Budget
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/plan">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/restaurant-5.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Plan
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/recipebook">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/app.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    My Recipe Book
                                </div>
                            </div>  
                        </Link>
                    </div>
                    <div className="homeWrapper mdl-grid mdl-cell--4-col mdl-cell--4-col-phone mdl-cell--4-col-tablet">
                        <Link to="/friends">
                            <div className="homeCards mdl-card mdl-shadow--4dp" >
                                <div className="mdl-card__title mdl-card--expand">
                                    <img src="/img/homeIco/svg/table.svg" className="homeIco" alt="exploreIcon"/>
                                </div>
                                <div className="mdl-card__actions menuItem">
                                    Friends
                                </div>
                            </div>  
                        </Link>
                    </div>
                </div>
        </Layout>
        );
    }
}

