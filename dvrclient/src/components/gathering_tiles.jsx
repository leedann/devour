import React from "react";
import moment from "moment";
import {history} from './app.jsx'

//the long tiles
class LongCards extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            day: moment(this.props.evt.StartTime).format("D"),
            hosting: this.props.evt.Hosting,
            name: this.props.evt.Name,
            request: this.props.request? this.props.request : false
        }
        this.handleEvntClick = this.handleEvntClick.bind(this)
    }

    handleEvntClick(evt) {
        //state in push goes to this.location.state
        if (evt.target.tagName !== 'I') {
            history.push("/social/event")
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
                                {doubleNum? <text className="circleText" x="35" y="59">{this.state.day}</text> : <text className="circleText" x="43" y="59">{this.state.day}</text>}
                            </svg>
                            <span className="gatheringName">
                                {this.state.name}
                            </span>
                        </div>
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
                                {this.state.name}
                            </span>
                        </div>
                    </span>
                </li>
            );
        }
    }
    render() {
        return(
            this.requestCheck()
        );
    }
}


//takes an array and a title and wraps the cards in a title
export class TitleWrap extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            titleName: this.props.titleName,
            arr: this.props.arr,
            request: this.props.request
        }
    }
    //will need to rerender page to show updated
    handleRequestClick(event) {
        var target = event.currentTarget;
        var buttonClass = target.className.split(" ");
        var targetID = target.parentNode.id;
        var snackbarContainer = document.getElementById("snackBar");
        var data = {
            message: "",
            timeout: 1000,
            actionText: 'Undo'
        };
        switch (buttonClass[0]) {
            case "accept":
                    target.style.color="#4CAF50";
                    data.message = "Event accepted!"
                    snackbarContainer.MaterialSnackbar.showSnackbar(data);
                break;
            case "reject":
                    target.style.color="#F50057";
                    data.message ="Event declined."
                    snackbarContainer.MaterialSnackbar.showSnackbar(data);
                break;
            default:
                break;
        }
        //temporary remove -- will need to remove from api then 
        var newArr = this.state.arr;
        newArr = newArr.filter(obj => {
            return obj.id != targetID
        })
        this.setState({
            arr: newArr
        })
        //dont forget to do an api call when we get it up!
    }

    render() {
        return (
                <div className="titleTile">
                    <div className="titleChildren mdl-grid">
                        <span className="titleName mdl-cell--12-col mdl-cell--8-col-phone">{this.state.titleName}</span>
                        <ul className="longCardContainer mld-list mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-grid">
                            {this.state.arr.map(evt =>
                                <LongCards key={evt.id} evt={evt} request={this.props.request}>
                                    {this.props.request? 
                                        <span id={evt.id} className="reqAnswer mdl-list__item-secondary-action">
                                            <i className="reject material-icons" onClick={event => this.handleRequestClick(event)}>clear</i>
                                            <i className="accept material-icons" onClick={event => this.handleRequestClick(event)}>done</i>
                                        </span>
                                    : ""}
                                </LongCards>
                            )}
                        </ul>
                    </div>
                </div>              
            );
    }
}


//this is the events page list
export default class GatheringList extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            events: this.props.events,
        }
    }


    //compares the dates in the array of events... orders them by earliest to latest
    compareDate(rec1, rec2) {
        var date1 = new Date(rec1.StartTime);
        var date2 = new Date(rec2.StartTime);
        if (date1 > date2) {
            return 1;
        }
        if (date1 < date2) {
            return -1;
        }
        return 0;
    }

    //re-orders the array to make sure that they are sorted correctly
    //TODO: work with this and actual data from an ajax call
    componentDidMount() {
        var allMonths = this.state.events;
        allMonths.sort(this.compareDate)
        var eventObj={}
        for (var i = 0; i < allMonths.length; i++) {
            var month = moment(allMonths[i].StartTime).format("MMMM")
            //if month isnt in the obj
            if (!eventObj[month]) {
                eventObj[month]=[]
                eventObj[month].push(allMonths[i]);
            }else {
                eventObj[month].push(allMonths[i]);
            }
        }
        this.setState({
            events: eventObj
        })
    }

    clickAction(e) {
        e.preventDefault()
    }

    render() {
        var eventObj = this.state.events
        return (
            <div className="md-grid">
                {
                    Object.keys(this.state.events).map(function(keyName, keyIndex) {
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

