import React, { Component } from 'react';

import { Grid, Card, Table, TableBody, TextField, Button } from '@material-ui/core';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';

import TitleBar from '../../../components/TitleBar';
import TitleBarTitle from '../../../components/TitleBarTitle';
import TitleBarButton from '../../../components/TitleBarButton';
import DataTable from '../../../components/DataTable';
import styles from './settingsStyle';
import Divider from '@material-ui/core/Divider';
import WithdrawStore from '../../../stores/WithdrawStore';
import SettingsStore from '../../../stores/SettingsStore';
import { ETHER } from '../../../util/Coin-type';
import i18n, { packageNS } from '../../../i18n';
import NumberFormat from 'react-number-format';
import PropTypes from 'prop-types';

const NumberFormatMXC=(props)=> {
	const { inputRef, onChange, ...other } = props;

	return (
		<NumberFormat
			{...other}
			getInputRef={inputRef}
			onValueChange={(values) => {
				onChange({
					target: {
						value: values.value
					}
				});
			}}
			suffix=" MXC"
		/>
	);
}

const NumberFormatPerc =(props) =>{
	const { inputRef, onChange, ...other } = props;

	return (
		<NumberFormat
			{...other}
			getInputRef={inputRef}
			onValueChange={(values) => {
				onChange({
					target: {
						value: values.value
					}
				});
			}}
			suffix=" %"
		/>
	);
}

class SystemSettings extends Component {
	constructor(props) {
		super(props);

		this.state = {};
	}

	componentDidMount() {
		this.loadSettings();
	}

	loadSettings = async () => {
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

	saveSettings = async () => {
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
		return (
			<Grid container spacing={3} className={this.props.classes.root}>
				<Grid item xs={12}>
					<Grid item container xs={6} direction="column">
						<TitleBar>
							<TitleBarTitle title={i18n.t(`${packageNS}:menu.settings.system_settings`)} />
						</TitleBar>
						<Divider light={true} />
						<div className={this.props.classes.breadcrumb}>
							<TitleBar>
								<TitleBarTitle
									component={Link}
									to="#"
									title="M2M Wallet"
									className={this.props.classes.link}
								/>
								<TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link} />
								<TitleBarTitle
									component={Link}
									to="#"
									title={i18n.t(`${packageNS}:menu.settings.control_panel`)}
									className={this.props.classes.link}
								/>
								<TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link} />
								<TitleBarTitle
									component={Link}
									to="#"
									title={i18n.t(`${packageNS}:menu.settings.system_settings`)}
									className={this.props.classes.link}
								/>
							</TitleBar>
						</div>
					</Grid>

					<Grid item container direction="column" xs={6} className={this.props.classes.settingsForm}>
						<TextField
							id="withdrawFee"
							label={i18n.t(`${packageNS}:menu.settings.withdraw_fee`)}
							className={this.props.classes.TextField}
							variant="filled"
							InputLabelProps={{
								shrink: true
							}}
							InputProps={{
								inputComponent: NumberFormatMXC
							}}
							margin="normal"
							value={this.state.withdrawFee}
							onChange={(e) => this.handleChange('withdrawFee', e)}
						/>

						<TextField
							id="downlinkPrice"
							label={i18n.t(`${packageNS}:menu.settings.downlink_price`)}
							className={this.props.classes.TextField}
							variant="filled"
							InputLabelProps={{
								shrink: true
							}}
							InputProps={{
								inputComponent: NumberFormatMXC
							}}
							margin="normal"
							value={this.state.downlinkPrice}
							onChange={(e) => this.handleChange('downlinkPrice', e)}
						/>

						<TextField
							id="percentageShare"
							label={i18n.t(`${packageNS}:menu.settings.percentage_share`)}
							className={this.props.classes.TextField}
							variant="filled"
							InputLabelProps={{
								shrink: true
							}}
							InputProps={{
								inputComponent: NumberFormatPerc
							}}
							margin="normal"
							value={this.state.percentageShare}
							onChange={(e) => this.handleChange('percentageShare', e)}
						/>

						<TextField
							id="lbWarning"
							label={i18n.t(`${packageNS}:menu.settings.low_balance`)}
							className={this.props.classes.TextField}
							variant="filled"
							InputLabelProps={{
								shrink: true
							}}
							InputProps={{
								inputComponent: NumberFormatMXC
							}}
							margin="normal"
							value={this.state.lbWarning}
							onChange={(e) => this.handleChange('lbWarning', e)}
						/>
					</Grid>
					<Grid container item xs={6} direction="row" justify="flex-end" spacing={2}>
						<Button variant="contained" className={this.props.classes.Button} onClick={this.loadSettings}>
							{i18n.t(`${packageNS}:menu.settings.cancel`)}
						</Button>

						<Button className={this.props.classes.Button} onClick={this.saveSettings}>
							{i18n.t(`${packageNS}:menu.settings.save_changes`)}
						</Button>
					</Grid>
				</Grid>
			</Grid>
		);
	}
}

export default withStyles(styles)(withRouter(SystemSettings));
