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
        if (this.state.title == "taste profile") {
            var index = surveyRoute.indexOf(location.pathname);
            this.setState({
                leftAction: surveyRoute[index-1],
                rightAction: surveyRoute[index+1]
            })
        }
    }

    surveyHelper() {

    }

    surveyFooter() {
        return (
            <footer className="devourFooter mdl-mini-footer mdl-layout mdl-js-layout mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                <div className="footLeft mdl-grid mdl-mini-footer__left-section">
                    <Link to={this.state.leftAction}><i className="paginate material-icons">chevron_left</i></Link>
                </div>
                <div className="mdl-layout-spacer"></div>
                <Link to='/home'><i className="homeIcon material-icons">restaurant</i></Link>
                <div className="mdl-layout-spacer"></div>
                <div className="footRight mdl-grid mdl-mini-footer__right-section">
                    <Link to={this.state.rightAction}><i className="paginate material-icons">chevron_right</i></Link>
                </div>
            </footer>
        );
    }

    splitFooter() {
        return (
            <footer className="devourFooter mdl-mini-footer mdl-layout mdl-js-layout mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                <div className="footLeft mdl-grid mdl-mini-footer__left-section">
                    <button className="mdl-cell--stretch mdl-button mdl-js-button mdl-button--accent">
                        <span>{this.state.left}</span>
                    </button>
                </div>
                <div className="mdl-layout-spacer"></div>
                <Link to='/home'><i className="homeIcon material-icons">restaurant</i></Link>
                <div className="mdl-layout-spacer"></div>
                <div className="footRight mdl-grid mdl-mini-footer__right-section">
                    <button className="mdl-cell--stretch mdl-button mdl-js-button mdl-button--accent">
                        <span>{this.state.right}</span>
                    </button>
                </div>
            </footer>
        );
    }



render() {
        switch (this.state.title) {
            case "devour":
                return (<span></span>);
                break;
            case "taste profile":
                return this.surveyFooter();
                break;
            case "recipes":
                return (<span></span>);
                break; 
            case "social":
                return this.splitFooter();
                break;
            case "recipe book":
                return (<span></span>)
                break;
            case "budget":
                return this.splitFooter()
                break;
            case "plan":
                return this.splitFooter();
                break;
        }
    }
}