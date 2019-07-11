import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";

class ModifyEthAccountForm extends FormComponent {

  onChange = (event) => {
    const { id, value } = event.target;
    
    this.setState({
      [id]: value
    });
  }

  render() {
    if (this.props.activeAccount === undefined) {
      return(<div></div>);
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={() => this.props.onSubmit({
          currentAccount: this.state.newAccount,
          username: this.state.username,
          password: this.state.password
        })}
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
          id="newAccount"//it is defined current account in swagger
          label="New account"
          margin="normal"
          value={this.state.newaccount}
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
          label="Pass word"
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
