import React, { Component } from 'react';

import { Grid, Table, TableBody, Switch,Button } from '@material-ui/core';
import TableCell from '@material-ui/core/TableCell';
import TableRow from '@material-ui/core/TableRow';
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from '@material-ui/core/styles';
import HistoryStore from '../../../stores/HistoryStore';
import TitleBar from '../../../components/TitleBar';
import TitleBarTitle from '../../../components/TitleBarTitle';
import DataTable from '../../../components/DataTable';
import styles from './organizationStyle';
import Cancel from 'mdi-material-ui/Cancel';

class UserList extends Component {
	constructor(props) {
		super(props);
		this.getPage = this.getPage.bind(this);
		this.getRow = this.getRow.bind(this);
	}

	getPage(limit, offset, callbackFunc) {}

	getRow(obj, index) {
		return (
			<TableRow key={index}>
			
			</TableRow>
		);
	}

	render() {
		return (
			<Grid container spacing={3} className={this.props.classes.root} justify="space-between" direction="row">
				<Grid item container xs={8} alignItems="flex-start">
					<Grid item xs={12}>
						<TitleBar>
							<TitleBarTitle title="Organization Name's User List" />
						</TitleBar>
					</Grid>

					<Grid item xs={12}>
						<DataTable
							header={
								<TableRow>
									<TableCell>User</TableCell>
									<TableCell>Type</TableCell>
									<TableCell>Last Access</TableCell>
									<TableCell>Permissions</TableCell>
									<TableCell />
								</TableRow>
							}
							getPage={this.getPage}
							getRow={this.getRow}
						/>
					</Grid>
				</Grid>

				<Grid item container xs={3} alignItems="flex-start"  className={this.props.classes.userListSettings}>
            <Grid item xs={12}>
              <h4>User Name Settings</h4>
						 	 <p className={this.props.classes.link}>Last active: today</p>
               <p className={this.props.classes.link}><Cancel /> Block user</p>
					
            </Grid>
						

					
						
            <Grid item xs={12}>
						<h4>Permissions</h4>
						<Table className={this.props.classes.cardTable}>
							<TableBody>
								<TableRow>
									<TableCell>Read</TableCell>
									<TableCell align="right">
										<Switch />
									</TableCell>
								</TableRow>
								<TableRow>
									<TableCell>Topup</TableCell>
									<TableCell align="right">
										<Switch />
									</TableCell>
								</TableRow>
								<TableRow>
									<TableCell>Add Gateway/Device</TableCell>
									<TableCell align="right">
										<Switch />
									</TableCell>
								</TableRow>
								<TableRow>
									<TableCell>Universal</TableCell>
									<TableCell align="right">
										<Switch />
									</TableCell>
								</TableRow>
							</TableBody>
						</Table>
            </Grid>
            <Grid container item xs={12} direction="row" justify="center" spacing={2}>

            <Grid item xs={6}>
                <Button variant="contained" className={this.props.classes.Button}>CANCEL</Button>
            </Grid>
           
            <Grid item xs={6}>
                <Button className={this.props.classes.Button}>CONFIRM</Button>
            </Grid>
   

            </Grid>
					</Grid>
				</Grid>
	
		);
	}
}

export default withStyles(styles)(withRouter(UserList));
