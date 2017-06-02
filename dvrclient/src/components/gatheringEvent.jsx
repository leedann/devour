import React from "react"
import Layout from "./Layout.jsx";
import { Link } from 'react-router-dom';
import moment from "moment";
import {history} from './app.jsx';
import { Button, Popup, Form, Input } from 'semantic-ui-react'
import LoaderExampleLoader from './loading.jsx';
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));
const baseurl = 'https://dvrapi.leedann.me/v1/'

//TODO: redir if user is not authenticated without rendering anything
//append id to the end of this url
export default class GatheringPage extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            today: "",
            event: "",
            host: "",
            users: "",
            recipes: "",
        }
    }

    componentDidMount() {
        var id = this.props.match.params.eventid
        var req = {
            method: 'GET',
            headers: myHeader,
        }
        fetch(baseurl+'events/' + id, req)
        .then((resp) => resp.json())
        .then(data => {
            fetch(baseurl+'users/me', req)
                .then((resp) => resp.json())
                .then(me => {
                this.setState({
                    event: data.event,
                    host: data.host,
                    users: data.allUsers,
                    recipes: data.recipes,
                    me: me
                })
            })
        })
    }
    goToRecipe(e) {
        e.preventDefault()
        var location = {
            pathname: '/recipes/' + this.state.event.id,
            state: {event: this.state.event}
        }
        history.push(location)
    }
    deleteEv(e) {
        e.preventDefault()
            var req = {
                method: 'DELETE',
                headers: myHeader,
            }
            fetch(baseurl+'events/'+this.props.match.params.eventid, req)
            .then((resp) => {
                history.push('/home')
            })
            .catch((error) => {
                alert("There was something wrong with your request");
            })
    }
    handleSubmit(e) {
        e.preventDefault()
        let email = document.getElementById("form-input-control-friend-email");
        myHeader.append('Link', email.value);
        var data = {
                "email": email.value,
            }
            data = JSON.stringify(data);
            var req = {
                method: 'LINK',
                headers: myHeader,
                body: data
            }
            fetch(baseurl+'events/'+this.props.match.params.eventid, req)
            .then((resp) => resp.json())
            .then(data => console.log(data))
            .catch((error) => {
                alert("There was something wrong with your request");
            })

    }

    //Event page -- hard coded
    render() {
        if (!this.state.event) {
            return (<LoaderExampleLoader />);
        }
        var doubleNum = day >= 10;
        console.log(this.state)
        var startTime = this.state.event.startAt.substring(0, this.state.event.startAt.length - 1)
        var endTime = this.state.event.endAt.substring(0, this.state.event.endAt.length - 1)
        var day = moment(startTime).format("D");
        var dayName = moment(startTime).format("LLLL");
        dayName = dayName.split(",", 1)
        var month = moment(startTime).format("MMMM");
        var time = moment(startTime).format("LT");
        console.log(this.state)
        return(
            <Layout title="social">
                <div className="eventWrapper mdl-grid mdl-layout">
                    <div className="mdl-card eventTitleCard mdl-shadow--2dp mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <img src={this.state.host.photoURL} alt="hostimg" className="usrPhoto" />
                        <p className="evtCardName">{this.state.host.firstName}</p>
                        <span className="evtCardSecondary textAccent textSecondary">host</span>
                    </div>
                    <div className="mdl-card eventTitleCard mdl-shadow--2dp mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <svg className="dayIcon eventTitleCircle">
                            <circle cx="45" cy="45" r="45" className="dayCircle" />
                            {doubleNum? <text className="evtCircleText textSecondary" x="26" y="56">{day}</text> : <text className="evtCircleText textSecondary" x="25" y="56">{day}</text>}
                        </svg>
                        <p className="evtCardName">{dayName}</p>
                        <span className="evtCardSecondary textAccent textSecondary">{month}</span>
                    </div>
                    <div className="mdl-card eventWideTitleCard mdl-shadow--3dp mdl-cell mdl-cell--12-col mdl-cell--8-col-phone">
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Invite Some Friends! </span>
                        </div>
                        <Form className=" mdl-grid mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <Form.Field required maxLength='100' width={12} id='form-input-control-friend-email' control={Input} placeholder='email' />
                        
                        <Form.Button  size='large' onClick={(e) => this.handleSubmit(e)}>Invite!</Form.Button>
                        </Form>
                            
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Name</span>
                        </div>
                        <span className="eventDesc">{this.state.event.name}</span>
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Description</span>
                        </div>
                        <span className="eventDesc">{this.state.event.description}</span>
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Time</span>
                        </div>
                        <span className="timePadding"><p className="mdl-color-text--amber-700">Begins at:</p>{moment(startTime).format('LLLL')}</span>
                        <span className="timePadding"><p className="mdl-color-text--amber-700">Ends at: </p>{moment(endTime).format('LLLL')}</span>
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Friends</span>
                        </div>
                        <span className="eventDesc">
                            {/*gonna have to change this to friends or users*/}
                            {this.state.users.map(usr =>
                            <img key={usr.email} src={usr.photoURL} alt="usrphoto" className="smallUsrPhoto"/>
                            )}
                        </span>
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Recipes Suggested: </span>
                        <span className="eventDesc">
                            {this.state.recipes? this.state.recipes.map(food =>
                                <p className="eventDesc mdl-cell--12-col mdl-cell--8-col-phone">{food}</p>
                            ) : ""}
                        </span>
                        </div>
                    </div>
                    <button onClick={(e) => this.goToRecipe(e)}className="recommendRecipes mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell-8-col-tablet">
                        Recommend Recipes
                    </button>
                    {this.state.me.id === this.state.host.id ? 
                    <Popup
                        trigger={<Button icon='trash outline' label='Delete Event' color='red'/>}
                        content={<Button id={this.state.event.id} color='red' onClick={(event) => this.deleteEv(event)} content='Confirm' />}
                        on='click'
                        position='top right'
                    />
                    : ""}
                </div>
            </Layout>
        );
    }
}