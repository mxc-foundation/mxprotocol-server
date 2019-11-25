import theme from "../theme";

const SideNavStyle = {
    drawerPaper: {
      position: "fixed",
      width: 270,
      paddingTop: 94,
      paddingLeft: 0,
      paddingRight: 0,
      backgroundColor: theme.palette.secondary.secondary,
      color: theme.palette.textPrimary.main,
      boxShadow: '1px 1px 5px 0px rgba(29, 30, 33, 0.5)',
    },
    select: {
      paddingTop: theme.spacing.unit,
      paddingLeft: theme.spacing.unit * 3,
      paddingRight: theme.spacing.unit * 3,
      paddingBottom: theme.spacing.unit * 1,
    },
    selected: {
      color: theme.palette.common.white,
    },
    card: { // LPWAN Server options
      width: '100%',
      height: 250,
      position: 'absolute',
      bottom: 5,
      color: theme.palette.textPrimary.main,
      fontSize: 14,
      paddingLeft: 0,
      paddingRight: 0,
    },
    static: {
      position: 'static'
    },
    iconStyle: {
      color: theme.palette.common.white,
    },
    divider: {
      padding: 0,
      color: theme.palette.darkBG.main,
      width: '80%',
    },
  };
  
export default SideNavStyle;
