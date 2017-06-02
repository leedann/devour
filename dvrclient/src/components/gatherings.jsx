import React from "react"
import Layout from "./Layout.jsx";
import {store} from "./shared-state.js";
import GatheringList from "./gathering_tiles.jsx";
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));
// import moment from "moment";

//TODO:
// the views will change from past to upcoming-- need to get events and sort events
// pass in params to the gather tiles based on which 'view'

//the gatherings page
export default class Gatherings extends React.Component {
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
        return(
            <Layout title="social">
                <GatheringList />
            </Layout>
        );
    }
}