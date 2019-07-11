import { createMuiTheme } from "@material-ui/core/styles";

const tealHighLight = '#00FFD9';
const blueMxcBrand = '#09006E';
const white = '#F9FAFC';
const linkTextColor = '#CAFCF5';

const theme = createMuiTheme({
    palette: {
      primary: { main: tealHighLight }, 
      secondary: { main: blueMxcBrand }, 
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
    },
    typography: {
      //useNextVariants: true,
      subheading: {
        color: white
      },
      title: {
        color: white
      },
      fontFamily: [
        'Montserrat',
      ].join(','),
    },
    overrides: {
      MuiInput: {
        root: {
          color: white
        },
        underline: {
          "&:before": {
            borderBottom: `1px solid #F9FAFC`
          }
        }
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
      MuiInputBase: {
        input: {
          color: white,
          "&:-webkit-autofill": {
            WebkitBoxShadow: "0 0 0 1000px #F9FAFC inset"
          }
        }
      },
      MuiTable: {
        root: {
          background: white,
        }
      },
      MuiDivider: {
        light: {
          backgroundColor: '#FFFFFF50',
        },
        /* dark: {
          backgroundColor: '#1C147870',
          padding: 500, 
        } */
      },
      MuiTableCell: {
        head: {
          background: '#0C0270',
          color: white,
          fontWeight: 'bold',
          padding: 10, 
        },
        body: {
          background: '#0C0270',
          color: white,
        }
      },
      MuiPaper: {
        root: {
          backgroundColor: '#0C0270',
          padding: 10,
        }
      },
      MuiTablePagination: {
        root: {
          color: white,
          background: '#0C0270',
        }
      },
      MuiButton: { 
        root: {
          backgroundColor: '#00FFD9',
          width: 135,
          height: 50,
          fontWeight: 'bolder',
          marginRight: 5,
          boxShadow: '0 4px 8px 0 rgba(0, 0, 0, 0.2)',
          backgroundColor: '#00FFD9',
            "&:hover": {
              backgroundColor: '#00CCAE',
            }
        },
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
        }
      },
      MuiSvgIcon: {
        root: {
          fill: '#F9FAFC',
        },
      },
    },
});
  
export default theme;