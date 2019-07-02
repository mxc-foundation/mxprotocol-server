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
    if (this.state.object === undefined) {
      return(<div></div>);
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={() => this.props.onSubmit({
          newaccount: this.state.newaccount
        })}
      >
        <TextField
          id="newaccount"
          label="New account"
          margin="normal"
          value={this.state.newaccount}
          placeholder="Type here" 
          onChange={this.onChange}
          inputProps={{
            pattern: "[\\w-]+",
          }}
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
