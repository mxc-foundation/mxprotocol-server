import { createMuiTheme } from "@material-ui/core/styles";
import { border } from "@material-ui/system";
//import { teal } from "@material-ui/core/colors";

const tealHighLight = '#00FFD9';
const tealHighLight20 = '#00FFD920';
const blueMxcBrand = '#09006E';
const blueBG = '#070033';
const overlayBG = '#0C027060';
const white = '#F9FAFC';
const linkTextColor = '#BBE9E8';

const theme = createMuiTheme({
    palette: {
      primary: { main: tealHighLight, secondary: tealHighLight20 }, 
      secondary: { main: blueMxcBrand, secondary: overlayBG }, 
      darkBG: { main: blueBG }, 
      textPrimary: {main: white}, 
      textSecondary: {main: linkTextColor} 
    },
    MuiListItemIcon: {
      root: {
        color: white
      }
    },
    //tab 
    MuiTypography: {
      root: {
        color: white,
      },
      body1: {
        color: white,
      },
      colorTextSecondary: {
        color: white,
      },
    },
    typography: {
      //useNextVariants: true,
      subheading: {
        color: white,
        "&:hover": {
          color: 'white',
        },
      },
      title: {
        color: white
      },
      fontFamily: [
        'Montserrat',
      ].join(','),
    },
    overrides: {
      MuiTypography: {
        root: {
          color: white,
        },
        body1: {
          color: white,
          fontSize: '0.8rem'
        },
        body2: {
          color: white,
          fontSize: '0.7rem'
        },
        colorTextSecondary: {
          color: white,
        },
        headline: {
          color: white
        },
        caption: {
          color: white
        },
      },
      MuiInput: {
        root: {
          color: white
        },
        underline: {
          "&:before": {
            borderBottom: `1px solid #F9FAFC`
          },
          "&:hover": {
            borderBottom: `1px solid #00FFD9`
          }
        },
      },
      MuiAppBar: {
        root: {
          //width: '1024px',
          color: white
        },
        positionFixed: {
          left: 'inherit',
          right: 'inherit'
        }
      },
      MuiSelect: {
        icon: {
          color: white,
          right: 0,
          position: 'absolute',
          pointerEvents: 'none',
        }
      },
      MuiIconButton: {
        root: {
          color: white,
        }
      },
/*       MuiInputBase: {
        input: {
          color: '#F9FAFC',
          fontWeight: "bolder",
          "&:-webkit-autofill": {
            WebkitBoxShadow: "0 0 0 1000px #F9FAFC inset"
          }
        }
      }, */
      MuiDivider: {
        root: {
          backgroundColor: '#00000040',
          margin: '5px 0px 5px 0px',
        },
        light: {
          backgroundColor: '#FFFFFF50',
        }
      },
      MuiTable: {
        root: {
          background: 'transparent',
          //minWidth: 840,
        }
      },
      MuiTableCell: {
        head: {
          color: white,
          fontWeight: '800',
          fontSize: '1em',
          padding: 10, 
        },
        body: {
          background: 'none',
          color: white,
          //maxWidth: 140,
          whiteSpace: 'nowrap', 
          //overflow: 'hidden',
          textOverflow: 'ellipsis',
          fontWeight: '400', 
        },
        root: {
          padding: '4px 5px',
          //maxWidth: 140,
          whiteSpace: 'nowrap', 
          //overflow: 'hidden',
          textOverflow: 'ellipsis',
          borderBottom: 'solid 1px #070033',
          lineHeight: '40px',
          textAlign: 'left',
        }
      },
      MuiPaper: {
        root: {
          backgroundColor: overlayBG,
          padding: 10,
        }
      },
      MuiTablePagination: {
        root: {
          color: white,
          background: 'none',
        }
      },
      MuiButton: { 
        root: {
          background: tealHighLight,
          color: blueMxcBrand,
          width: 160,
          height: 50,
          fontWeight: 'bolder',
          marginRight: 5,
          boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2)',
          "&:hover": {
            backgroundColor: "#00CCAE",
          },
        },
        outlined: {
          backgroundColor: 'transparent',
          color: tealHighLight,
          //padding: 30,
          fontWeight: 900,
          lineHeight: 1.5,
          borderWidth: 2,
          borderColor: white,
          "&:hover": {
            backgroundColor: tealHighLight20,
            borderColor: "#00CCAE",
            color: white,
          },
        },
/*         link: {
          color: tealHighLight,
          //padding: 30,
          fontWeight: 900,
          lineHeight: 1.5,
          "&:hover": {
            color: "#00CCAE",
          },
        }, */
        label: {
          color: blueMxcBrand
        },
        text: { 
          color: white, 
          padding: 6,
        },
      },
      MuiFormLabel: { 
        root: { 
          color: white, 
        },
      },
      MuiFormHelperText: { 
        root: { 
          color: white, 
        },
      },
      MuiPrivateTabScrollButton:{
        root: {
          width: 0
        }
      },
      MuiTab: {
        root: {
          color: white,
        },
        textColorPrimary: {
          color: white
        },
      },
      MuiSvgIcon: {
        root: {
          fill: '#F9FAFC80',
        },
      },
      MuiDialog: {
        color: white,
        root: {
          color: white,
          boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2)',
        },
      },
      MuiMenu: {
        paper: {
          backgroundColor: blueBG,
          marginTop: '50px',
          color: white
        }
      }
    },
});
  
export default theme;