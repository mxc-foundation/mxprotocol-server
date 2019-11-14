import React, { Component } from "react";

import Grid from "@material-ui/core/Grid";
import { withRouter } from "react-router-dom";
import { withStyles } from "@material-ui/core/styles";

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

class SuperNodeNetworkActivityHistory extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return(
            <Grid container spacing={24}>
                <Grid item xs={12}>
                    <a href={"https://www.google.de"} target="_blank">
                        Coming soon, click here for more information.
                    </a>
                </Grid>
            </Grid>
        );
    }
}

export default withStyles(styles)(withRouter(SuperNodeNetworkActivityHistory));