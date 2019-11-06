import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
import Spinner from "../../components/ScaleLoader";
import { SUPER_ADMIN } from "../../util/M2mUtil";
import SupernodeStore from "../../stores/SupernodeStore";
import styles from "./superNodeEthStyle"
import { ETHER } from "../../util/Coin-type";


class SuperNodeEth extends Component {
  constructor(props) {
    super(props);
    this.state = {
      activeAccount:'0',
      loading:false
    };
 
  }
  
  componentDidMount() {
    this.loadData();
  }
  
  componentDidUpdate() {
    this.loadData();
  }
  
  loadData() {
   
      SupernodeStore.getSuperNodeActiveMoneyAccount(ETHER, SUPER_ADMIN, resp => {
        this.setState({
          activeAccount: resp.supernodeActiveAccount,
        });
      });
    
  }


  render() {
    return(
      <Grid container spacing={24}>
        <Spinner on={this.state.loading}/>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
                <TitleBar className={this.props.classes.padding}>
                  <TitleBarTitle title="Control Panel" />
                </TitleBar>
             
            </div>
        </Grid>
        <Grid item xs={6} className={this.props.classes.column}>
         
          <TextField
          id="to"
          label="Super Node ETH Account"
          margin="normal"
          value={this.state.activeAccount ||"Can not find any account."}
          InputProps={{
            readOnly: true,
          }}
          fullWidth
        />
          <TitleBarTitle component={Link} to={`/modify-account/${SUPER_ADMIN}`} title="CHANGE SUPER NODE ETH ACCOUNT" />
        </Grid>
        <Grid item xs={6}>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(SuperNodeEth));