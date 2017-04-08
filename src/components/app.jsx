import React from "react";
import LoginTile from "./login.jsx";

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    componentDidMount() {
    }

    render() {
        return (
            <div className="mdl grid mdl-layout mdl-js-layout mdl-color--grey-200 startPage">
                <main className="mdl-grid mdl-layout__content mdl-cell mdl-cell--12-col mainContent">
                    <div className="titleTag mdl-cell mdl-cell--12-col">
                        <h2 className="titleName mdl-color-text--amber-400">DEVOUR</h2>
                        <h4 className="titleName mdl-color-text--primary">Eat together. Your way.</h4>
                    </div>
                    <LoginTile></LoginTile>
                </main>
            </div>

        );
    }
}
