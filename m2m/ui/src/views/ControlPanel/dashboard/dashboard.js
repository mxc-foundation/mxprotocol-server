import React, { Component } from "react";

import {Grid,Card,Table,TableBody,TextField} from "@material-ui/core";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
import { withRouter, Link } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";
import HistoryStore from "../../../stores/HistoryStore";
import TitleBar from "../../../components/TitleBar";
import TitleBarTitle from "../../../components/TitleBarTitle";
import TitleBarButton from "../../../components/TitleBarButton";
import DataTable from "../../../components/DataTable";
import styles from "./dashboardStyle"

class Dashboard extends Component {
  constructor(props) {
    super(props);
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {

   

  }

  getRow(obj, index) {
    return(
      <TableRow key={index}>
        <TableCell>{obj.org}</TableCell>
        <TableCell>{obj.timestamp}</TableCell>
        <TableCell>{obj.value}</TableCell>
        <TableCell>{obj.type}</TableCell>
        <TableCell>{obj.income}</TableCell>
      </TableRow>
    );
  }

  render() {
    return(
      <Grid container spacing={3} className={this.props.classes.root}>
      <Grid item xs={12}>
        <TitleBar>
         <TitleBarTitle title="Welcome, SuperAdmin" />
        </TitleBar>
        </Grid>
   
        <Grid item xs={8}>
        
          <DataTable
            header={
              <TableRow>
                <TableCell>Org</TableCell>
                <TableCell>Timestamp</TableCell>
                <TableCell>Value</TableCell>
                <TableCell>Type</TableCell>
                <TableCell>Income</TableCell>
              </TableRow>
            }
            getPage={this.getPage}
            getRow={this.getRow}
          />
        </Grid>

        <Grid item xs={4}>
        <Grid container direction="column" spacing={10}>
        <Grid item xs={12}>
        <Card  className={this.props.classes.card}>
        <Table className={this.props.classes.cardTable}>
          <TableBody>
            <TableRow >
              <TableCell>Today's income</TableCell>
              <TableCell align="right">1.244MXC</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Monthly Balance</TableCell>
              <TableCell align="right"><span>1.244MXC</span></TableCell>
            </TableRow>
            <TableRow>
              <TableCell></TableCell>
              <TableCell align="right"><b>set alert</b></TableCell>
            </TableRow>
          </TableBody>
        </Table>

        </Card>
        </Grid>
        <Grid item container direction="column" xs={12}>

        <h4>General Settings</h4>
          <TextField
        id="standard-number"
        label="Widthdraw Fee"
        className={this.props.classes.TextField}
        variant="filled"
        type="number"
 
        InputLabelProps={{
          shrink: true,
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
          shrink: true,
        }}
        margin="normal"
      />
      
      <h4>System</h4>
        <Table className={this.props.classes.cardTable}>
          <TableBody>
            <TableRow>
              <TableCell>Monthly Downtime</TableCell>
              <TableCell align="right">3 hours</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Tickets opened</TableCell>
              <TableCell align="right"><b>1</b></TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Tickets closed</TableCell>
              <TableCell align="right"><b>5</b></TableCell>
            </TableRow>
          </TableBody>
        </Table>

        </Grid>
        </Grid>
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles) (withRouter(Dashboard));