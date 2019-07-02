import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import Button from "@material-ui/core/Button";
import InputAdornment from '@material-ui/core/InputAdornment';

import { withRouter } from "react-router-dom";

class WithdrawForm extends FormComponent {

  onChange = (event) => {
    const { id, value } = event.target;
    
    console.log(value)
    if(!((event.keyCode > 95 && event.keyCode < 106)
      || (event.keyCode > 47 && event.keyCode < 58) 
      || event.keyCode == 8)) {
        return false;
    }
    console.log(this.state.amount)

    this.setState({
      // [id]: value.replace(' MXC', '')
      [id]: value
    });
  }

  render() {
    if (this.props.txinfo === undefined) {
      return(<div>loading...</div>);
    }

    const extraButtons = <>
      <Button color="primary" type="button" disabled={false} >Cancel</Button>
    </>;

    return(
      <Form
        submitLabel={this.props.submitLabel}
        extraButtons={extraButtons}
        onSubmit={() => this.props.onSubmit({
          amount: this.state.amount,
          destination: this.state.destination,
          txFee: this.props.txinfo.withdrawFee
        })}
      >
        <TextField
          id="amount"
          //bgcolor="primary.main"
          label="Amount"
          //helperText="The name may only contain words, numbers and dashes."
          margin="normal"
          //value={this.state.amount + ' MXC'}
          value={this.state.amount}
          onChange={this.onChange}
          className={this.props.classes.root}
          required
          fullWidth
          type="number"
          inputProps={{
            maxLength: 4,
            min: 0,
            max: 1000
          }}
        />
        
        <TextField
        id="txFee"
          label="Transaction fee"
          margin="normal"
          value={this.props.txinfo.withdrawFee || ""}
            className={this.props.classes.root}

          required
          fullWidth
        />
        
        <TextField
          id="destination"
          label="Destination"
          helperText="ETH Account."
          margin="normal"
          value={this.state.destination}
          onChange={this.onChange}
          className={this.props.classes.root}
          InputProps={{
            pattern: "[\\w-]+",
          }}
          
          required
          fullWidth
        />
      </Form>
    );
  }
}

export default withRouter(WithdrawForm);
