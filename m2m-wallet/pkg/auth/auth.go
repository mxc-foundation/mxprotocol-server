package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"net/http"
	"regexp"
)

type errStruct struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int64  `json:"code,omitempty"`
	Details []byte `json:"details,omitempty"`
}

var ctxAuth struct {
	authServer string
	authUrl    string
}

func Setup(conf config.MxpConfig) error {
	log.Info("setup auth service")

	ctxAuth.authServer = conf.General.AuthServer
	ctxAuth.authUrl = conf.General.AuthUrl
	return nil
}

func VerifyRequestViaAuthServer(ctx context.Context, requestServiceName string) (api.ProfileResponse, error) {
	log.WithField("request service", requestServiceName).Info()

	info, err := tokenMiddleware(ctx)
	if err != nil {
		return api.ProfileResponse{}, err
	}

	errInfo := errStruct{}
	err = json.Unmarshal(*info, &errInfo)
	if err != nil {
		fmt.Println(err)
	}

	if errInfo.Error != "" {
		return api.ProfileResponse{}, errors.New(errInfo.Error)
	}

	userInfo := api.ProfileResponse{}
	err = json.Unmarshal(*info, &userInfo)
	if err != nil {
		fmt.Println(err)
	}
	return userInfo, nil
}

func tokenMiddleware(ctx context.Context) (*[]byte, error) {
	tokenStr, err := getTokenFromContext(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "get token from context error")
	} else {
		res, err := getRequest(ctxAuth.authServer+ctxAuth.authUrl, tokenStr)
		if err != nil {
			return nil, errors.Wrap(err, "no response from lora app server")
		}
		return res, nil
	}
}

var validAuthorizationRegexp = regexp.MustCompile(`(?i)^bearer (.*)$`)

func getTokenFromContext(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("invalid algorithm")
	}

	token, ok := md["authorization"]
	if !ok || len(token) == 0 {
		return "", errors.New("no authorization-data in metadata")
	}

	match := validAuthorizationRegexp.FindStringSubmatch(token[0])

	// authorization header should respect RFC1945
	if len(match) == 0 {
		log.Warning("Deprecated Authorization header : RFC1945 format expected : Authorization: <type> <credentials>")
		return token[0], nil
	}

	return match[1], nil
}

//send get request to lora app server
func getRequest(url, jwtToken string) (*[]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	authStr := "Bearer " + jwtToken
	req.Header.Add("Authorization", authStr)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	return &body, nil
}

type InternalServerAPI struct {
	serviceName string
}

func NewInternalServerAPI() *InternalServerAPI {
	return &InternalServerAPI{serviceName: "internal get jwt"}
}

func (s *InternalServerAPI) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	requestBody, err := json.Marshal(map[string]string{
		"password": req.Password,
		"username": req.Username,
	})

	if err != nil {
		log.Warn(err)
		return &api.LoginResponse{}, err
	}

	request, err := http.NewRequest("POST", ctxAuth.authServer+"/api/internal/login", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	if err != nil {
		log.Warn(err)
		return &api.LoginResponse{}, err
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return &api.LoginResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &api.LoginResponse{}, err
	}

	// parse response
	errInfo := errStruct{}
	err = json.Unmarshal(body, &errInfo)
	if err != nil {
		fmt.Println("unmarshal err", err)
	}

	if errInfo.Error != "" {
		return &api.LoginResponse{}, err
	}

	var output map[string]string
	err = json.Unmarshal(body, &output)
	if err != nil {
		fmt.Println("unmarshal response", err)
	}
	return &api.LoginResponse{Jwt: output["jwt"]}, nil
}
