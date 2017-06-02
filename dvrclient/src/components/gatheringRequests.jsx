import React from "react";
import GRequestsList from "./gatheringRequest_tiles.jsx";
import {store} from "./shared-state.js";
import { Link } from 'react-router-dom';
import Layout from "./Layout.jsx";
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));

//The actually events requests the user has
export default class GatheringRequests extends React.Component {
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
    }
    componentWillUnmount() {
        this.unsub();
    }
    render() {
        console.log(this.state)
        return(
            <Layout title="social">
                <Link to="/social" className="eventLink">
                    <span className="eventsRedir textAccent">
                        <i className="eventsRedirButton material-icons">keyboard_backspace</i>
                        <p>Events</p>
                    </span>
                </Link>
                <GRequestsList events={this.state.pendingEvents} />
            </Layout>
        );
    }
}