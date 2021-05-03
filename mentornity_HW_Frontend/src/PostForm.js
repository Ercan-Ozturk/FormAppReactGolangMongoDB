import React, {Component} from "react";

import "./style.css";

class PostForm extends Component{
    constructor(props) {
        super(props);

        this.state = {
            Name: '',
            Email: '',
            Message: '',
            nameError:'',
            emailError:'',
            messageError:''
        }
    }
    validate = () => {
        let nameError = ""
        let emailError = ""
        let messageError = ""
        if(!this.state.Name){
            nameError = 'Empty Name'
        }
        if(!this.state.Message){
            messageError = 'Empty Message'
        }

        if(!this.state.Email.includes('@')){
            emailError = 'Invalid Email'
        }
        if(emailError || nameError || messageError){
            this.setState({emailError, nameError, messageError})
            return false
        }
        return true

    }
    changeHandler = e =>{
        this.setState({[e.target.name]: e.target.value})
    }
    submitHandler = e =>{

        e.preventDefault()

        const isValid = this.validate()

        if(isValid){
            console.log(this.state)

            const postURL = "http://localhost:8080/api/items" //Our previously set up route in the backend
            //const postURL =  "https://webhook.site/2783210c-5025-41dd-b8e8-74b43c4aa044"
            fetch(postURL, {
                method: 'POST',
                mode: 'no-cors',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                    //'Content-Type': 'X-www-form-urlencoded',
                },
                body: JSON.stringify({ // We should keep the fields consistent for managing this data later
                    Name: this.state.Name,
                    Email:this.state.Email,
                    Message:this.state.Message
                })
            })
                .then(()=>{
                    // Once posted, the user will be notified
                    alert('You have been added to the system!');
                })

        }
        else{
            alert("Invalid input!")
        }






    }


    render() {
        const {Name, Email, Message} = this.state
        return(
            <div>
                <div className="header">
                    <h1>Mentornity Form App</h1>
                </div>
                <form onSubmit={this.submitHandler}>
                    <div className="field">
                        <input placeholder="Name" type="text" name = "Name" value={Name} onChange={this.changeHandler}/>
                    </div>
                    <div className="field">
                        <input placeholder="Email" type="text" name = "Email" value={Email} onChange={this.changeHandler}/>
                    </div>
                    <div className="field">
                        <input placeholder="Message" type="text" name = "Message" value={Message} onChange={this.changeHandler}/>
                    </div>
                    <div className="field button">
                        <button type= "submit"> Submit</button>
                    </div>

                </form>
            </div>


        )
    }
}

export  default  PostForm