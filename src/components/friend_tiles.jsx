import React from "react"

export default class Friend_Cards extends React.Component {
    constructor(props) {
        super(props)
    }

    render() {
        return(
                <div className="homeWrapper mdl-grid mdl-cell--12-col homeCard mdl-cell--8-col-phone">
                    <div className="mdl-card mdl-shadow--2dp" >
                        <p>Hello</p>
                    </div>  
                </div>
        );
    }
}