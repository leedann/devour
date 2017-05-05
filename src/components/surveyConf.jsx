import React from "react"
import { Link } from 'react-router-dom';
export default class SurveyEnd extends React.Component {
    constructor(props) {
        super(props)
        this.state={
            
        }
    }
    render() {
        return(
            <div className="welcomePage mdl-layout mdl-js-layout">
                <div className="messageContainer mdl-layout mdl-js-layout">
                    <h3 className="textAccent">All Finished!</h3>
                    <p>visit settings in the main menu to edit your <span className="textAccent">lifestyle preferences</span></p>
                    <Link to='/recipes'><button className="dirToRecipes mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect">Explore Recipes</button></Link>
                </div>
            </div>
        );
    }
}