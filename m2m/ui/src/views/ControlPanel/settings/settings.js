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
import { ETHER } from "../../../util/Coin-type";

class Settings extends Component {
	constructor(props) {
		super(props);
    this.loadData = this.loadData.bind(this);
    this.state = {

    };
	}

  componentDidMount() {
    this.loadSettings();
  }

  loadSettings = async () => {
    try {
      const organizationID = 0;
      //this.setState({loading: true})
      WithdrawStore.getWithdrawFee(ETHER, organizationID, resp => {
        this.setState({withdrawFee:resp.withdrawFee});
      });
    
  }catch(e){

  }

}

  saveSettings = async()=>{
    try {
     
      let body = {
        "moneyAbbr": "Ether",
        "orgId": "0",
        "withdrawFee": this.state.withdrawFee 
      }
    
      WithdrawStore.setWithdrawFee(ETHER, 0, body, resp => {
        console.log(resp)
      });
    
  }catch(e){
  }
}

	render() {
		return (
			<Grid container spacing={3} className={this.props.classes.root}>
				<Grid item xs={12}>
					<Grid item container xs={6} direction="column" >
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

					<Grid item container direction="column" xs={6} className={this.props.classes.settingsForm}>
						<TextField
							id="withdrawFee"
							label="Widthdraw Fee"
							className={this.props.classes.TextField}
							variant="filled"
							type="number"
							InputLabelProps={{
								shrink: true
							}}
							margin="normal"
              value = {this.state.withdrawFee}
						/>

						<TextField
							id="downlinkPrice"
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
							id="percentageShare"
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
							id="lbWarning"
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
            <Grid container item xs={6} direction="row" justify="flex-end" spacing={2}>
				
						<Button variant="contained" className={this.props.classes.Button}>
							CANCEL
						</Button>
			
						<Button className={this.props.classes.Button}>SAVE CHANGES</Button>
					
				</Grid>
				</Grid>
			
			</Grid>
		);
	}
}

export default withStyles(styles)(withRouter(Settings));
