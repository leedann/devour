import React from "react";
import { Link } from 'react-router-dom';

//the buttons for the header
const buttons= {
    "/home": ["/shop", "/settings", "/info"],
    "/recipes": ["/search", "/filter"],
    "/social": ["/create", "/info"],  
    "/create": ["/create", "/info"],  
    "/info": ["/create", "/info"],  
}
//the header of every page, the buttons vary from page to page
export default class Header extends React.Component {
    constructor(props) {
        super(props);
        //gets the first portion of the page -- matches the first part of the path
        var strings = location.pathname.split("/", 2)
        this.state={
            icon1: this.props.page[0],
            icon2: this.props.page[1],
            icon3: this.props.page[2],
            title: this.props.title, 
            choices: buttons["/" + strings[1]]
        }
    }
    componentDidMount() {
    }

    //the header for the survey pages
    surveyHeader() {
        return (
                <header className="layoutHeader mdl-color--white mdl-layout__header">
                    <div className="headerRow mdl-layout__header-row">
                        <span className="headerTitle textSecondary mdl-layout-title">{this.state.title}</span>
                    </div>
                </header>
        );
    }

    //the normal headers with buttons
    normalHeader() {
        return (
            <header className="layoutHeader mdl-color--white mdl-layout__header">
                <div className="headerRow mdl-layout__header-row">
                    <Link to="/home"><i className="homeButton material-icons">restaurant</i></Link>
                    <span className="headerTitle textSecondary mdl-layout-title">{this.state.title}</span>
                    <nav className="mdl-navigation">
                        <Link to={this.state.choices[0]}> <i className="headerico material-icons">{this.state.icon1}</i></Link>
                        <Link to={this.state.choices[1]}> <i className="headerico material-icons">{this.state.icon2}</i></Link>
                    {this.state.choices[2]? <Link to={this.state.choices[2]}> <i className="headerico material-icons">{this.state.icon3}</i></Link> : ""}
                    </nav>
                </div>
            </header>
        );
    }



    render() {
        if (this.state.title !== "taste profile") {
            return (this.normalHeader());
        }else {
            return(this.surveyHeader());
        }
    }
}