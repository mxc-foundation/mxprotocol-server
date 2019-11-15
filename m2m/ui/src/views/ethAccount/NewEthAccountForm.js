import React from "react";

import TextField from '@material-ui/core/TextField';
import i18n, { packageNS } from '../../i18n';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
class NewEthAccountForm extends FormComponent {

  state = {
    createAccount: '',
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

  onSubmit = () => {
    this.props.onSubmit({
      action: 'createAccount',  
      createAccount: this.state.createAccount,
      currentAccount: this.state.createAccount,
      username: this.state.username,
      password: this.state.password
    });

    this.clear();
  }

  render() {
    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={this.onSubmit}
      >
        <TextField
          id="createAccount"//it is defined current account in swagger
          label={i18n.t(`${packageNS}:menu.withdraw.new_account`)}
          margin="normal"
          value={this.state.createAccount}
          placeholder="0x0000000000000000000000000000000000000000" 
          onChange={this.onChange}
          inputProps={{
            pattern: "^0x[a-fA-F0-9]{40}$",
          }}

          autoComplete='off'
          required
          fullWidth
        />

        <TextField
          id="username"//it is defined current account in swagger
          label={i18n.t(`${packageNS}:menu.withdraw.username`)}
          margin="normal"
          value={this.state.username}
          placeholder={i18n.t(`${packageNS}:menu.withdraw.type_here`)}
          onChange={this.onChange}
          autoComplete='off'
          required
          fullWidth
        />

        <TextField
          id="password"//it is defined current account in swagger
          label={i18n.t(`${packageNS}:menu.withdraw.password`)}
          margin="normal"
          value={this.state.password}
          placeholder={i18n.t(`${packageNS}:menu.withdraw.type_here`)}
          onChange={this.onChange}
          type="password"
          autoComplete="off"
          required
          fullWidth
        />
       
      </Form>
    );
  }
}

export default NewEthAccountForm;
