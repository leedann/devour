import React from "react";
import { Link } from 'react-router-dom';

export default class Footer extends React.Component {
    constructor(props) {
        super(props);
        this.state={
            left: this.props.page[0],
            right: this.props.page[1],
            title: this.props.title
        }
    }

    componentDidMount() {

    }

    surveyFooter() {
        return (
            <footer className="devourFooter mdl-mini-footer mdl-layout mdl-js-layout mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                <div className="footLeft mdl-grid mdl-mini-footer__left-section">
                    <i className="paginate material-icons">chevron_left</i>
                </div>
                <div className="mdl-layout-spacer"></div>
                <Link to='/home'><i className="homeIcon material-icons">restaurant</i></Link>
                <div className="mdl-layout-spacer"></div>
                <div className="footRight mdl-grid mdl-mini-footer__right-section">
                    <i className="paginate material-icons">chevron_right</i>
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