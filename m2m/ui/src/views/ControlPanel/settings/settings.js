import React, { Component } from 'react';

import { Grid } from '@material-ui/core';
import { withRouter } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';

import TitleBar from '../../../components/TitleBar';
import TitleBarTitle from '../../../components/TitleBarTitle';
import styles from '../../ethAccount/EthAccountStyle';
import WithdrawStore from '../../../stores/WithdrawStore';
import SettingsStore from '../../../stores/SettingsStore';
import { ETHER } from '../../../util/Coin-type';
import i18n, { packageNS } from '../../../i18n';
import CardContent from "@material-ui/core/CardContent";
import Card from "@material-ui/core/Card";
import SettingsForm from "./settingsForm"


class SystemSettings extends Component {
	constructor() {
		super();
		this.state = {
			downlinkPrice: '',
			percentageShare: '',
			lbWarning: '',
			withdrawFee: ''
		};
		this.loadSettings = this.loadSettings.bind(this);
		this.saveSettings = this.saveSettings.bind(this);
	}

	componentDidMount() {
		this.loadSettings();
	}

	loadSettings() {
		try {
			const organizationID = 0;
			//this.setState({loading: true})

			WithdrawStore.getWithdrawFee(ETHER, organizationID, (resp) => {
				this.setState({ withdrawFee: resp.withdrawFee });
			});

			SettingsStore.getSystemSettings((resp) => {
				this.setState({
					downlinkPrice: resp.downlinkFee,
					percentageShare: resp.transactionPercentageShare,
					lbWarning: resp.lowBalanceWarning
				});
			});
		} catch (e) {}
	};

	saveSettings() {
		try {
			let bodyWF = {
				moneyAbbr: 'Ether',
				orgId: '0',
				withdrawFee: this.state.withdrawFee
			};

			let bodySettings = {
				downlinkFee: this.state.downlinkPrice,
				lowBalanceWarning: this.state.lbWarning,
				transactionPercentageShare: this.state.percentageShare
			};

			WithdrawStore.setWithdrawFee(ETHER, 0, bodyWF, (resp) => {});

			SettingsStore.setSystemSettings(bodySettings, (resp) => {});
		} catch (e) {}
	};

	handleChange = (name, event) => {
		this.setState({
			[name]: event.target.value
		});
	};

  render() {
    return(
      <Grid container spacing={24}>
        <Grid item xs={12} className={this.props.classes.divider}>
          <div className={this.props.classes.TitleBar}>
              <TitleBar className={this.props.classes.padding}>
							<TitleBarTitle title={i18n.t(`${packageNS}:menu.settings.system_settings`)} />
						</TitleBar>
          </div>
        </Grid>
        <Grid item xs={6} className={this.props.classes.column}>
          <Card className={this.props.classes.card}>
            <CardContent>
                <SettingsForm
                  submitLabel={i18n.t(`${packageNS}:menu.eth_account.confirm`)}
                  onSubmit={this.onSubmit}
				  downlinkPrice={this.state.downlinkPrice}
				  percentageShare={this.state.percentageShare}
				  lbWarning={this.state.lbWarning}
				  withdrawFee={this.state.withdrawFee}
                />

            </CardContent>
          </Card>

        </Grid>
        <Grid item xs={6}>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(SystemSettings));
