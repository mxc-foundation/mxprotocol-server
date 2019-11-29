import React, { Component } from "react";
import { Link } from "react-router-dom";
import Grid from "@material-ui/core/Grid";
import i18n, { packageNS } from '../../i18n';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";

import GatewayForm from "./GatewayForm";
import Modal from "../../components/Modal";
import Divider from '@material-ui/core/Divider';
import GatewayStore from "../../stores/GatewayStore.js";
import WalletStore from "../../stores/WalletStore.js";
//import WithdrawBalanceInfo from "./WithdrawBalanceInfo";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import styles from "./GatewayStyle"
import { ETHER } from "../../util/Coin-type"
import { SUPER_ADMIN } from "../../util/M2mUtil"

function getDlPrice(orgId) {
    return new Promise((resolve, reject) => {
        WalletStore.getDlPrice(orgId, resp => {
            resolve(resp.downLinkPrice);
        });
    });
}

class GatewayLayout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      modal: null,
      downlinkFee: 0
    };
  }

  loadData = async () => {
    //console.log(this.props);
  }

  loadData = async () => {
    try {
      const orgId = this.props.match.params.organizationID;
      this.setState({loading: true})
      var downlinkFee = await getDlPrice(orgId);

      this.setState({
        downlinkFee
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
    if (this.props === oldProps) {
      return;
    }
    this.loadData();
  }
  
  showModal(modal) {
    this.setState({ modal });
  }

  onSubmit = (e, apiWithdrawReqRequest) => {
    e.preventDefault();
  }

  handleCloseModal = () => {
    this.setState({
      modal: null
    })
  }

  onSelectChange = (gateway) => {
    const { gwId, gwMode } = gateway;
    //console.log('device', device);
    GatewayStore.setGatewayMode(this.props.match.params.organizationID, gwId, gwMode, data => {
      this.props.history.push(`/gateway/${this.props.match.params.organizationID}`);
    });
  }

  onConfirm = (data) => {
    
  }
  
  render() {
    return (
        <Grid container spacing={24} className={this.props.classes.backgroundColor}>
            {this.state.modal && 
            <Modal title={i18n.t(`${packageNS}:menu.messages.confirmation`)} description={i18n.t(`${packageNS}:menu.messages.confirmation_text`)} onClose={this.handleCloseModal} open={!!this.state.modal} data={this.state.modal} onConfirm={this.onConfirm} />}
            <Grid item xs={12} className={this.props.classes.divider}>
            <div className={this.props.classes.TitleBar}>
                <TitleBar className={this.props.classes.padding}>
                  <TitleBarTitle title={i18n.t(`${packageNS}:menu.gateways.gateways`)} />
                </TitleBar>    
                {/* <Divider light={true}/> */}
                <div className={this.props.classes.between}>
                <TitleBar>
                  <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                  <TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link}/>
                  <TitleBarTitle component={Link} to="#" title={i18n.t(`${packageNS}:menu.gateways.gateways`)} className={this.props.classes.link}/>
                </TitleBar>
                </div>
            </div>
          </Grid>
        <Grid item xs={12} className={this.props.classes.divider}>
          <Grid item xs={6} className={this.props.classes.divider}>
            <GatewayForm
              submitLabel={i18n.t(`${packageNS}:menu.gateways.gateways`)}
              downlinkFee={this.state.downlinkFee}
              onSelectChange={this.onSelectChange}
              onSubmit={this.onSubmit}
            />
          </Grid>
        </Grid>
        </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(GatewayLayout));
