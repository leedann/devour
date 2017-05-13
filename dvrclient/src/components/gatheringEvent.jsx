import React from "react"
import Layout from "./Layout.jsx";
import { Link } from 'react-router-dom';
import moment from "moment";

//redir if user is not authenticated without rendering anything
//append id to the end of this url
var d2 = new Date('12/12/2017');
var day = moment(d2).format("D");
var dayName = moment(d2).format("LLLL")
dayName = dayName.split(",", 1)
var month = moment(d2).format("MMMM");
var evnt2={
    id: 2,
    Name: "Dinner at Danny's",
    Description: "Come hang out at my house! No need to bring anything",
    Hosting: false,
    Time: "6:00pm",
    StartTime: d2
}
var going=["https://www.gravatar.com/avatar/0e58cf2f03c9a0a5d6965154837cd595", "https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50"]
//This is the page for a particular event
export default class GatheringPage extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            today: ""
        }
    }

    componentDidMount() {
    }

    render() {
        var doubleNum = day >= 10;
        return(
            <Layout title="social">
                <div className="eventWrapper mdl-grid mdl-layout">
                    <div className="mdl-card eventTitleCard mdl-shadow--2dp mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <img src="https://www.gravatar.com/avatar/b46909970f20fab89ca815d749ec80b3" alt="hostimg" className="usrPhoto" />
                        <p className="evtCardName">Danny</p>
                        <span className="evtCardSecondary textAccent textSecondary">host</span>
                    </div>
                    <div className="mdl-card eventTitleCard mdl-shadow--2dp mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                        <svg className="dayIcon eventTitleCircle">
                            <circle cx="45" cy="45" r="45" className="dayCircle" />
                            {doubleNum? <text className="evtCircleText textSecondary" x="26" y="56">{day}</text> : <text className="evtCircleText textSecondary" x="35" y="56">{day}</text>}
                        </svg>
                        <p className="evtCardName">{dayName}</p>
                        <span className="evtCardSecondary textAccent textSecondary">{month}</span>
                    </div>
                    <div className="mdl-card eventWideTitleCard mdl-shadow--3dp mdl-cell mdl-cell--12-col mdl-cell--8-col-phone">
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Name</span>
                        </div>
                        <span className="eventDesc">{evnt2.Name}</span>
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Description</span>
                        </div>
                        <span className="eventDesc">{evnt2.Description}</span>
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Time</span>
                        </div>
                        <span className="eventDesc">{evnt2.Time}</span>
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Friends</span>
                        </div>
                        <span className="eventDesc">
                            {going.map(usr =>
                            <img src={usr} alt="usrphoto" className="smallUsrPhoto"/>
                            )}
                        </span>
                    </div>
                    <Link to="/recipes">
                        <button className="recommendRecipes mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell-8-col-tablet">
                            Recommend Recipes
                        </button>
                    </Link>
                </div>

            </Layout>
        );
    }
}