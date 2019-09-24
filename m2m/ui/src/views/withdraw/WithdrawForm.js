import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
//import Button from "@material-ui/core/Button";
import Spinner from "../../components/ScaleLoader"
import { withRouter } from "react-router-dom";

class WithdrawForm extends FormComponent {
  
  state = {
    amount: ''
  }
 
  onChange = (event) => {
    const { id, value } = event.target;
    
    this.setState({
      [id]: value
    });
  }

  clear() {
    this.setState({
      amount: ''
    })
  }

  render() {
    if (this.props.txinfo === undefined) {
      return(<Spinner on={this.state.loading}/>);
    }
    
    const wLimit = this.props.txinfo.balance - this.props.txinfo.withdrawFee;
    const { txinfo } = this.props;
    
    return(
      <Form
        submitLabel={this.props.submitLabel}
        //extraButtons={extraButtons}
        onSubmit={(e) => this.props.onSubmit(e, {
          amount: parseFloat(this.state.amount),
        })}
      >
        <TextField
          id="amount"
          label="Amount"
          margin="normal"
          value={this.state.amount}
          placeholder="Type here" 
          onChange={this.onChange}
          autoComplete='off'
          
          required
          fullWidth
          type="number"
          inputProps={{
            min: 0,
            max: wLimit
          }}
        />
        
        <TextField
          id="txFee"
          label="Transaction fee"
          margin="normal"
          
          value={this.props.txinfo.withdrawFee || "0"}
          InputProps={{
            readOnly: true,
          }}
          fullWidth
        />
        
        <TextField
          id="destination"
          label="Destination"
          helperText="ETH Account."
          margin="normal"
          value={this.props.txinfo.account || ""}
          onChange={this.onChange}
          
          InputProps={{
            readOnly: true,
          }}
          
          fullWidth
        />
      </Form>
    );
  }
}

export default withRouter(WithdrawForm);
