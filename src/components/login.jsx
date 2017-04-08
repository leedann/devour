import React from "react";

export default class LoginTile extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            tile: this.login()
        };
    }

    componentDidMount() {
        //login is true--need to change to registration
        if (this.state.login) {
            this.setState({
                login: false,
                tile: this.register()
             })
        }else {
            this.setState({
                login: true,
                tile: this.login()
             })
        }
    }

    login() {
        return (
        <div className="mdl-card mdl-shadow--2dp loginCard">
                        <div className="mdl-card__title mdl-color--amber-400 mdl-color-text--white">
                            <h2 className="mdl-card__title-text">Login</h2>
                        </div>
                        <div className="mdl-card__supporting-text">
                            <form action="#">
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="text" id="username" />
                                    <label className="mdl-textfield__label" for="username">Username...</label>
                                </div>
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="password" id="userpass" />
                                    <label className="mdl-textfield__label" for="userpass">Password...</label>
                                </div>
                            </form>
                        </div>
                        <div className="mdl-grid mdl-cell mdl-cell--12-col lrContainer">
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone login">
                                <button className="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Log in</button>
                            </div>
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone register">
                                <button onClick={() => this.componentDidMount()}className="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Register</button>
                            </div>
                        </div>
                    </div>
        );
    }
    register() {
        return (
        <div className="mdl-card mdl-shadow--2dp registerCard">
                        <div className="mdl-card__title mdl-color--amber-400 mdl-color-text--white">
                            <h2 className="mdl-card__title-text">Registration</h2>
                        </div>
                        <div className="mdl-card__supporting-text">
                            <form action="#">
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="text" id="firstName" required autoFocus/>
                                    <label className="mdl-textfield__label" for="firstName">First Name..</label>
                                </div>
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="text" id="lastName" required/>
                                    <label className="mdl-textfield__label" for="lastName">Last Name..</label>
                                </div>
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="password" id="pass" placeholder="Password"required/>
                                </div>
                                <div className="mdl-textfield mdl-js-textfield">
                                    <input className="mdl-textfield__input" type="password" id="pass2" placeholder="Re-Enter Password"required/>
                                </div>
                            </form>
                        </div>
                        <div className="mdl-grid mdl-cell mdl-cell--12-col lrContainer">
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone login">
                                <button className="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Log in</button>
                            </div>
                            <div className="mdl-cell mdl-cell--6-col mdl-cell--2-col-phone register">
                                <button className="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Register</button>
                            </div>
                        </div>
                    </div>
        );
    }


    render() {
        return (this.state.tile);
    }
}