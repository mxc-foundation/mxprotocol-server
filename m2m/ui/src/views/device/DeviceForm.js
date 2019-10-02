import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableRow from "@material-ui/core/TableRow";
//import Wallet from "mdi-material-ui/OpenInNew";
//import Typography from '@material-ui/core/Typography';
import TableCellExtLink from "../../components/TableCellExtLink";
import { DV_MODE_OPTION } from "../../util/data"

import { withRouter } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";
import DeviceStore from "../../stores/DeviceStore.js";
import TitleBar from "../../components/TitleBar";
import NativeSelects from "../../components/NativeSelects";
import SwitchLabels from "../../components/Switch";
import TitleBarButton from "../../components/TitleBarButton";
import DataTable from "../../components/DataTable";
import Admin from "../../components/Admin";

const styles = {
  flex: {
      display: 'flex',
      alignItems: 'center',
  },
  flex2: {
    left: 'calc(100%/3)',
    
  },
  maxW:{
    maxWidth: 120
  }
};
//<Wallet color="primary.main"  />
class DeviceForm extends Component {
  constructor(props) {
    super(props);
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
    this.state = {
      haveGateway: false
    };
  }
  
  

  getPage(limit, offset, callbackFunc) {
    DeviceStore.getDeviceList(this.props.match.params.organizationID, offset, limit, data => {
        callbackFunc({
            totalCount: parseInt(data.count),
            result: data.devProfile
          });
      }); 
  }

  onSelectChange = (e) => {
    const device = {
      dvId: e.dvId, 
      dvMode: e.target.value
    }
  
    this.props.onSelectChange(device);
  }

  onSwitchChange = (dvId, available) => {
    const device = {
      dvId, 
      available
    }
  
    this.props.onSwitchChange(device);
  }

  getRow(obj, index) {
    const url = `/#/organizations/${this.props.match.params.organizationID}/devices/${obj.mac}`;
    let dValue = null;
    const options = DV_MODE_OPTION;
    
    switch(obj.mode) {
        case options[1].value:
        dValue = options[1];
        break;
        case options[2].value:
        dValue = options[2];
        break;
        default:
        dValue = options[0];
        break;
    }  

    return(
      <TableRow key={ index }>
        <TableCellExtLink align={'left'} to={url}>{obj.name}</TableCellExtLink>
        <TableCell align={'left'}>{obj.lastSeenAt}</TableCell>
        <TableCell><span className={this.props.classes.flex}><SwitchLabels on={ obj.mode } dvId={obj.devEui} onSwitchChange={ this.onSwitchChange } /></span></TableCell>
        <TableCell><span><NativeSelects options={options} defaultValue={dValue} mode={ obj.mode } dvId={obj.devEui} onSelectChange={ this.onSelectChange } /></span></TableCell>
      </TableRow>
    );
  }

  render() {
    return(
      <Grid container spacing={24}>
        <TitleBar
          buttons={
            <Admin organizationID={this.props.match.params.organizationID}>
              <TitleBarButton
                label="Filter"
                //icon={<Plus />}
              />
            </Admin>
          }
        >
        
        </TitleBar>
        <Grid item xs={12}>
          <DataTable
            header={
              <TableRow>
                <TableCell align={'center'}>Device</TableCell>
                <TableCell align={'center'}>Status</TableCell>
                <TableCell align={'center'}>Available</TableCell>
                <TableCell className={this.props.classes.maxW} align={'center'}>Mode</TableCell>
              </TableRow>
            }
            getPage={this.getPage}
            getRow={this.getRow}
          />
        </Grid>
      </Grid>
    );
  }
}

export default withStyles(styles)(withRouter(DeviceForm));