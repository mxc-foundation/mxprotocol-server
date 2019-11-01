import theme from "../../theme";

const SuperNodeEthStyle = {
    tabs: {
        borderBottom: "1px solid " + theme.palette.divider,
        height: "49px",
      },
      navText: {
        fontSize: 14,
      },
      TitleBar: {
        height: 115,
        width: '50%',
        light: true,
        display: 'flex',
        flexDirection: 'column'
      },
      card: {
        minWidth: 180,
        width: 220,
        backgroundColor: "#0C0270",
      },
      divider: {
        padding: 0,
        color: '#FFFFFF',
        width: '100%',
      },
      padding: {
        padding: 0,
      },
      column: {
        display: 'flex',
        flexDirection: 'column',
      },
      link: {
        textDecoration: "none",
        fontWeight: "bold",
        fontSize: 12,
        color: theme.palette.textSecondary.main,
        opacity: 0.7,
          "&:hover": {
            opacity: 1,
          }
      },
  };
  
export default SuperNodeEthStyle;
