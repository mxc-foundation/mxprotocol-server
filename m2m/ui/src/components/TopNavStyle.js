import blue from "@material-ui/core/colors/blue";
import theme from "../theme";

const TopNavStyle = {
    appBar: {
        zIndex: theme.zIndex.drawer + 1,
        /*backgroundColor: '#09006E',*/
        backgroundColor: theme.palette.secondary.main,
        height: 84
    },
    menuButton: {
        marginLeft: -12,
        marginRight: 10,
    },
    hidden: {
        display: "none",
    },
    flex: {
        flex: 1,
        paddingLeft: 40,
    },
    brandLogo: {
        height: 32,
        marginLeft: 20
    },
    logo: {
        height: 32,
        marginLeft: -45,
    },
    search: {
        marginRight: 3 * theme.spacing.unit,
        color: theme.palette.common.white,
        background: blue[400],
        width: 450,
        padding: 5,
        borderRadius: 3,
    },
    avatar: {
        background: theme.palette.secondary.main,
        color: theme.palette.secondary.main,
    },
    chip: {
        background: theme.palette.secondary.main,
        color: theme.palette.common.white,
        marginRight: theme.spacing.unit,
        "&:hover": {
          background: theme.palette.primary.secondary,
        },
        "&:active": {
          background: theme.palette.primary.main,
        },
        "&:visited": {
            background: theme.palette.primary.main,
        },
      },
    iconButton: {
        color: theme.palette.common.white,
        marginRight: theme.spacing.unit,
    },
    iconStyle: {
        color: theme.palette.primary.main,
    },
    noPadding: {
        padding: 0,
        color: theme.palette.textPrimary.main
    }
  };
  
export default TopNavStyle;
