import React from "react"
import Layout from "./Layout.jsx";
import {store} from "./shared-state.js";
const baseurl = 'https://dvrapi.leedann.me/v1/'
const myHeader = new Headers();
myHeader.append('Authorization', localStorage.getItem("Authorization"));
export default class FriendsList extends React.Component {
    constructor(props) {
        super(props);
        //not authorized
        if (!localStorage.getItem("Authorization")) {
            this.props.history.push('/')
        }
        this.state= store.getState()
    }

    componentDidMount() {
        this.unsub = store.subscribe(() => this.setState(store.getState()));
        var req = {
            method: 'GET',
            headers: myHeader,
        }
        fetch(baseurl+'users', req)
        .then((resp) => resp.json())
        .then(data => {
            this.setState({
                source: data
            })
        })
    }
    componentWillUnmount() {
        this.unsub();
    }
    render() {
        return(
            <Layout title="social">
                
            </Layout>
        );
    }
}

