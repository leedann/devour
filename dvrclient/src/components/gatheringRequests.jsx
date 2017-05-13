import React from "react";
import GRequestsList from "./gatheringRequest_tiles.jsx";
import { Link } from 'react-router-dom';
import Layout from "./Layout.jsx";

export default class GatheringRequests extends React.Component {
    render() {
        var d1 = new Date('09/09/2017');
        var d2 = new Date('12/12/2017');
        var evnt1={
            id: 1,
            Name: "Surprise Party For Morty",
            Hosting: false,
            StartTime: d1
        }
        var evnt2={
            id: 2,
            Name: "Dinner at Danny's",
            Hosting: false,
            StartTime: d2
        }
        var test = [evnt1, evnt2];
        return(
            <Layout title="social">
                <Link to="/social" className="eventLink">
                    <span className="eventsRedir textAccent">
                        <i className="eventsRedirButton material-icons">keyboard_backspace</i>
                        <p>Events</p>
                    </span>
                </Link>
                <GRequestsList events={test} />
            </Layout>
        );
    }
}