import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";

class ModifyEthAccountForm extends FormComponent {

  state = {
    newaccount: '',
    username: '',
    password: ''
  }
  
  onChange = (event) => {
    const { id, value } = event.target;
    
    this.setState({
      [id]: value
    });
  }

  clear() {
    this.setState({
      username: '',
      password: '',
      newaccount: ''
    })
  }

  submit = () => {
    this.props.onSubmit({
      action: 'modifyAccount',  
      currentAccount: this.state.newaccount,
      createAccount: this.state.newaccount,
      username: this.state.username,
      password: this.state.password
    })

    this.setState({
      username: '',
      password: '',
      newaccount: ''
    })
  }

  render() {
    if (this.props.activeAccount === undefined) {
      return(<div></div>);
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={this.submit}
      >
        <TextField
          id="activeAccount"
          label="Current account"
          margin="normal"
          value={this.props.activeAccount || ""}

          InputProps={{
            readOnly: true,
          }}
          fullWidth
        />

        <TextField
          id="newaccount"//it is defined current account in swagger
          label="New account"
          margin="normal"
          value={this.state.newaccount}
          placeholder="0x0000000000000000000000000000000000000000" 
          onChange={this.onChange}
          /* inputProps={{
            pattern: "^0x[a-fA-F0-9]{40}$",
          }} */
            
          autoComplete='off'
          required
          fullWidth
        />

        <TextField
          id="username"//it is defined current account in swagger
          label="User name"
          margin="normal"
          value={this.state.username}
          placeholder="Type here" 
          onChange={this.onChange}
          inputProps={{
            pattern: "[\\w-]+",
          }}
          autoComplete='off'
          required
          fullWidth
        />

        <TextField
          id="password"//it is defined current account in swagger
          label="Password"
          margin="normal"
          value={this.state.password}
          placeholder="Type here" 
          onChange={this.onChange}
          inputProps={{
            pattern: "[\\w-]+",
          }}
          type="password"
          autoComplete="off"
          required
          fullWidth
        />
        
        {/* <TitleBarButton
            key={1}
            label="Go to Etherscan.io"
            icon={<LinkI />}
            color="secondary"
            onClick={this.deleteOrganization}
        /> */}
          
      </Form>
    );
  }
}

export default ModifyEthAccountForm;
