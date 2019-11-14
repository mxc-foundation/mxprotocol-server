import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";

import Grid from '@material-ui/core/Grid';
import i18n, { packageNS } from '../../i18n';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Divider from '@material-ui/core/Divider';
import MoneyStore from "../../stores/MoneyStore";
import SessionStore from "../../stores/SessionStore";
import SupernodeStore from "../../stores/SupernodeStore";
import ModifyEthAccountForm from "./ModifyEthAccountForm";
import NewEthAccountForm from "./NewEthAccountForm";
import styles from "./EthAccountStyle";
import { ETHER } from "../../util/Coin-type";
import { SUPER_ADMIN } from "../../util/M2mUtil";

function verifyUser (resp) {
  const login = {};
  login.username = resp.username;
  login.password = resp.password;
  
  return new Promise((resolve, reject) => {
    SessionStore.login(login, (resp) => {
      if(resp){
        resolve(resp);
      } else {
        alert(`${i18n.t(`${packageNS}:menu.withdraw.incorrect_username_or_password`)}`);
        return false;
      }
    })
  });
}

function modifyAccount (req, orgId) {
  req.moneyAbbr = ETHER;
  req.orgId = orgId;
  return new Promise((resolve, reject) => {
    MoneyStore.modifyMoneyAccount(req, resp => {
      resolve(resp);
    })
  });
}

function createAccount (req, orgId) {
  req.moneyAbbr = ETHER;
  return new Promise((resolve, reject) => {
    SupernodeStore.addSuperNodeMoneyAccount(req, orgId, resp => {
      resolve(resp);
    })
  });
}

class ModifyEthAccount extends Component {
  constructor() {
      super();
      this.state = {};
      this.loadData = this.loadData.bind(this);
    }
    
    componentDidMount() {
      this.loadData();
    }
    
    loadData() {
      const orgId = this.props.match.params.organizationID;

      if (orgId === SUPER_ADMIN) {
        SupernodeStore.getSuperNodeActiveMoneyAccount(ETHER, orgId, resp => {
          this.setState({
            activeAccount: resp.supernodeActiveAccount,
          });
        });
      }else{
        MoneyStore.getActiveMoneyAccount(ETHER, orgId, resp => {
          this.setState({
            activeAccount: resp.activeAccount,
          });
        });
      }
    }

    componentDidUpdate(oldProps) {
      if (this.props === oldProps) {
        return;
      }

      this.loadData();
    }

    onSubmit = async (resp) => {
      const orgId = this.props.match.params.organizationID;
      
      try {
        if(resp.username !== SessionStore.getUsername() ){
          alert(`${i18n.t(`${packageNS}:menu.withdraw.incorrect_username_or_password`)}`);
          return false;
        }
        const isOK = await verifyUser(resp);
        
        if(orgId == 0 && isOK) {
          const res = await createAccount(resp, orgId);
          if(res.status){
            window.location.reload();
          }
        }else{
          const res = await modifyAccount(resp, orgId);
          if(res.status){
            window.location.reload();
          }
        } 
      } catch (error) {
        console.error(error);
        this.setState({ error });
      }
    } 

  render() {
    
    const organizationID = this.props.match.params.organizationID;
    const isAdmin = (organizationID === SUPER_ADMIN)?true:false;
    
    return(
      <Grid container spacing={24}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
                <TitleBar className={this.props.classes.padding}>
                  <TitleBarTitle title={i18n.t(`${packageNS}:menu.withdraw.eth_account`)}  />
                </TitleBar>
                {/* <Divider light={true}/>
                <div className={this.props.classes.breadcrumb}>
                <TitleBar>
                  <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                  <TitleBarTitle title="/" className={this.props.classes.navText}/>
                  <TitleBarTitle component={Link} to="#" title="ETH Account" className={this.props.classes.link}/>
                </TitleBar>
                </div> */}
            </div>
        </Grid>
        <Grid item xs={6} className={this.props.classes.column}>
          {this.state.activeAccount &&
            <ModifyEthAccountForm
              submitLabel={i18n.t(`${packageNS}:menu.withdraw.confirm`)}
              onSubmit={this.onSubmit}
              activeAccount={this.state.activeAccount}
            />
          }
          {!this.state.activeAccount &&  
          <NewEthAccountForm
            submitLabel={i18n.t(`${packageNS}:menu.withdraw.confirm`)}
            onSubmit={this.onSubmit}
            //isAdmin={isAdmin}
          />
          }
        </Grid>
        <Grid item xs={6}>
          
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(ModifyEthAccount));