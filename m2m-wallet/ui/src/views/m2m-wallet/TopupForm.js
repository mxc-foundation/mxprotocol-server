import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import TitleBarButton from "../../components/TitleBarButton";
import LinkI from "mdi-material-ui/Link";
import { withRouter } from "react-router-dom";

class TopupForm extends FormComponent {
  render() {
    if (this.props.txinfo === undefined) {
      return(<div></div>);
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
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
        <TitleBarButton
            key={1}
            label="CHANGE ETH ACCOUNT"
            icon={<LinkI />}
            color="secondary"
            /* onClick={this.deleteOrganization} */
        />
          
      </Form>
    );
  }
}

export default withRouter(TopupForm);
