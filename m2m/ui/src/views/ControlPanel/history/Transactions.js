import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";
import TableCell from "@material-ui/core/TableCell";
import theme from "../../../theme";

const styles = {
    maxW140: {
        maxWidth: 140,
        //backgroundColor: "#0C0270",
        whiteSpace: 'nowrap',
        overflow: 'hidden',
        textOverflow: 'ellipsis'
    },
    flex:{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center'

    }
};

class SuperNodeTransactions extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return(
            <Grid container spacing={24}>
                <Grid item xs={6}>
                    <TableCell align={this.props.align}>
                    <span style={
                        {
                            textDecoration: "none",
                            color: theme.palette.primary.main,
                            cursor: "pointer",
                            padding: 0,
                            fontWeight: "bold",
                            fontSize: 20,
                            opacity: 0.7,
                            "&:hover": {
                                opacity: 1,
                            }
                        }
                    } className={this.props.classes.link} onClick={()=>{
                        window.open("www.mxc.org", '_blank');
                    }}>Coming soon, more information about transaction history, please click here</span>
                    </TableCell>
                </Grid>
            </Grid>
        );
    }
}

export default withStyles(styles)(withRouter(SuperNodeTransactions));