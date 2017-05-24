import React from "react";
import Layout from "./Layout.jsx";
const baseurl='https://api.yummly.com/v1/api/';
const apiHeader = new Headers();
apiHeader.append("X-Yummly-App-ID", "f3a4ad3c");
apiHeader.append("X-Yummly-App-Key", "8377c0df84a59fd60a40c2cb0ce4a964");
import LoaderExampleLoader from './loading.jsx';
//the cards for the recipes
class RecipeCard extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            fullDetails: null
        }
    }

    componentWillMount() {
        var data = {
            method: 'GET',
            headers: apiHeader
        }
        fetch(baseurl + `recipe/${this.props.details.id}`, data)
        .then((resp) => resp.json())
        .then(data => {
            console.log(data)
            this.setState({
                fullDetails: data
            })
        })
    }

    //toggles the info table 
    showInfo() {
        var tab = document.getElementById(this.props.details.id + "info")
        tab.classList.toggle("hidden")
    }

    //the snackbar confirmation for adding a recipe -- this doesnt work until you refresh the Recipe
    //will not work when you redir from a page
    addRecipe() {
        var snackbarContainer = document.getElementById("snackBar");
        //add a call to the yummly api with the params
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

    handleInfo(id, type) {
        var target = document.getElementById(id + type);
        var content = document.getElementById(id + type+"content");
        switch (type) {
            case "nutrition":
                var other = document.getElementById(id + "ingredients");
                var otherContent = document.getElementById(id + "ingredientscontent");
                var last = document.getElementById(id + "detail");
                var lastContent = document.getElementById(id + "detailcontent");
                target.classList.add('active');
                other.classList.remove('active');
                last.classList.remove('active')
                lastContent.classList.add('hidden');
                otherContent.classList.add('hidden');
                content.classList.remove('hidden');
                break;
            case "ingredients":
                other = document.getElementById(id + "nutrition");
                otherContent = document.getElementById(id + "nutritioncontent");
                last = document.getElementById(id + "detail");
                lastContent = document.getElementById(id + "detailcontent");
                target.classList.add('active');
                other.classList.remove('active');
                last.classList.remove('active');
                otherContent.classList.add('hidden');
                lastContent.classList.add('hidden');
                content.classList.remove('hidden');
                break;
            case "detail":
                other = document.getElementById(id + "nutrition");
                otherContent = document.getElementById(id + "nutritioncontent");
                last = document.getElementById(id + "ingredients");
                lastContent = document.getElementById(id + "ingredientscontent");
                target.classList.add('active');
                other.classList.remove('active');
                last.classList.remove('active');
                otherContent.classList.add('hidden');
                lastContent.classList.add('hidden');
                content.classList.remove('hidden');
                break;
            default:
                break;
        }
    }

    //currently splits the ingredients and desc but will have to work with yummly
    render() {
        if (!this.state.fullDetails) {
            return (<LoaderExampleLoader />);
        }
        var styles = {
            backgroundImage: `url(${this.state.fullDetails.images[0].hostedLargeUrl})`,
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
                    <i className="material-icons">query_builder</i>
                    <span className="recipeName mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">
                    {this.state.fullDetails.totalTime}</span>

                    <span className="recipeTime mdl-cell mdl-cell--6-col mdl-cell--2-col-phone mdl-cell--4-col-tablet">{this.props.details.CookTime}</span>
                </div>
                
                <div className="hidden recipeInfo" id={this.props.details.id + "info"}>
                    <div className="ui secondary menu">
                        <a id={this.props.details.id + "nutrition"}className="active item" onClick={() => this.handleInfo(this.props.details.id, "nutrition")}>Nutrition</a>
                        <a id={this.props.details.id + "ingredients"}className="item" onClick={() => this.handleInfo(this.props.details.id, "ingredients")}>Ingredients</a>
                        <a id={this.props.details.id + "detail"}className="item" onClick={() => this.handleInfo(this.props.details.id, "detail")}>Details</a>
                    </div>
                    <div className="ui bottom attached segment">
                        <div id={this.props.details.id + "ingredientscontent"} className="hidden">
                            <ul className="recipeList">
                                {this.state.fullDetails.ingredientLines.map((d, i) =>
                                <li key={this.props.details.id + "ing"+i}>{i+1+`. `+d}</li>
                                
                                )}
                            </ul>
                        </div>
                        <div id={this.props.details.id + "nutritioncontent"} >
                            <ul className="recipeList">
                                {this.state.fullDetails.nutritionEstimates.map((d, i) =>
                                <li key={this.props.details.id + "nut"+i}>{
                                    d.value + d.unit.abbreviation +' '+ (d.description? d.description : 'Fat')
                                    }</li>
                                
                                )}
                            </ul>
                        </div>
                        <div id={this.props.details.id + "detailcontent"} >
                            <ul className="recipeList">
                                <p>For detailed instructions visit: </p>
                                <a href={this.state.fullDetails.source.sourceRecipeUrl} target="_blank">{this.state.fullDetails.source.sourceDisplayName}</a>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

// the page of the recipes
export default class RecipePage extends React.Component {
    constructor(props) {
        super(props)
        this.state={
        }
    }

    componentDidMount() {
        var req = {
            method: 'GET',
            headers: apiHeader
        }

        fetch('http://api.yummly.com/v1/api/recipes?&maxResult=2&start=0&requirePictures=true', req)
        .then((resp) => resp.json())
        .then(data => {
            this.setState({
                matches: data.matches
            })
        })
    }

    render() {
        //this is the test array-- will change to the prop with api call to recipes
        return(
            <Layout title="recipes">
                <p className="titleName">View Recipes</p>
                <div className="recipeWrapper mdl-grid">
                    {this.state.matches? this.state.matches.map((rec) => {
                        return <RecipeCard key={rec.id} details={rec}/>
                    }) : ""}
                </div>
            </Layout>
        );
    }
}