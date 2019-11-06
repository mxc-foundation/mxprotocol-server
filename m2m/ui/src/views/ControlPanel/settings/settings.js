import React, { Component } from 'react';

import { Grid, Card, Table, TableBody, TextField } from '@material-ui/core';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';
import HistoryStore from '../../../stores/HistoryStore';
import TitleBar from '../../../components/TitleBar';
import TitleBarTitle from '../../../components/TitleBarTitle';
import TitleBarButton from '../../../components/TitleBarButton';
import DataTable from '../../../components/DataTable';
import styles from './settingsStyle';
import Divider from '@material-ui/core/Divider';

class Settings extends Component {
	constructor(props) {
		super(props);
		this.getPage = this.getPage.bind(this);
	}

	getPage(limit, offset, callbackFunc) {}

	render() {
		return (
			<Grid container spacing={3} className={this.props.classes.root}>
				<Grid item xs={12}>
					<Grid item container xs={6} direction="column" className={this.props.classes.divider} padding={12}>
						<TitleBar>
							<TitleBarTitle title="System Settings" />
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
									title="Control Panel"
									className={this.props.classes.link}
								/>
								<TitleBarTitle component={Link} to="#" title="/" className={this.props.classes.link} />
								<TitleBarTitle
									component={Link}
									to="#"
									title="System Settings"
									className={this.props.classes.link}
								/>
							</TitleBar>
						</div>
					</Grid>

					<Grid item container direction="column" xs={6}>
						<TextField
							id="standard-number"
							label="Widthdraw Fee"
							className={this.props.classes.TextField}
							variant="filled"
							type="number"
							InputLabelProps={{
								shrink: true
							}}
							margin="normal"
						/>

						<TextField
							id="standard-number"
							label="Downlink Price"
							className={this.props.classes.TextField}
							variant="filled"
							type="number"
							InputLabelProps={{
								shrink: true
							}}
							margin="normal"
						/>

						<TextField
							id="standard-number"
							label="Percentage Share per transaction"
							className={this.props.classes.TextField}
							variant="filled"
							type="number"
							InputLabelProps={{
								shrink: true
							}}
							margin="normal"
						/>

						<TextField
							id="standard-number"
							label="Low Balance warning"
							className={this.props.classes.TextField}
							variant="filled"
							type="number"
							InputLabelProps={{
								shrink: true
							}}
							margin="normal"
						/>
					</Grid>
				</Grid>
			</Grid>
		);
	}
}

export default withStyles(styles)(withRouter(Settings));
