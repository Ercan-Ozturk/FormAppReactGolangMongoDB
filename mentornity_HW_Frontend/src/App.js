import logo from './logo.svg';
import React, { Component } from 'react';
import './App.css';
import PostForm from "./PostForm";



class App extends Component {

    render(){
        return(
            <div className="App">
                <PostForm/>
            </div>
        )
    }


}

export default App;
