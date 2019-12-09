import dispatcher from "../dispatcher";
import history from '../history';


export function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response
  } else {
    throw response.json();
  }
};

export function errorHandler(error) {
  if(error.response === undefined) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "error",
        message: error.message,
      },
    });
  } else {
    if (error.response.obj.code === 16) {
      onLoginFailed();
    } else {
      dispatcher.dispatch({
        type: "CREATE_NOTIFICATION",
        notification: {
          type: "error",
          message: error.response.obj.error + " (code: " + error.response.obj.code + ")",
        },
      });
    }
  }
};

export function onLoginFailed() {
  console.error('login failed');
  history.push("/login");
}

export function errorHandlerLogin(error) {
  if(error.response === undefined) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "error",
        message: error.message,
      },
    });
  } else {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "error",
        message: error.response.obj.error + " (code: " + error.response.obj.code + ")",
      },
    });
  }
};

export function errorHandlerIgnoreNotFound(error) {
  if (error.response === undefined) {
    dispatcher.dispatch({
      type: "CREATE_NOTIFICATION",
      notification: {
        type: "error",
        message: error.message,
      },
    });
  } else {
    if (error.response.obj.code === 16 && history.location.pathname !== "/login") {
      onLoginFailed();
    } else if (error.response.obj.code !== 5) {
      dispatcher.dispatch({
        type: "CREATE_NOTIFICATION",
        notification: {
          type: "error",
          message: error.response.obj.error + " (code: " + error.response.obj.code + ")",
        },
      });
    }
  }
};

export function errorHandlerIgnoreNotFoundWithCallback(callbackFunc) {
  return function(error) {
    if (error.response.obj.code === 5) {
      callbackFunc(null);
    } else {
      errorHandlerIgnoreNotFound(error);
    }
  }
}
