import React from "react";
import Diet_Tile from "./diet_tiles.jsx"
import Layout from "./Layout.jsx";

export default class SurveyGoals extends React.Component {
    constructor(props) {
        super(props)
        this.state = {};
    }

    render() {
        //test array
        var goals = ["All of these Reasons", "Plan group meals", "Learn to cook", "Eat Healthy", "Find Recipes", "Plan Meals", "Entertain", "Stick to a Diet", "Budget"]
        return(
            <Layout title="taste profile">
                <div className="surveywrapper mdl-layout mdl-js-layout">
                    <header className="mdl-layout__header mdl-layout__header-row center_title mdl-layout__header--transparent mdl-color--white mdl-color-text--yellow-500 ">select yo goals</header>
                    <div className="mdl-grid dietGrid">
                        {goals.map((d) => <Diet_Tile title={d}/>)}
                    </div>
                </div>
            </Layout>
        );
    }
}