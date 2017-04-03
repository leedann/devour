import React from "react";

// const bIma = "url(/img/frontpage.jpg)";
export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    componentDidMount() {
    }

    render() {
        // var img = {
        //     backgroundImage: bIma,

        // };
        return (
            <div className="mdl grid mdl-layout mdl-js-layout mdl-color--grey-200 startPage">
                <main className="mdl-grid mdl-layout__content mdl-cell mdl-cell--12-col mainContent">
                    <div className="titleTag mdl-cell mdl-cell--12-col">
                        <h2 className="titleName mdl-color-text--amber-400">DEVOUR</h2>
                        <h4 className="titleName mdl-color-text--primary">Eat together. Your way.</h4>
                    </div>

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
                                <button className="mdl-button mdl-button--colored mdl-js-button mdl-js-ripple-effect">Register</button>
                            </div>
                        </div>
                    </div>
                </main>
            </div>

        );
    }
}
