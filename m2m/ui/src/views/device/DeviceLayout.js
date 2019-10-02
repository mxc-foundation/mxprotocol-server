import React, { Component } from "react";
import { Link } from "react-router-dom";
import Divider from '@material-ui/core/Divider';
import Grid from "@material-ui/core/Grid";
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Typography from '@material-ui/core/Typography';

import DeviceStore from "../../stores/DeviceStore.js";
import DeviceForm from "./DeviceForm";
import Modal from "../../components/Modal";
//import WithdrawBalanceInfo from "./WithdrawBalanceInfo";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import styles from "./DeviceStyle"
import { DV_INACTIVE, DV_FREE_GATEWAYS_LIMITED, DV_WHOLE_NETWORK } from "../../util/data"
import { CONFIRMATION, CONFIRMATION_TEXT, INVALID_ACCOUNT, INVALID_AMOUNT } from "../../util/Messages"

class DeviceLayout extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loading: false,
      modal: null,
      mod: null
    };
  }

  loadData() {
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

  onConfirm = (data) => {
    
  }
  
  onSelectChange = (device) => {
    const { dvId, dvMode } = device;
    
    DeviceStore.setDeviceMode(this.props.match.params.organizationID, dvId, dvMode, data => {
      this.props.history.push(`/device/${this.props.match.params.organizationID}`);
    });
  }

  onSwitchChange = (device) => {
    const { dvId, on } = device;
    
    let mod = DV_FREE_GATEWAYS_LIMITED;
    if(!on){
     mod = DV_INACTIVE;   
    }
    DeviceStore.setDeviceMode(this.props.match.params.organizationID, dvId, mod, data => {
      this.props.history.push(`/device/${this.props.match.params.organizationID}`);
    });
  }

  render() {
    return (
        <Grid container spacing={24} className={this.props.classes.backgroundColor}>
            {this.state.modal && 
            <Modal title={CONFIRMATION} description={CONFIRMATION_TEXT} onClose={this.handleCloseModal} open={!!this.state.modal} data={this.state.modal} onConfirm={this.onConfirm} />}
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              <TitleBar className={this.props.classes.padding}>
                <TitleBarTitle title="Devices" />
              </TitleBar>
              <Divider light={true}/>
              <div className={this.props.classes.breadcrumb}>
              <TitleBar>
                <TitleBarTitle component={Link} to="#" title="M2M Wallet" className={this.props.classes.link}/> 
                <TitleBarTitle title="/" className={this.props.classes.navText}/>
                <TitleBarTitle component={Link} to="#" title="Devices" className={this.props.classes.link}/>
              </TitleBar>
              </div>
          </div>

        </Grid>
        <Grid item xs={12} className={this.props.classes.divider}>
        <Grid item xs={6}>
          <DeviceForm
            submitLabel="Devices"
            onSubmit={this.onSubmit}
            onSelectChange={this.onSelectChange}
            onSwitchChange={this.onSwitchChange}
          />
          </Grid>
        </Grid>
        </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(DeviceLayout));
