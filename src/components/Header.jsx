import React from "react";
import { Link } from 'react-router-dom';

const buttons= {
    "/home": ["shop", "settings"],
    "/recipes": ["search", "filter"],
    "/social": ["create", "info"],  
    "/create": ["create", "info"],  
    "/info": ["create", "info"],  
}
export default class Header extends React.Component {
    constructor(props) {
        super(props);
        this.state={
            icon1: this.props.page[0],
            icon2: this.props.page[1],
            title: this.props.title, 
            choices: buttons[location.pathname]
        }
    }
    componentDidMount() {
    }
    surveyHeader() {
        return (
            <div className="headerContent mdl-layout mdl-js-layout mdl-layout--fixed-header mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                <header className="layoutHeader mdl-color--white mdl-layout__header">
                    <div className="headerRow mdl-layout__header-row">
                        <div className="mdl-layout-spacer"></div>
                        <span className="textSecondary mdl-layout-title">{this.state.title}</span>
                        <div className="mdl-layout-spacer"></div>
                    </div>
                </header>
            </div>
        );
    }

    normalHeader() {
        return (
        <div className="headerContent mdl-layout mdl-js-layout mdl-layout--fixed-header mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
            <header className="layoutHeader mdl-color--white mdl-layout__header">
                <div className="headerRow mdl-layout__header-row">
                    <Link to="/home"><i className="homeButton material-icons">restaurant</i></Link>
                    <div className="mdl-layout-spacer"></div>
                    <span className="textSecondary mdl-layout-title">{this.state.title}</span>
                    <div className="mdl-layout-spacer"></div>
                    <nav className="mdl-navigation">
                        <Link to={this.state.choices[0]}> <i className="headerico material-icons">{this.state.icon1}</i></Link>
                        <Link to={this.state.choices[1]}> <i className="headerico material-icons">{this.state.icon2}</i></Link>
                    </nav>
                </div>
            </header>
        </div>
        );
    }



    render() {
        if (this.state.title != "taste profile") {
            return (this.normalHeader());
        }else {
            return(this.surveyHeader());
        }
    }
}