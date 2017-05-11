import React from "react"
import Textfield from 'react-mdl/lib/Textfield';
import moment from "moment"
import Layout from "./Layout.jsx";

export default class CreateEvent extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            today: ""
        }
    }

    componentDidMount() {
    }

    render() {
        return(
            <Layout title="social">
                <div className="eventWrapper mdl-layout mdl-js-layout">
                    <span className="titleName mdl-cell--12-col mdl-cell--8-col-phone">create event</span>
                    <div className="eventCard mdl-shadow--2dp mdl-grid">
                        <div className="eventName mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Event Name</span>
                            <Textfield
                                className="eventTextField mdl-cell--12-col mdl-cell--8-col-phone"
                                onChange={() => {}}
                                label="Event Name"
                                style={{width: '25vw'}}
                            />
                        </div>
                        <div className="eventDesc mdl-grid mdl-cell mdl-cell--6-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                        <span className="eventHeader mdl-cell--12-col mdl-cell--8-col-phone">Description</span>
                            <Textfield
                                className="eventTextField mdl-cell--12-col mdl-cell--8-col-phone"
                                onChange={() => {}}
                                label="Important details go here!"
                                rows={3}
                                style={{width: '200px'}}
                            />
                        </div>
                        <div className="eventDate mdl-grid mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                            
                        </div>
                    </div>
                </div>
            </Layout>
        );
    }
}