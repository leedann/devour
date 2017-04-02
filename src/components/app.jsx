import React from "react";

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    componentDidMount() {
    }

    render() {
        return (
            <main className="mdl-layout__content mdl-color-text--blue-grey-600">
                <h2>DEVOUR</h2>
                <h3>eat together.your way.</h3>
            

                <div className="mdl-grid mdl-card mdl-shadow--2dp startCard">
                    {/*<form action="submit">
                        <div className="mdl-textfield__input" type="text" id="username">
                            <label className="mdl-textfield__label" for="username">Username: </label>
                        </div>
                        <div className="mdl-textfield__input" type="password" id="password">
                        </div>
                        <button class="mdl-button mdl-js-button mdl-button--primary">
                        </button>
                    </form>*/}
                </div>
            </main>
        );
    }
}
