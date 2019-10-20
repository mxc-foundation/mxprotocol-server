import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import TitleBarTitle from "../../components/TitleBarTitle";
import { withRouter, Link  } from "react-router-dom";
import Button from "@material-ui/core/Button";

class TopupForm extends FormComponent {

  handleOpenAXS = () => {
    window.location.replace(`http://wallet.mxc.org/`);
  } 

  render() {
    const extraButtons = <>
      <Button color="primary.main" onClick={this.handleOpenAXS} type="button" disabled={false}>USE AXS WALLET</Button>
    </>;
    
    if (this.props.reps === undefined) {
      return(
        <Form
          submitLabel={this.props.submitLabel}
          extraButtons={extraButtons}
          onSubmit={this.onSubmit}
        >
          <TitleBarTitle component={Link} to={'#'} title="THERE IS NO DATA TO DISPLAY." />
        </Form>
      );
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        extraButtons={extraButtons}
        onSubmit={this.onSubmit}
      >
        <TextField
          id="to"
          label="From"
          margin="normal"
          value={this.props.reps.account || "Can not find any account."}
          InputProps={{
            readOnly: true,
          }}
          fullWidth
        />
        <TextField
          id="to"
          label="To"
          margin="normal"
          value={this.props.reps.superNodeAccount || "Can not find any account."}
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
