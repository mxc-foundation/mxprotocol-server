import React, { Component } from "react";
import { Link, withRouter } from "react-router-dom";

import Grid from "@material-ui/core/Grid";
import TitleBarTitle from "../../../components/TitleBarTitle";
import { withStyles } from "@material-ui/core/styles";
import styles from "../../withdraw/WithdrawStyle"
import Button from "@material-ui/core/Button";
import theme from "../../../theme";
import TableCell from "@material-ui/core/TableCell";


class SuperAdminWithdraw extends Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    render() {
        return (
            <Grid container spacing={24} className={this.props.classes.backgroundColor}>
                {/*<Grid item xs={12} className={this.props.classes.divider}>
                <div className={this.props.classes.TitleBar}>
                    <TitleBarTitle title="Withdraw" />
                </div>
            </Grid>*/}

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
                    } className={this.props.classes.link} >
                        Coming soon...
                    </span>
                </TableCell>
            </Grid>

            </Grid>
        );
    }
}

export default withStyles(styles)(withRouter(SuperAdminWithdraw));