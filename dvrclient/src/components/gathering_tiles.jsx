import React from "react";
import moment from "moment";
import {store} from "./shared-state.js";
import {history} from './app.jsx';
import LoaderExampleLoader from './loading.jsx';
import { Button, Popup } from 'semantic-ui-react'
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));
const baseurl = 'https://dvrapi.leedann.me/v1/';

//the long tiles
class LongCards extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            day: moment(this.props.evt.startAt).format("D"),
            hosting: "",
            name: this.props.evt.Name,
            request: this.props.request? this.props.request : false,
            event: ""
        }
        this.handleEvntClick = this.handleEvntClick.bind(this)
    }

    componentDidMount() {
        var req = {
            method: 'GET',
            headers: myHeader,
        }
        fetch(baseurl+'events/' + this.props.evt.id, req)
        .then((resp) => resp.json())
        .then(data => {
            this.setState({
                event: data.event,
                hosting: data.host
            })
        })
    }

    //clicking the event tile will send you to the particular event
    handleEvntClick(evt) {
        //state in push goes to this.location.state
        if (evt.target.tagName !== 'I') {
            //will need to render a different event based on referrer
            history.push("/social/event/" + this.props.evt.id)
        }
    }

    //checks to see if the longcards are of the request variety
    requestCheck() {
        var doubleNum = this.state.day >= 10
        if (this.state.request) {
            return (
                <li onClick={evt => this.handleEvntClick(evt)} className="listItem mdl-list__item mdl-list mdl-cell mdl-cell--4-col mdl-cell--8-col-phone mdl-cell--8-col-tablet mdl-shadow--2dp">
                    <span className="longCardContent mdl-list__item-primary-content">
                        <div className="dateTitleWrap">
                            <svg className="dayIcon">
                                <circle cx="50" cy="50" r="23" className="dayCircle" />
                                {/*resizing of the day icon based on the length of day*/}
                                {doubleNum? <text className="circleText" x="35" y="59">{this.state.day}</text> : <text className="circleText" x="43" y="59">{this.state.day}</text>}
                            </svg>
                            <span className="gatheringName">
                                {this.state.event.name}
                            </span>
                        </div>
                    {/*the accept and reject button*/}
                    {this.props.children}
                    </span>
                </li>
            );
        }else {
            //if the user is hosting, outline it
            var stroke = this.state.hosting? "#F50057" : "";
            return (
                <li onClick={evt => this.handleEvntClick(evt)} className="listItem mdl-list__item mld-list mdl-cell mdl-cell--4-col mdl-cell--8-col-phone mdl-cell--8-col-tablet mdl-shadow--2dp">
                    <span className="longCardContent mdl-list__item-primary-content">
                        <div className="dateTitleWrap">
                            <svg className="dayIcon">
                                <circle cx="50" cy="50" r="23" stroke={stroke}className="dayCircle" />
                                {doubleNum? <text className="circleText" x="35" y="59">{this.state.day}</text> : <text className="circleText" x="43" y="59">{this.state.day}</text>}
                            </svg>
                            <span className="gatheringName">
                                {this.state.event.name}
                            </span>
                        </div>
                    </span>
                </li>
            );
        }
    }
    render() {
        if (!this.state.event) {
            return (<LoaderExampleLoader />);
        }
        return(
            this.requestCheck()
        );
    }
}


//takes an array and a title and wraps the cards in a title (ie. events requests, or events/ possible friends?)
export class TitleWrap extends React.Component {
    constructor(props) {
        super(props)
        //not authorized
        if (!localStorage.getItem("Authorization")) {
            this.props.history.push('/')
        }
        this.state= store.getState()
    }
    componentDidMount() {
        this.unsub = store.subscribe(() => this.setState(store.getState()));
    }
    componentWillUnmount() {
        this.unsub();
    }
    //will need to rerender page to show the snackbar (weird)
    handleAccept(event) {
        var target = event.currentTarget;
        event.preventDefault()
        var data = {
            "eventid": target.id,
            "attendanceStatus": "Attending",
        }
        data = JSON.stringify(data);
        var req = {
            method: 'PATCH',
            headers: myHeader,
            body: data
        }
        fetch(baseurl+'attendance', req)
    }
    handleReject(event) {
        var target = event.currentTarget;
        event.preventDefault()
        var data = {
            "eventid": target.id,
            "attendanceStatus": "Not Attending",
        }
        data = JSON.stringify(data);
        var req = {
            method: 'PATCH',
            headers: myHeader,
            body: data
        }
        fetch(baseurl+'attendance', req)
    }


    render() {
        return (
                <div className="titleChildren mdl-grid">
                    <span className="titleName mdl-cell--12-col mdl-cell--8-col-phone">{this.props.titleName}</span>
                    <ul className="longCardContainer mld-list mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-grid">
                        {this.props.arr.map(evt =>
                            <LongCards key={evt.id} evt={evt} request={this.props.request}>
                                {this.props.request? 
                                    <span id={evt.id} className="reqAnswer mdl-list__item-secondary-action">
                                    <Popup
                                        trigger={<Button icon='checkmark' />}
                                        content={<Button id={evt.id} color='green' onClick={(event) => this.handleAccept(event)} content='Accept' />}
                                        on='click'
                                        position='top right'
                                    />
                                    <Popup
                                        trigger={<Button icon='remove'/>}
                                        content={<Button id={evt.id} color='red' onClick={(event) => this.handleReject(event)} content='Reject' />}
                                        on='click'
                                        position='top right'
                                    />
                                    </span>
                                : ""}
                            </LongCards>
                        )}
                    </ul>
                </div>            
            );
    }
}


//this is the events page list
export default class GatheringList extends React.Component {
    constructor(props) {
        super(props);
        //not authorized
        if (!localStorage.getItem("Authorization")) {
            this.props.history.push('/')
        }
        this.state= store.getState()
    }


    //compares the dates in the array of events... orders them by earliest to latest
    compareDate(rec1, rec2) {
        var first = rec1.startAt.substring(0, 10)
        var second = rec2.startAt.substring(0, 10)
        var date1 = new Date(first);
        var date2 = new Date(second);
        if (date1 > date2) {
            return 1;
        }
        if (date1 < date2) {
            return -1;
        }
        return 0;
    }

    componentDidMount() {
        this.unsub = store.subscribe(() => this.setState(store.getState()));
    }
    componentWillUnmount() {
        this.unsub();
    }

    render() {
        if (!this.state.upcomingEvents) {
            return <h2 className="mdl-layout__header mdl-layout__header-row mdl-layout__header--transparent textAccent ">no events to show!</h2>
        }
        var allMonths = this.state.upcomingEvents;
        allMonths.sort(this.compareDate)
        var eventObj={}
        //gets all the months of the events and orders them
        for (var i = 0; i < allMonths.length; i++) {
            var fix = allMonths[i].startAt.substring(0, allMonths[i].startAt.length - 1)
            var month = moment(fix).format("MMMM")
            //if month isnt in the obj
            if (!eventObj[month]) {
                eventObj[month]=[]
                eventObj[month].push(allMonths[i]);
            }else {
                eventObj[month].push(allMonths[i]);
            }
        }
        return (
            <div className="md-grid">
                {
                    Object.keys(eventObj).map(function(keyName, keyIndex) {
                        if (moment(keyName, "MMMM").isValid()) {
                            return <TitleWrap key={keyName} arr={eventObj[keyName]} titleName={keyName} />
                        }
                        return "";
                    })
                }
            </div>
        );
    }
}

