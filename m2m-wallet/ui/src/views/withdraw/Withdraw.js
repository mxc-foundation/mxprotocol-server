import React, { Component } from "react";
import { Link } from "react-router-dom";

import Grid from "@material-ui/core/Grid";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import MoneyStore from "../../stores/MoneyStore";
import WithdrawStore from "../../stores/WithdrawStore";
import SupernodeStore from "../../stores/SupernodeStore";
import WalletStore from "../../stores/WalletStore";
import WithdrawForm from "./WithdrawForm";
import Modal from "./Modal";
import WithdrawBalanceInfo from "./WithdrawBalanceInfo";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import Divider from '@material-ui/core/Divider';
import styles from "./WithdrawStyle"



const coinType = "Ether";

function formatNumber(number) {
  return number.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

function loadWithdrawFee(coinType) {
  return new Promise((resolve, reject) => {
    WithdrawStore.getWithdrawFee(coinType,
      resp => {
        Object.keys(resp).forEach(attr => {
          const value = resp[attr];

          if (typeof value === 'number') {
            resp[attr] = formatNumber(value);
          }
        });
        resp.moneyAbbr = coinType;
        resolve(resp);
      })
  });
}

function loadCurrentAccount(coinType, organizationID) {
  
  return new Promise((resolve, reject) => {
    if (organizationID == 0) {
      SupernodeStore.getSuperNodeActiveMoneyAccount(coinType, resp => {
        resolve(resp.supernodeActiveAccount);
        
      });
    }else{
      MoneyStore.getActiveMoneyAccount(coinType, organizationID, resp => {
        resolve(resp.activeAccount);
        
      });
    }
  });
}

      
function loadWalletBalance(organizationID) {
  return new Promise((resolve, reject) => {
    WalletStore.getWalletBalance(organizationID,
      resp => {
        Object.keys(resp).forEach(attr => {
          const value = resp[attr];
  
          if (typeof value === 'number') {
            resp[attr] = formatNumber(value);
          }
        });
        resolve(resp);
      });
  });
}

class Withdraw extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      modal: null
    };
  }

  loadData = async () => {
    try {
      
      this.setState({loading: true})
      var result = await loadWithdrawFee(coinType);
      var wallet = await loadWalletBalance(this.props.match.params.organizationID);
      var account = await loadCurrentAccount(coinType, this.props.match.params.organizationID);
      
      /* this.setState({
        activeAccount: resp.supernodeActiveAccount,
      }); */

      const txinfo = {};
      txinfo.withdrawFee = result.withdrawFee;
      txinfo.balance = wallet.balance;
      txinfo.account = account;

      this.setState({
        txinfo
      });
      this.setState({loading: false})
    } catch (error) {
      this.setState({loading: false})
      console.error(error);
      this.setState({ error });
    }
  }

  componentDidMount() {
    this.loadData();
  }

  componentDidUpdate(oldProps) {
    if (this.props.match.url === oldProps.match.url) {
      return;
    }
  }

  showModal(modal) {
    this.setState({ modal });
  }

  onSubmit = (e, apiWithdrawReqRequest) => {
    e.preventDefault();
    this.showModal(apiWithdrawReqRequest);
  }

  handleCloseModal = () => {
    this.setState({
      modal: null
    })
  }

  onConfirm = (data) => {
    this.setState({loading: true});
    WithdrawStore.WithdrawReq(data, resp => {
      this.setState({loading: false});
    });
  }

  render() {
    return (
      <Grid container spacing={24} className={this.props.classes.backgroundColor}>
        {this.state.modal && 
          <Modal title="title" description="description" onClose={this.handleCloseModal} open={!!this.state.modal} data={this.state.modal} onConfirm={this.onConfirm} />}
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title="Withdraw" />
              </TitleBar>
              <Divider light={true}/>
              <div className={this.props.classes.breadcrumb}>
              <TitleBar>
                <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                <TitleBarTitle title="/" className={this.props.classes.navText}/>
                <TitleBarTitle component={Link} to="#" title="Withdraw" className={this.props.classes.link}/>
              </TitleBar>
              </div>
          </div>

        </Grid>
        <Grid item xs={6} className={this.props.classes.divider}></Grid>
        <Grid item xs={12} className={this.props.classes.divider}>

        </Grid>
        <Grid item xs={6}>
          <WithdrawForm
            submitLabel="Withdraw"
            txinfo={this.state.txinfo} {...this.props}
            onSubmit={this.onSubmit}
          />
        </Grid>
        <Grid item xs={2}>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(Withdraw));