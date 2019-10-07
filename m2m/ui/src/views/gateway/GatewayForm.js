import React, { Component } from "react";
import { withRouter } from 'react-router-dom';
import { withStyles } from "@material-ui/core/styles";

import Grid from "@material-ui/core/Grid";
import TableCell from "@material-ui/core/TableCell";
import TableCellExtLink from "../../components/TableCellExtLink";
import TableRow from "@material-ui/core/TableRow";
import { GW_MODE_OPTION, GW_INACTIVE } from "../../util/Data"

import GatewayStore from "../../stores/GatewayStore.js";
import TitleBar from "../../components/TitleBar";
import NativeSelects from "../../components/NativeSelects";
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

class GatewayForm extends Component {
  constructor(props) {
    super(props);
    this.getPage = this.getPage.bind(this);
    this.getRow = this.getRow.bind(this);
  }

  getPage(limit, offset, callbackFunc) {
    GatewayStore.getGatewayList(this.props.match.params.organizationID, offset, limit, data => {
        callbackFunc({
            totalCount: parseInt(data.count),
            result: data.gwProfile
          });
      }); 
  }

  onSelectChange = (e) => {
    const gatewayInfo = {
        gwId: e.gwId, 
        gwMode: e.target.value
    }
    
    this.props.onSelectChange(gatewayInfo);
  }

  onSwitchChange(e) {
    //console.log('e.target.value', e.target.value);
  }

  getRow(obj, index) {
    const url = `/#/organizations/${this.props.match.params.organizationID}/gateways/${obj.mac}`;
    let dValue = null;
    const options = GW_MODE_OPTION;
    
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
    
    let on = (obj.mode !== GW_INACTIVE) ? true : false;
    
    return(
      <TableRow key={index}>
        <TableCellExtLink align={'left'} for={'lora'} to={url}>{obj.name}</TableCellExtLink>
        <TableCell align={'left'}>{obj.lastSeenAt.substring(0, 19)}</TableCell>
        <TableCell align={'right'}>{this.props.downlinkFee}</TableCell>
        <TableCell align={'left'}><NativeSelects options={options} defaultValue={dValue} default={ obj.mod } gwId={obj.id} onSelectChange={ this.onSelectChange } /></TableCell>
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
                <TableCell align={'left'}>Gateway</TableCell>
                <TableCell align={'left'}>Status</TableCell>
                <TableCell align={'right'}>Downlink Price</TableCell>
                <TableCell align={'left'}>Mode</TableCell>
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

export default withStyles(styles)(withRouter(GatewayForm));
