import React from "react";
import { Link } from 'react-router-dom';

//the routes from welcome all the way to the end of the survey screens
const surveyRoute = ["/welcome", "/survey", "/selectgoals", "/allergyinfo", "/surveyend"];
//The footer of every page
export default class Footer extends React.Component {
    constructor(props) {
        super(props);
        this.state={
            left: this.props.page[0],
            right: this.props.page[1],
            title: this.props.title,
            leftAction: "/home",
            rightAction: "/home"
        }
    }

    //Gets the next and previous routes for the bottom nav
    componentDidMount() {
        if (this.state.title === "taste profile") {
            var index = surveyRoute.indexOf(location.pathname);
            this.setState({
                leftAction: surveyRoute[index-1],
                rightAction: surveyRoute[index+1]
            })
        }
    }

    //TODO: helper for the upcoming and past buttons at the bottom
    surveyHelper() {

    }

    //a blank footer for pages that do not require any buttons
    blankFooter() {
        return(
            <div className="phantom">
                <footer className="devourFooter mdl-mini-footer mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                    <div className="footLeft mdl-grid mdl-mini-footer__left-section">
                    </div>
                    <div className="centerFoot mdl-grid mdl-mini-footer__left-section">
                    </div>
                    <div className="footRight mdl-grid mdl-mini-footer__right-section">
                    </div>
                </footer>
            </div>
        );
    }

    //buttons for the survey -- will use the nav keywords to link to pages forwards and backwards
    surveyFooter() {
        return (
            <div className="phantom">
                <footer className="devourFooter mdl-mini-footer mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                    <div className="footLeft mdl-grid mdl-mini-footer__left-section">
                        <Link to={this.state.leftAction}><i className="paginate material-icons">chevron_left</i></Link>
                    </div>
                    <div className="centerFoot mdl-grid mdl-mini-footer__left-section">
                        <Link to='/home'><i className="homeIcon material-icons">restaurant</i></Link>
                    </div>
                    <div className="footRight mdl-grid mdl-mini-footer__right-section">
                        <Link to={this.state.rightAction}><i className="paginate material-icons">chevron_right</i></Link>
                    </div>
                </footer>
            </div>
        );
    }

    //split footer-- upcoming and past 
    splitFooter() {
        return (
            <div className="phantom">
                <footer className="devourFooter mdl-mini-footer mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                    <div className="footLeft mdl-grid mdl-mini-footer__left-section">
                        <button id="upcomingButton" className="mdl-cell--stretch mdl-button mdl-js-button mdl-button--accent">
                            <span>{this.state.left}</span>
                        </button>
                    </div>
                    <div className="centerFoot mdl-grid mdl-mini-footer__left-section">
                        <Link to='/home'><i className="homeIcon material-icons">restaurant</i></Link>
                    </div>
                    <div className="footRight mdl-grid mdl-mini-footer__right-section">
                        <button id="pastButton" className="mdl-cell--stretch mdl-button mdl-js-button mdl-button--accent">
                            <span>{this.state.right}</span>
                        </button>
                    </div>
                </footer>
            </div>
        );
    }


//renders a different footer based on the title of the state (the title is the main page section)
render() {
        switch (this.state.title) {
            case "devour":
                return this.blankFooter();
            case "taste profile":
                return this.surveyFooter();
            case "recipes":
                return (<span></span>);
            case "social":
                return this.splitFooter();
            case "recipe book":
                return this.blankFooter();
            case "budget":
                return this.splitFooter()
            case "plan":
                return this.splitFooter();
            default:
                break;
        }
    }
}