import React from "react";
import {login, resetPassword} from "../helpers/auth.js"
import {ref, firebaseAuth } from '../config/config.js'

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            errorMessage: "",
            login: true,
            passNoMatch: false
        };

        this.loginTile= this.loginTile.bind(this)
        this.registerTile= this.registerTile.bind(this)
    }

    componentDidMount() {

    }

    passwordMatch(e) {
        e.preventDefault()
        let pass1 = document.getElementById("pass");
        let pass2 = document.getElementById("pass2");
        if (pass1.value != pass2.value) {
            this.setState({
                errorMessage: "",
                login: false,
                passNoMatch: true
            })
        }else {
            this.setState({
                errorMessage: "",
                login: false,
                passNoMatch: false
            })
        }
    }

    handleTileChange(e) {
        e.preventDefault()
        if (this.state.login) {
            this.setState({
                errorMessage: "",
                login: false
            })
        }else {
            this.setState({
                errorMessage: "",
                login: true
            })
        }
    }

    handleLogin(e) {
        e.preventDefault();
        let email  = document.getElementById("email");
        let password = document.getElementById("userpass");
        firebaseAuth().signInWithEmailAndPassword(email.value, password.value)
            .then(function() {
                this.setState({
                    errorMessage: "",
                    login: true
                })
            })
            .catch((error) => {
                this.setState({
                    errorMessage: "Invalid email or password",
                    login: true
                });
            });
    }

    loginTile() {
        var errorMessage = this.state.errorMessage;
        return (
        <div className="mdl-card mdl-shadow--2dp loginCard">
                        <div className="mdl-card__title mdl-color--amber-400 mdl-color-text--white">
                            <h2 className="mdl-card__title-text">Login</h2>
                        </div>
                        <div className="mdl-card__supporting-text">
                            <ErrorMessage errorMessage={errorMessage? this.state.errorMessage : ""}/>
                            <form action="#">
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="email" id="email" ref={(email) => this.state.email = email}/>
                                    <label className="mdl-textfield__label" htmlFor="email">Email...</label>
                                </div>
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="password" id="userpass" ref={(pw) => this.state.pw = pw} />
                                    <label className="mdl-textfield__label" htmlFor="userpass">Password...</label>
                                </div>
                            </form>
                        </div>
                        <div className="mdl-grid mdl-cell mdl-cell--12-col lrContainer">
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone login">
                                <button onClick= {(e) => this.handleLogin(e)}className="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Log in</button>
                            </div>
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone register">
                                <button onClick={(e) => this.handleTileChange(e)}className="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Register</button>
                            </div>
                            <a href="#" className="forgot mdl-color-text--primary mdl-cell--12-col">Forgot your password?</a>
                        </div>
                    </div>
        );
    }
    registerTile() {
        var passMatch = this.state.passNoMatch;
        return (
        <div className="mdl-card mdl-shadow--2dp registerCard">
                        <div className="mdl-card__title mdl-color--amber-400 mdl-color-text--white">
                            <h2 className="mdl-card__title-text">Registration</h2>
                        </div>
                        <div className="mdl-card__supporting-text">
                            <form action="#">
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="text" id="firstName" placeholder="First Name" required autoFocus/>
                                </div>
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="text" id="lastName" placeholder="Last Name" required/>
                                </div>
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="email" id="email" placeholder="Email" required/>
                                </div>
                                {passMatch? <span className="notMatch" id="notMatch" hidden>Passwords do not match</span> : ""}
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="password" id="pass" placeholder="Password"required/>
                                </div>
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" onChange={(e) => this.passwordMatch(e)}type="password" id="pass2" placeholder="Re-Enter Password"required/>
                                </div>
                            </form>
                        </div>
                        <div className="mdl-grid mdl-cell mdl-cell--12-col lrContainer">
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone login">
                                <a href="#" className="backToLogin mdl-color-text--primary" onClick={(e) => this.handleTileChange(e)}>Already have an account? </a>
                            </div>
                            <div className="mdl-cell mdl-cell--1-col mdl-cell--1-col-phone register">
                                <button className="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Register</button>
                            </div>
                        </div>
                    </div>
        );
    }

    render() {
        return (
            <div className="mdl grid mdl-layout mdl-js-layout mdl-color--grey-200 startPage">
                <main className="mdl-grid mdl-layout__content mdl-cell mdl-cell--12-col mainContent">
                    <div className="titleTag mdl-cell mdl-cell--12-col">
                        <h2 className="titleName mdl-color-text--amber-400">DEVOUR</h2>
                        <h4 className="titleName mdl-color-text--primary">Eat together. Your way.</h4>
                    </div>
                    {this.state.login? this.loginTile() : this.registerTile()}
                </main>
            </div>

        );
    }
}

export class ErrorMessage extends React.Component {
    constructor(props) {
        super(props)
    }

    render() {
        if (this.props.errorMessage) {
            return (<div className="errorMessage">{this.props.errorMessage}</div>);
        }else {
            return <p> </p>
        }
    }
}