import React from "react"
import {store} from "./shared-state.js";
import {TitleWrap} from "./gathering_tiles.jsx"

//this is the list of event requests
export default class GRequestsList extends React.Component {
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

    //renders the title and the tiles
    render() {
        if (!this.props.events) {
            return <h2 className="mdl-layout__header mdl-layout__header-row mdl-layout__header--transparent textAccent ">no pending invites</h2>
        }
        var allMonths = this.props.events;
        allMonths.sort(this.compareDate)
        return(
            <TitleWrap key="requests" titleName="requests" arr={allMonths} request={true}/>
        );
    }
}