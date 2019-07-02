import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import TitleBarButton from "../../components/TitleBarButton";
import TitleBarTitle from "../../components/TitleBarTitle";
import { withRouter, Link  } from "react-router-dom";
import Button from "@material-ui/core/Button";

class TopupForm extends FormComponent {
  render() {
    if (this.props.txinfo === undefined) {
      return(<div></div>);
    }
    
    const extraButtons = <>
      <Button color="primary" type="button" disabled={false} >USE AXS WALLET</Button>
    </>;

    return(
      <Form
        submitLabel={this.props.submitLabel}
        extraButtons={extraButtons}
        onSubmit={this.onSubmit}
      >
        <TextField
          id="amount"
          label="Amount"
          //helperText="Send MXC amount from."
          margin="normal"
          value={this.props.txinfo.sendToken || ""}
          onChange={this.onChange}
          inputProps={{
            pattern: "[\\w-]+",
          }}
          required
          fullWidth
        />
        <TextField
          id="to"
          label="From Account"
          margin="normal"
          value={this.props.txinfo.from || ""}
          onChange={this.onChange}
          required
          fullWidth
        />
        <TextField
          id="to"
          label="To Account"
          margin="normal"
          value={this.props.txinfo.to || ""}
          onChange={this.onChange}
          required
          fullWidth
        />
        <TitleBarTitle to="/#" title="CHANGE ETH ACCOUNT" />
          
      </Form>
    );
  }
}

export default withRouter(TopupForm);
