import theme from "../theme";

const SideNavStyle = {
    drawerPaper: {
      position: "fixed",
      width: 270,
      paddingTop: theme.spacing.unit * 10,
      paddingLeft: 0,
      paddingRight: 0,
      backgroundColor: theme.palette.secondary.secondary,
      color: theme.palette.textPrimary.main,
      fontSize: 'bold',
      boxShadow: '1px 1px 5px 0px rgba(29, 30, 33, 0.5)',
    },
    select: {
      paddingTop: theme.spacing.unit,
      paddingLeft: theme.spacing.unit * 3,
      paddingRight: theme.spacing.unit * 3,
      paddingBottom: theme.spacing.unit * 1,
    },
    selected: {
      fontSize: 'bold', 
      color: theme.palette.common.white,
    },
    card: { // lora server options
      width: '100%',
      height: 250,
      position: 'absolute',
      bottom: 5,
      backgroundColor: '#09006E20',
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
      color: '#1C1478',
      width: '100%',
    },
  };
  
export default SideNavStyle;
