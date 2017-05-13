import React from "react";
import DietTile from "./diet_tiles.jsx"
import Layout from "./Layout.jsx";

export default class Survey extends React.Component {
    constructor(props) {
        super(props)
        this.state = {};
    }

    render() {
        //test array
        var diets = ["Everything", "Vegetarian", "Vegan", "Low Carb", "Pescetarian", "Paleo", "Gluten-Free", "Flexitarian", "Pollotarian"]
        return(
            <Layout title="taste profile">
                <header className="mdl-layout__header mdl-layout__header-row center_title mdl-layout__header--transparent mdl-color--white mdl-color-text--yellow-500 ">choose your menu</header>
                <div className="mdl-grid dietGrid">
                    {diets.map((d) => <DietTile title={d}/>)}
                </div>
            </Layout>
        );
    }
}