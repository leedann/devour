import React from "react";
import { Link } from 'react-router-dom';
import {history} from './app.jsx'

//Registration page
export default class Register extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            passMatch: "",
            errmess: ""
        };
    }

    componentDidMount() {

    }

    //if the password does not match dont let register
    passwordMatch(e) {
        e.preventDefault()
        let pass1 = document.getElementById("pass");
        let pass2 = document.getElementById("pass2");
        if (pass1.value === pass2.value) {
            this.setState({
                passMatch: false
            })
        }else {
            this.setState({
                passMatch: true
            })
        }
    }

    //handles submitting the registration
    handleSubmit(event) {
        event.preventDefault();
        let firstName=document.getElementById("firstName").value
        let lastName=document.getElementById("lastName").value
        let email=document.getElementById("email").value
        let password=document.getElementById("pass").value
        let passconf=document.getElementById("pass2").value
        let username = "";
        //if passwords are matching, POST to the url
        if (!this.state.passMatch) {
            var data = {
                "Email": email,
                "Password": password,
                "PasswordConf": passconf,
                "Username": username,
                "FirstName": firstName,
                "LastName": lastName
            }
            //api post call to the users db
            if (data) {

            }
        }
        //will need to move to end of successful api call
        history.push("/welcome");
    }

    registerTile() {
        return (
            <div className="mdl-card mdl-shadow--2dp registerCard">
                <div className="mdl-card__title mdl-color--pink-A400 mdl-color-text--white">
                    <h2 className="mdl-card__title-text">Registration</h2>
                </div>
                <div className="mdl-card__supporting-text">
                    {this.state.errmess? <span className="errmess" id="errmess">{this.state.errmess}</span> : ""}
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
                        {this.state.passMatch? <span className="notMatch" id="notMatch">Passwords do not match</span> : ""}
                        <div className="mdl-textfield mdl-js-textfield">
                            <input className="mdl-textfield__input" type="password" id="pass" placeholder="Password"required/>
                        </div>
                        <div className="mdl-textfield mdl-js-textfield">
                            <input className="mdl-textfield__input" onChange={(e) => this.passwordMatch(e)}type="password" id="pass2" placeholder="Re-Enter Password"required/>
                        </div>
                        <div className="mdl-grid mdl-cell mdl-cell--12-col lrContainer">
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone login">
                                <Link to="/"><span className="backToLogin textAccent">Already have an account? </span></Link>
                            </div>
                            <div className="mdl-cell mdl-cell--1-col mdl-cell--1-col-phone register">
                                <button className="mdl-button textAccent mdl-js-button mdl-js-ripple-effect" onClick={(e) => this.handleSubmit(e)}>Register</button>
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
                    {this.registerTile()}
                </main>
            </div>

        );
    }
}