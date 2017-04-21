import React from "react";

export default class DietTile extends React.Component {
    constructor(props) {
        super(props)
    }

    clickAction(e) {
        e.preventDefault()
        var target = e.currentTarget;
        target.style.backgroundColor = '#E0E0E0';
        // TODO: add redux to save all selections for diets
    }

    render() {
        return (
            <div className="mdl-card mdl-shadow--2dp mdl-cell--4-col dietCard" onClick={(e) => this.clickAction(e)} >
                <div className="mdl-card__title dietName">{this.props.title}</div>
            </div>
        );
    }
}