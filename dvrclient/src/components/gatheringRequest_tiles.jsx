import React from "react"
import {TitleWrap} from "./gathering_tiles.jsx"

export default class GRequestsList extends React.Component {
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

    componentDidMount(){
        var allMonths = this.state.events;
        allMonths.sort(this.compareDate)
        this.setState({
            events: allMonths
        })
    }

    render() {
        return(
            <TitleWrap key="requests" titleName="requests" arr={this.state.events} request={true}/>
        );
    }
}