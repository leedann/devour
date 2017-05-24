import React from "react"
import Layout from "./Layout.jsx";
import GatheringList from "./gathering_tiles.jsx";
// import moment from "moment";

//TODO:
// the views will change from past to upcoming-- need to get events and sort events
// pass in params to the gather tiles based on which 'view'
const d1 = new Date('07/09/2016');
const d2 = new Date('07/12/2016');
const d3 = new Date('09/20/2017');
const d4 = new Date('09/26/2017');
const d5 = new Date('10/15/2017');
const d6 = new Date('12/12/2017');
const evnt1={
    id: 1,
    Name: "Girl's Night Out",
    Hosting: false,
    StartTime: d1
}
const evnt2={
    id: 2,
    Name: "Grandma's Bday",
    Hosting: true,
    StartTime: d2
}
const evnt3={
    id: 3,
    Name: "Brunch with the fam",
    Hosting: false,
    StartTime: d3
}
const evnt4={
    id: 4,
    Name: "Movie Night",
    Hosting: false,
    StartTime: d5
}
const evnt5={
    id: 5,
    Name: "Eating alone",
    Hosting: true,
    StartTime: d1
}
const evnt6={
    id: 6,
    Name: "Dinner with bae",
    Hosting: false,
    StartTime: d4
}
const evnt7={
    id: 7,
    Name: "Dinner at Danny's",
    Hosting: false,
    StartTime: d6
}
const test = [evnt2, evnt3, evnt1, evnt4, evnt5, evnt6, evnt7];

//the gatherings page
export default class Gatherings extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            events: test
        }
    }

    componentDidMount() {
    }

    filterFuture() {
        var today = new Date();
        var filteredTest = test.filter((evnt) => {
            return (today > evnt.StartTime)
        });
        this.setState({
            events: filteredTest
        });
    }

    filterPast() {
        var today = new Date();
        var filteredTest = test.filter((evnt) => {
            return (today < evnt.StartTime)
        });
        this.setState({
            events: filteredTest
        });
    }
    render() {
        return(
            <Layout title="social">
                <GatheringList events={this.state.events} />
            </Layout>
        );
    }
}