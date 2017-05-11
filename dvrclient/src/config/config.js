import * as firebase from "firebase"

const CONFIG = {
  apiKey: "AIzaSyALrZU2ERCUu_wAfJEZtfb6s1eXgqZm7Hk",
  authDomain: "devour-bf0f3.firebaseapp.com",
  databaseURL: "https://devour-bf0f3.firebaseio.com",
  projectId: "devour-bf0f3",
  storageBucket: "devour-bf0f3.appspot.com",
  messagingSenderId: "1014896556090"
};

firebase.initializeApp(CONFIG);

export const ref = firebase.database().ref()
export const firebaseAuth = firebase.auth