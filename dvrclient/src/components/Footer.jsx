import React from "react";
import { Link } from 'react-router-dom';

const surveyRoute = ["/welcome", "/survey", "/selectgoals", "/allergyinfo", "/surveyend"];
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

    componentDidMount() {
        if (this.state.title === "taste profile") {
            var index = surveyRoute.indexOf(location.pathname);
            this.setState({
                leftAction: surveyRoute[index-1],
                rightAction: surveyRoute[index+1]
            })
        }
    }

    surveyHelper() {

    }

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