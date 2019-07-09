import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import TitleBarTitle from "../../components/TitleBarTitle";
import { withRouter, Link  } from "react-router-dom";
import Button from "@material-ui/core/Button";

class TopupForm extends FormComponent {

  handleOpenAXS = () => {
    window.location.replace(`http://www.google.com`);
  } 

  render() {
    if (this.props.reps === undefined) {
      return(<div></div>);
    }
    //console.log('this.props.reps', this.props.reps);
    const extraButtons = <>
      <Button color="primary" onClick={this.handleOpenAXS} type="button" disabled={false}>USE AXS WALLET</Button>
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

          value={this.props.reps.amount || ""}
          InputProps={{
            readOnly: true,

          }}

          fullWidth
        />
        <TextField
          id="to"
          label="From"
          margin="normal"
          value={this.props.reps.from || ""}
          InputProps={{
            readOnly: true,
          }}
          fullWidth
        />
        <TextField
          id="to"
          label="To"
          margin="normal"
          value={this.props.reps.to || ""}
          InputProps={{
            readOnly: true,
          }}
          fullWidth
        />
        <TitleBarTitle component={Link} to={`/modify-account/${this.props.orgId}`} title="CHANGE ETH ACCOUNT" />
      </Form>
    );
  }
}

export default withRouter(TopupForm);
