import React from "react";

import TextField from '@material-ui/core/TextField';
import i18n, { packageNS } from '../../i18n';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import Divider from '@material-ui/core/Divider';
import Button from "@material-ui/core/Button";
import Typography from '@material-ui/core/Typography';
import InputAdornment from '@material-ui/core/InputAdornment';
import StakeStore from "../../stores/StakeStore";
//import Spinner from "../../components/ScaleLoader"
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import styles from "./StakeStyle"

class StakeForm extends FormComponent {
  
  state = {
    amount: '',
    revenue_rate: 0
  }

  componentDidMount(){
    this.loadData();
  }
  
  loadData = () => {
    const resp = StakeStore.getStakingPercentage(this.props.match.params.organizationID);
    resp.then((res) => {
      let revenue_rate = 0;
      revenue_rate = res.stakingPercentage;
      if(revenue_rate){
        this.setState({
          revenue_rate
        })
      }
    })
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
      <Button  variant="outlined" color="inherit" onClick={this.handleOpenAXS} type="button" disabled={false}>{i18n.t(`${packageNS}:menu.staking.cancel`)}</Button>
    </>;

    return(
        <Form
            submitLabel={ this.props.isUnstake ? i18n.t(`${packageNS}:menu.messages.confirm_unstake`) : i18n.t(`${packageNS}:menu.messages.confirm_stake`) }
            extraButtons={extraButtons}
            onSubmit={(e) => this.props.confirm(e, {
              amount: parseFloat(this.props.amount),
              action: this.props.isUnstake
            })}
        >
            <Typography  /* className={this.props.classes.title} */ gutterBottom>
                {this.props.label}
            </Typography>
            <Divider light={true}/>
            <TextField
                id="amount"
                label={i18n.t(`${packageNS}:menu.common.amount`)}
                margin="normal"
                value={this.props.amount}
                onChange={this.props.onChange}
                required={!this.props.isUnstake}
                InputProps={{
                  min: 0,
                    readOnly: this.props.isUnstake,
                    endAdornment: <InputAdornment position="end">MXC</InputAdornment>,
                }}
                fullWidth
            />
            <TextField
                id="revRate"
                label={i18n.t(`${packageNS}:menu.messages.revenue_rate`)}
                margin="normal"
                
                value={this.state.revenue_rate}
                InputProps={{
                    readOnly: true,
                    endAdornment: <InputAdornment position="end">{i18n.t(`${packageNS}:menu.staking.monthly`)}</InputAdornment>,
                }}
                fullWidth
            />
        </Form>
    );
  }
}

export default withStyles(styles)(withRouter(StakeForm));
