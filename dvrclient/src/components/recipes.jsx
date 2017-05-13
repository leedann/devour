import React from "react"
import Layout from "./Layout.jsx";


var rec1={
    ID: "1",
    Name: "Eggs Benedict",
    Ingredients: "1 egg;1/2 english muffin;1 slice of ham;3 Tbsp hollandaise sauce;paprika;salt and pepper",
    Description: "1. Poach Egg.;2. Lightly toast English muffin half.;3. Top with ham slice, (tucking square corners under to conform to round muffin).;4. Top ham with poached egg.;5. Cover all with 3 tablespoons of Hollandaise Sauce (See recipe #25) for (Quickly Hollandaise Sauce).",
    PhotoURL: "https://cdn.mtlblog.com/uploads/144916_4a0a45916e74f824a523f85a3285a1177d663b18.jpg",
    CookTime: "20m"
}
var rec2={
    ID: "2",
    Name: "Buffalo Chicken Mac and Cheese",
    Ingredients: "1 (16oz) package elbow macaroni;1 rotisserie-roasted chicken;6 Tbsp Butter;6 Tbsp all-purpose flour;1 pinch ground black pepper;2 cups shredded cheddar cheese;2 cups shredded monterey jack cheese;1/2 cup hot sauce",
    Description: "1. Bring a large pot of lightly salted water to a boil. Cook macaroni in the boiling water, stirring occasionally until tender yet firm to the bite, 8 minutes. Drain.;2. Cut wings and legs off rotisserie chicken. Skin and bone wings and legs; chop or shred dark meat into bite-size pieces.;3. Melt butter in a large Dutch oven over medium heat. Whisk in flour gradually until a thick paste forms. Cook until golden, about 1 minute. Pour in milk, whisking constantly, until thickened and bubbling, about 5 minutes. Continue to cook until sauce is smooth, about 1 minute more. Reduce heat and season sauce with black pepper.;4. Stir Cheddar and Monterey Jack cheese into the sauce until melted and combined. Stir in hot sauce, adjusting to reach desired level of spiciness. Add blue cheese, chicken, and macaroni; mix well to combine.",
    PhotoURL: "http://thenerderypublic.com/wp-content/uploads/2014/01/5429708155_3e7e18ecf2_thumb.jpg",
    CookTime: "35m"
}
var rec3={
    ID: "3",
    Name: "Pork Chops with Fresh Tomato, Onion, Garlic, and Feta",
    Ingredients: "2 Tbsp olive oil;1 large onion, halved and thinly sliced;4 pork loin chops, 1 inch thick;salt to taste;black pepper to taste;garlic powder to taste;1/2 pint red grape tomatoes, halved; 1/2 pint yellow grape tomatoes, halved; 3 cloves garlic, diced;1 Tbsp dried basil; 2 1/2 tsp balsamic vinegar;4oz feta cheese, crumbled",
    Description: "1. Heat 1 tablespoon oil in a skillet over medium heat. Stir in the onion and cook until golden brown. Set aside.;2. Heat 1/2 tablespoon oil in the skillet. Season pork chops with salt, pepper, and garlic powder, and place in the skillet. Cook to desired doneness. Set aside and keep warm.;3. Heat remaining oil in the skillet. Return onions to skillet, and stir in tomatoes, garlic, and basil. Cook and stir about 3 minutes, until tomatoes are tender. Mix in balsamic vinegar, and season with salt and pepper. Top chops with the onion and tomato mixture, and sprinkle with feta cheese to serve.",
    PhotoURL: "http://chewnibblenosh.com/wp-content/uploads/2015/08/Pork-Chops-with-Fresh-Tomato-Onion-Garlic-and-Feta-Chew-Nibble-Nosh-.png",
    CookTime: "35m"
}


class RecipeCard extends React.Component {
    constructor(props) {
        super(props)
        this.state={

        }
    }

    showInfo() {
        var tab = document.getElementById(this.props.details.ID + "info")
        tab.classList.toggle("hidden")
    }

    addRecipe() {
        var snackbarContainer = document.getElementById("snackBar");
        var recipeInfo = this.props.details
        //make ajax with recipe info to add to user's book
        var handler = function(event) {

        };
        var data = {
        message: 'Added ' + this.props.details.Name,
        timeout: 2000,
        actionhandler: handler,
        actionText: 'Undo'
        };
        snackbarContainer.MaterialSnackbar.showSnackbar(data);
    }

    render() {
        var allIng = this.props.details.Ingredients.split(";")
        var allDesc = this.props.details.Description.split(";")
        var styles = {
            backgroundImage: `url(${this.props.details.PhotoURL})`,
            backgroundSize: 'cover'
        }
        return(
            <div className="recipeCard mdl-card mdl-shadow--2dp mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">
                <div className="imageContainer mdl-card__title mdl-card--expand" style={styles}>
                    <button className="favButton mdl-button mdl-js-button mdl-button--fab mdl-js-ripple-effect mdl-button--accent" onClick={() => this.addRecipe()}>
                        <i className="material-icons">favorite_border</i>
                    </button>
                </div>
                <div className="recipeActions mdl-card__actions mdl-card--border mdl-grid" onClick={() => this.showInfo()}>
                    <span className="recipeName mdl-cell mdl-cell--12-col mdl-cell--8-col-phone mdl-cell--8-col-tablet">{this.props.details.Name}</span>
                    <i className="material-icons">query_builder</i>
                    <span className="recipeTime mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">{this.props.details.CookTime}</span>
                </div>
                <div className="recipeInfo mdl-tabs mdl-js-tabs mdl-js-ripple-effect hidden" id={this.props.details.ID + "info"}>
                    <div className="mdl-tabs__tab-bar">
                        <a href={"#desc-panel" + this.props.details.ID} className="mdl-tabs__tab is-active">Description</a>
                        <a href={"#ing-panel" + this.props.details.ID} className="mdl-tabs__tab">Ingredients</a>
                    </div>
                    <div className="mdl-tabs__panel is-active" id={"desc-panel" + this.props.details.ID}>
                        <ul>
                            {allDesc.map((d, i) =>
                            <li key={this.props.details.ID + "desc"+i}>{d}</li>)}
                        </ul>
                    </div>
                    <div className="mdl-tabs__panel" id={"ing-panel" + this.props.details.ID}>
                        <ul>
                            {allIng.map((d, i) =>
                            <li key={this.props.details.ID + "ing"+i}>{d}</li>)}
                        </ul>
                    </div>
                </div>
            </div>
        );
    }
}

export default class RecipePage extends React.Component {
    constructor(props) {
        super(props)
        this.state={
        }
    }

    componentDidMount() {
    }

    render() {
        //this is the test array-- will change to the prop with api call to recipes
        var arr = [rec1, rec2, rec3]
        return(
            <Layout title="recipes">
                <p className="titleName">View Recipes</p>
                <div className="recipeWrapper mdl-grid">
                    {arr.map(rec => 
                    <RecipeCard key={rec.ID} details={rec}/>
                    )}
                </div>
            </Layout>
        );
    }
}