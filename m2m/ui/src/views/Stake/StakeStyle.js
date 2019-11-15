import theme from "../../theme";

const DeviceStyles = {
  title: {
    color: '#00FFD9',
    fontSize: 14,
    padding: 6,
  },
  pos: {
    marginBottom: 12,
    color: '#FFFFFF',
    textAlign: 'right',
  },
  between: {
    display: 'flex',
    justifyContent:'spaceBetween'
  },
  flex: {
    display: 'flex',
  },
  navText: {
    fontSize: '0.85rem !important',
  },
  TitleBar: {
    width: '100%',
    light: true,
    display: 'flex',
    flexDirection: 'column',
    padding: '0px 0px 50px 0px' 
  },
  divider: {
    padding: 0,
    color: '#FFFFFF',
    width: '100%',
  },
  padding: {
    padding: 0,
  },
  link: {
    textDecoration: "none",
    fontWeight: "bold",
    fontSize: '1rem',
    color: theme.palette.textSecondary.main,
    opacity: 0.7,
      "&:hover": {
        opacity: 1,
      }
  },
  subTitle2:{
    textDecoration: "none",
    padding: 0,
    fontWeight: "bold",
    fontSize: 12,
    color: theme.palette.textPrimary.main,
    cursor: "pointer",
    opacity: 0.7,
      "&:hover": {
        opacity: 1,
      } 
  },
  infoBox: {
    height: 150,
    width: '100%',
    borderWidth: 2,
    borderLeft: '5px solid #4d89e5',
    backgroundColor: '#142257',
    padding: 10,
    margin: 10
  },
  infoBoxSucceed: {
    height: 150,
    width: '100%',
    borderWidth: 2,
    borderLeft: '5px solid #60af4e',
    backgroundColor: '#142257',
    padding: 10,
    margin: 10
  },
  infoBoxError: {
    height: 150,
    width: '100%',
    borderWidth: 2,
    borderLeft: '5px solid #df434e',
    backgroundColor: '#142257',
    padding: 10,
    margin: 10
  },
  pRight: {
    paddingRight: 10
  },
  pLeft: {
    paddingLeft: 10
  },
  urStake: {
    width: '100%',
    height: 60,
    backgroundColor: '#142257',
    padding: 10,
    margin: 10
  },
  mxc: {
    fontWeight: 'bolder'
  }       
};
  
export default DeviceStyles;