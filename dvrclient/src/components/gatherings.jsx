import React from "react"
import Layout from "./Layout.jsx";
import GatheringList from "./gathering_tiles.jsx"
//TODO:
// the views will change from past to upcoming-- need to get events and sort events
// pass in params to the gather tiles based on which 'view'
export default class Gatherings extends React.Component {
    constructor(props) {
        super(props)
    }

    render() {
        var d1 = new Date('07/09/2017');
        var d2 = new Date('07/12/2017');
        var d3 = new Date('09/20/2017');
        var d4 = new Date('09/26/2017');
        var d5 = new Date('10/15/2017');
        var evnt1={
            id: 1,
            Name: "Girl's Night Out",
            Hosting: false,
            StartTime: d1
        }
        var evnt2={
            id: 2,
            Name: "Grandma's Bday",
            Hosting: true,
            StartTime: d2
        }
        var evnt3={
            id: 3,
            Name: "Brunch with the fam",
            Hosting: false,
            StartTime: d3
        }
        var evnt4={
            id: 4,
            Name: "Movie Night",
            Hosting: false,
            StartTime: d5
        }
        var evnt5={
            id: 5,
            Name: "Eating alone",
            Hosting: true,
            StartTime: d1
        }
        var evnt6={
            id: 6,
            Name: "Dinner with bae",
            Hosting: false,
            StartTime: d4
        }
        var test = [evnt2, evnt3, evnt1, evnt4, evnt5, evnt6];
        return(
            <Layout title="social">
                <div className="gatheringWrapper mdl-layout mdl-js-layout">
                    <GatheringList events={test} />
                </div>
            </Layout>
        );
    }
}