import React from "react";

import TextField from '@material-ui/core/TextField';
import { Redirect } from 'react-router-dom'
import clsx from 'clsx'
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import Grid from "@material-ui/core/Grid";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import TitleBarButton from "../../components/TitleBarButton";
import Typography from '@material-ui/core/Typography';
import StakeForm from "./StakeForm";
import Button from "@material-ui/core/Button";
//import Button from "@material-ui/core/Button";
import Spinner from "../../components/ScaleLoader";
import { YOUR_STAKE, STAKE_SET_SUCCESS, DISMISS, LEARN_MORE, MXC, UNSTAKE, STAKE, HISTORY } from "../../util/Messages"
import { EXT_URL_STAKE } from "../../util/Data"
import { withStyles } from "@material-ui/core/styles";
import { withRouter, Link  } from "react-router-dom";
import styles from "./StakeStyle"
import { Divider } from "@material-ui/core";

class SetStake extends FormComponent {
  
  state = {
    amount: 0,
    revRate: 0,
  }

  componentWillMount(){
    if(false){
      this.props.history.push(`/stake/${this.props.match.params.organizationID}`);
    }
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

  confirmStake = (e, amount) => {
    e.preventDefault();
    console.log(amount);
    //const resp = SessionStore.getProfile();
    /* resp.then((res) => {
      let orgId = SessionStore.getOrganizationID();
      const isBelongToOrg = res.body.organizations.some(e => e.organizationID === SessionStore.getOrganizationID());
  
      OrganizationStore.get(orgId, resp => {
        openM2M(resp.organization, isBelongToOrg, '/withdraw');
      });
    }) */

  } 

  render() {
    /* if (this.props.txinfo === undefined) {
      return(<Spinner on={this.state.loading}/>);
    } */
    
    return(
        <Grid container spacing={24} className={this.props.classes.backgroundColor}>
            <Grid item xs={12} md={12} lg={12} className={this.props.classes.divider}>
                <div className={this.props.classes.TitleBar}>
                    {/* <TitleBar className={this.props.classes.padding}>
                        <TitleBarTitle title="Stake" />
                    </TitleBar> */}    
                    {/* <Divider light={true}/> */}
                    <div className={this.props.classes.between}>
                    <TitleBar>
                        <TitleBarTitle title={STAKE} />
                        {/* <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                        <TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link}/>
                        <TitleBarTitle component={Link} to="#" title="Devices" className={this.props.classes.link}/> */}
                    </TitleBar>
                    
                    <Button variant="outlined" color="inherit" onClick={this.handleOpenAXS} type="button" disabled={false}>{HISTORY}</Button>
                    </div>
                </div>
            </Grid>
            <Grid item xs={6} lg={6} spacing={24} className={this.props.classes.pRight}>
                <StakeForm label={UNSTAKE} onChange={this.onChange} amount={this.state.amount} revRate={this.state.revRate} confirmStake={this.confirmStake} />
            </Grid>
            <Grid item xs={6} lg={6} spacing={24} className={this.props.classes.pLeft}>
                <div className={clsx(this.props.classes.urStake, this.props.classes.between)}>
                    <Typography  /* className={this.props.classes.title} */ gutterBottom>
                        {YOUR_STAKE}
                    </Typography>
                    <Typography  /* className={this.props.classes.title} */ gutterBottom>
                        200 {MXC}
                    </Typography>
                </div>
                <div>
                    <div className={this.props.classes.infoBox}>
                    <Typography  /* className={this.props.classes.title} */ gutterBottom>
                        {STAKE_SET_SUCCESS}
                    </Typography>
                    <div className={this.props.classes.between}>
                        <Typography className={this.props.classes.title} gutterBottom>
                        {DISMISS}
                        </Typography>
                        <TitleBarTitle component={Link} to={EXT_URL_STAKE} title={LEARN_MORE} />
                    </div>
                    </div>
                </div>
            </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(SetStake));
