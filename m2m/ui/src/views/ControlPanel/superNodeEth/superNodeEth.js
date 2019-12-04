import React, { Component } from "react";
import { withRouter } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";
import Grid from '@material-ui/core/Grid';
import i18n, { packageNS } from '../../../i18n';
import TitleBar from "../../../components/TitleBar";
import TitleBarTitle from "../../../components/TitleBarTitle";
import SessionStore from "../../../stores/SessionStore";
import SupernodeStore from "../../../stores/SupernodeStore";
import ModifyEthAccountForm from "../../ethAccount/ModifyEthAccountForm";
import NewEthAccountForm from "../../ethAccount/NewEthAccountForm";
import styles from "../../ethAccount/EthAccountStyle"
import { ETHER } from "../../../util/Coin-type";
import { SUPER_ADMIN } from "../../../util/M2mUtil";
import CardContent from "@material-ui/core/CardContent";
import Card from "@material-ui/core/Card";


class SuperNodeEth extends Component {
    constructor() {
        super();
        this.state = {
          activeAccount:'0'
        };
        this.loadData = this.loadData.bind(this);
    }
  
    componentDidMount() {
        this.loadData();
    }

    loadData() {
      SupernodeStore.getSuperNodeActiveMoneyAccount(ETHER, SUPER_ADMIN, resp => {
        this.setState({
          activeAccount: resp.supernodeActiveAccount,
        });
      });

    }

    componentDidUpdate(oldProps) {
        if (this.props === oldProps) {
            return;
        }

        this.loadData();
    }


    verifyUser(resp) {
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

    modifyAccount(req, orgId) {
        req.moneyAbbr = ETHER;
        return new Promise((resolve, reject) => {
            SupernodeStore.addSuperNodeMoneyAccount(req, orgId, resp => {
                resolve(resp);
            })
        });
    }

    onSubmit = async (resp) => {
      
      try {
        if(resp.username !== SessionStore.getUsername() ){
          alert(`${i18n.t(`${packageNS}:menu.withdraw.incorrect_username_or_password`)}`);
          return false;
        }
        const isOK = await this.verifyUser(resp);
        
        if(isOK) {
          const res = await this.modifyAccount(resp, SUPER_ADMIN);
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
    return(
      <Grid container spacing={24}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title={i18n.t(`${packageNS}:menu.eth_account.eth_account`)} />
              </TitleBar>
          </div>
        </Grid>
        <Grid item xs={6} className={this.props.classes.column}>
          <Card className={this.props.classes.card}>
            <CardContent>
              {this.state.activeAccount &&
                <ModifyEthAccountForm
                  submitLabel={i18n.t(`${packageNS}:menu.eth_account.confirm`)}
                  onSubmit={this.onSubmit}
                  activeAccount={this.state.activeAccount}
                />
              }
              {!this.state.activeAccount &&  
                <NewEthAccountForm
                  submitLabel={i18n.t(`${packageNS}:menu.eth_account.confirm`)}
                  onSubmit={this.onSubmit}
                />
              }
            </CardContent>
          </Card>
          
        </Grid>
        <Grid item xs={6}>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(SuperNodeEth));