import React from "react"
import { Link } from 'react-router-dom';
//the welcoming page shown after registration
export default class Welcome extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            
        }
    }
    render() {
        return(
            <div className="welcomePage mdl-layout mdl-js-layout">
                <div className="messageContainer mdl-layout mdl-js-layout">
                    <h3 className="textAccent">Welcome</h3>
                    <p>take the lifestyle assessment to help us recommend recipes for <span className="textAccent">you and your friends!</span></p>
                    <Link to='/survey'><button className="beginSurveyButton mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect">let's go!</button></Link>
                    <Link to='/home'><div className="skipSurvey">skip survey</div></Link>
                </div>
            </div>
        );
    }
}