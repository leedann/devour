import React from "react"
import GRequestsList from "./gatheringRequest_tiles.jsx"
import Layout from "./Layout.jsx";

export default class GatheringRequests extends React.Component {
    constructor(props) {
        super(props)
    }

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
            Name: "Beach Bonfire Dinna Dinna",
            Hosting: true,
            StartTime: d2
        }
        var test = [evnt1, evnt2];
        return(
            <Layout title="social">
                <div className="gatheringWrapper mdl-layout mdl-js-layout">
                    <GRequestsList events={test} />
                    <div>
                        <div id="snackConfirm" className="mdl-js-snackbar mdl-snackbar">
                        <div className="mdl-snackbar__text"></div>
                            <button className="mdl-snackbar__action" type="button"></button>
                        </div>
                    </div>
                </div>
            </Layout>
        );
    }
}