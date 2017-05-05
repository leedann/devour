import React from "react";
import Header from "./Header.jsx";
import Footer from "./Footer.jsx";

const homeIcons=["shopping_cart", "settings"];
const eventFooter=["Upcoming", "Past"];
const recipeIcons=["search", "filter"];
const planFooter=["Day", "Week"];
const socialIcons=["library_add", "info"];
const budgetFooter=["Dashboard", "Spending"];
const base=["", "info"];
const blank=["",""]

export default class Layout extends React.Component {
    constructor(props) {
        super(props);
        this.state={
            title: this.props.title,
            page: "",
            footNames: ""
        }
    }

    componentDidMount() {
    }

    iconHelper() {
        switch (this.state.title) {
            case "devour":
                return homeIcons;
                break;
            case "taste profile":
                return blank;
                break;
            case "recipes":
                return recipeIcons;
                break; 
            case "social":
                return socialIcons;
                break;
            case "recipe book":
                return base;
                break;
            case "budget":
                return base;
                break;
            case "plan":
                return base;
                break;
        }
    }

    footerHelper() {
        switch (this.state.title) {
            case "devour":
                return blank;
                break;
            case "taste profile":
                return blank
                break;
            case "recipes":
                return blank
                break; 
            case "social":
                return eventFooter;
                break;
            case "recipe book":
                return blank;
                break;
            case "budget":
                return budgetFooter;
                break;
            case "plan":
                return planFooter;
                break;
        }
    }


    render() {
        var headerIcons = this.iconHelper();
        var footieWords = this.footerHelper();
        return (
            <div className="pageWrap mdl-layout mdl-js-layout">
                <Header title={this.state.title} page={headerIcons} />
                { this.props.children }
                <Footer title={this.state.title} page={footieWords}/>
            </div>
        );
    }
}


