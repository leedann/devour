import React from "react";
import { Link } from 'react-router-dom';
import {history} from './app.jsx'

//The login tile
export default class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            errmess: "",
        };
    }

    componentDidMount() {

    }


    handleSubmit(event) {
        event.preventDefault();
        let email=document.getElementById("email").value
        let password=document.getElementById("userpass").value
        var userinfo = {
            "Email": email,
            "Password": password
        }
        //users api call here
        if (userinfo) {

        }
        history.push('/home')
    }


    loginTile() {
        return (
            <div className="mdl-card mdl-shadow--2dp loginCard">
                <div className="mdl-card__title mdl-color--pink-A400 mdl-color-text--white">
                    <h2 className="mdl-card__title-text">Login</h2>
                </div>
                <div className="mdl-card__supporting-text">
                    {this.state.errmess? <span className="errmess" id="errmess">{this.state.errmess}</span> : ""}
                    <form action="#">
                        <div className="mdl-textfield mdl-js-textfield">
                            <input className="mdl-textfield__input" type="text" placeholder="Email" id="email"/>
                        </div>
                        <div className="mdl-textfield mdl-js-textfield">
                            <input className="mdl-textfield__input" type="password" placeholder="Password" id="userpass" />
                        </div>
                        <div className="mdl-grid mdl-cell mdl-cell--12-col lrContainer">
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone login">
                                <button onClick={(e) => this.handleSubmit(e)}className="mdl-button textAccent mdl-js-button mdl-js-ripple-effect">Log in</button>
                            </div>
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone register">
                                <Link to="/register"> <button className="mdl-button textAccent mdl-js-button mdl-js-ripple-effect">Register</button></Link>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
        );
    }

render() {
    return (
        <div className="mdl grid mdl-layout mdl-js-layout mdl-color--grey-200 startPage">
            <main className="mdl-grid mdl-layout__content mdl-cell mdl-cell--12-col mainContent">
                <div className="titleTag mdl-cell mdl-cell--12-col">
                    <h1 className="devourName mdl-color-text--amber-400">DEVOUR</h1>
                    <h4 className="devourTag textSecondary textAccent">Eat together. Your way.</h4>
                </div>
                {this.loginTile()}
            </main>
        </div>
        );
    }
}
