import React from "react";
import { Link } from 'react-router-dom';
import {history} from './app.jsx';
import DOBForm from './helpers/DOBdown.jsx';
import MonthDown from './helpers/dropdown.jsx';
import moment from "moment"; 
const baseurl = 'https://dvrapi.leedann.me/v1/'

const datefmt = {
    "January": "01",
    "February": "02",
    "March": "03",
    "April": "04",
    "May": "05",
    "June": "06",
    "July": "07",
    "August": "08",
    "September": "09",
    "October": "10",
    "November": "11",
    "December": "12"
}


//Registration page
export default class Register extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            passMatch: "",
            passLength: "",
            validDate: "",
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
    passLength(e) {
        e.preventDefault();
        let pass1 = document.getElementById("pass");
        let pass2 = document.getElementById("pass2");
        if (pass1.value < 6 || pass2.value < 6) {
            this.setState({
                passLength: false
            })
        }else {
            this.setState({
                passLength: true
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
        let month = document.getElementById("dobMonth").getElementsByClassName('text')[0].textContent;
        let day = document.getElementById("dobDay").value;
        let year = document.getElementById("dobYear").value;
        if (parseInt(day, 10) < 10) {
            day = "0" + parseInt(day, 10);
        }
        var date = `${datefmt[month]}/${day}/${year}`
        var momentCheck = `${year}-${datefmt[month]}-${day}`

        //if passwords are matching, POST to the url
        if (!this.state.passMatch && !this.state.passLength && moment(momentCheck).isValid()) {
            var data = {
                "Email": email,
                "Password": password,
                "PasswordConf": passconf,
                "FirstName": firstName,
                "LastName": lastName,
                "DOB": date
            }
            //api post call to the users db
            if (data) {
                data = JSON.stringify(data);
                var req = {
                    method: 'POST',
                    body: data
                }
                fetch(baseurl+'users', req)
                .then((resp) => {
                    localStorage.setItem("Authorization", resp.headers.get('Authorization'));
                    history.push('/welcome')
                })
                .catch((error) => {
                    this.setState({
                        errmess: error.response.data
                    })
                })
            }
        }else if (!moment(momentCheck).isValid()) {
            this.setState({
                validDate: true
            })
        }
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
                        <div>
                            {this.state.validDate? <span className="notValid" id="notValid">Please pick a valid date</span> : ""}
                            <p>Date of Birth:</p>
                            <MonthDown/>
                            <DOBForm/>
                        </div>
                        {this.state.passMatch? <span className="notMatch" id="notMatch">Passwords do not match</span> : ""}
                        <br/>
                        {this.state.passLength? <span className="notLength" id="notLength">Passwords must be at least 6 characters</span> : ""}
                        <div className="mdl-textfield mdl-js-textfield">
                            <input className="mdl-textfield__input" onChange={(e) => {
                                this.passLength(e)
                                }} type="password" id="pass" placeholder="Password"required/>
                        </div>
                        <div className="mdl-textfield mdl-js-textfield">
                            <input className="mdl-textfield__input" onChange={(e) => {
                                this.passwordMatch(e)
                                } }type="password" id="pass2" placeholder="Re-Enter Password"required/>
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