import React from "react";

import TextField from '@material-ui/core/TextField';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
import Button from "@material-ui/core/Button";
import Typography from '@material-ui/core/Typography';
import StakeStore from "../../stores/StakeStore";
import { REVENUE_RATE, AMOUNT, CONFIRM_STAKE, CONFIRM_UNSTAKE } from "../../util/Messages"
import Spinner from "../../components/ScaleLoader"
import { withRouter, Link  } from "react-router-dom";

class StakeForm extends FormComponent {
  
  state = {
    amount: ''
  }

  componentDidMount(){
    this.loadData();
  }
  
  loadData = () => {
     
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
    /* if (this.props.txinfo === undefined) {
      return(<Spinner on={this.state.loading}/>);
    } */
    const extraButtons = <>
      <Button color="primary.main" onClick={this.handleOpenAXS} type="button" disabled={false}>CANCEL</Button>
    </>;

    return(
        <Form
            submitLabel={ this.props.isUnstake ? CONFIRM_UNSTAKE: CONFIRM_STAKE }
            extraButtons={extraButtons}
            onSubmit={(e) => this.props.confirmStake(e, {
            amount: parseFloat(this.props.amount),
            })}
        >
            <Typography  /* className={this.props.classes.title} */ gutterBottom>
                {this.props.label}
            </Typography>
            <Divider light={true}/>
            <TextField
                id="amount"
                label={AMOUNT}
                margin="normal"
                value={this.props.amount}
                placeholder="Type here" 
                onChange={this.props.onChange}
                autoComplete='off'
                
                required
                fullWidth
                type="number"
                inputProps={{
                    min: 0,
                }}
            />
            
            <TextField
                id="txFee"
                label={REVENUE_RATE}
                margin="normal"
                
                value={this.props.revRate}
                InputProps={{
                    readOnly: true,
                }}
                fullWidth
            />
        </Form>
    );
  }
}

export default withRouter(StakeForm);
