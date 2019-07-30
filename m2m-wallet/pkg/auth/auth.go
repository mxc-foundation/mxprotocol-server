package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/api"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

type resCode int32

const (
	OK                       resCode = 0
	OrganizationIdRearranged resCode = 1
	JsonParseError           resCode = 2
	AuthFailed               resCode = 3
	OrganizationIdMisMatch   resCode = 4
)

type VerifyResult struct {
	Err  error
	Type resCode
}

type User struct {
	// User ID.
	// Will be set automatically on create.
	Id string `json:"id,omitempty"`
	// Username of the user.
	Username string `json:"username,omitempty"`
	// The session timeout, in minutes.
	SessionTtl int32 `json:"sessionTTL,omitempty"`
	// Set to true to make the user a global administrator.
	IsAdmin bool `json:"isAdmin,omitempty"`
	// Set to false to disable the user.
	IsActive bool `json:"isActive,omitempty"`
	// E-mail of the user.
	Email string `json:"email,omitempty"`
	Note  string `json:"note,omitempty"`
}

type Settings struct {
	DisableAssignExistingUsers bool `json:"disableAssignExistingUsers,omitempty"`
}

type OrganizationLink struct {
	// Organization ID.
	OrganizationId string `json:"organizationID,omitempty"`
	// Organization name.
	OrganizationName string `json:"organizationName,omitempty"`
	// User is admin within the context of this organization.
	IsAdmin bool `json:"isAdmin,omitempty"`
	// Created at timestamp.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type ProfileResponse struct {
	// User object.
	User User `json:"user,omitempty"`
	// Profile settings.
	Settings Settings `json:"settings,omitempty"`
	// Organizations to which the user is associated.
	Organizations []OrganizationLink `json:"organizations,omitempty"`
}

func VerifyRequestViaAuthServer(ctx context.Context, requestServiceName string, reqOrgId int64) (api.ProfileResponse, VerifyResult) {
	// parse jwt from context
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		log.WithError(err).Error("auth/VerifyRequestViaAuthServer")
		return api.ProfileResponse{}, VerifyResult{err, JsonParseError}
	}

	// request user profile with jwt
	userProfile, err := requestUserProfileWithJWT(tokenStr)
	if err != nil {
		log.WithError(err).Error("auth/VerifyRequestViaAuthServer")
		return api.ProfileResponse{}, VerifyResult{err, AuthFailed}
	}

	return isOrgListRearranged(userProfile, reqOrgId)
}

func isOrgListRearranged(userProfile ProfileResponse, orgId int64) (api.ProfileResponse, VerifyResult) {
	profile := api.ProfileResponse{}
	profile.User = &api.User{}
	profile.Settings = &api.ProfileSettings{}

	profile.User.Id = userProfile.User.Id
	profile.User.Username = userProfile.User.Username
	profile.User.SessionTtl = userProfile.User.SessionTtl
	profile.User.IsAdmin = userProfile.User.IsAdmin
	profile.User.IsActive = userProfile.User.IsActive
	profile.User.Email = userProfile.User.Email
	profile.User.Note = userProfile.User.Note

	profile.Settings.DisableAssignExistingUsers = userProfile.Settings.DisableAssignExistingUsers

	orgDeleted := true
	for _, v := range userProfile.Organizations {
		id, _ := strconv.ParseInt(v.OrganizationId, 10, 64)
		org := api.OrganizationLink{}
		org.OrganizationId = id
		org.IsAdmin = v.IsAdmin
		org.OrganizationName = v.OrganizationName
		org.CreatedAt = &timestamp.Timestamp{Seconds: int64(v.CreatedAt.Second()), Nanos: int32(v.CreatedAt.Nanosecond())}
		org.UpdatedAt = &timestamp.Timestamp{Seconds: int64(v.UpdatedAt.Second()), Nanos: int32(v.UpdatedAt.Nanosecond())}
		profile.Organizations = append(profile.Organizations, &org)

		if id == orgId {
			orgDeleted = false
		}
	}


	if orgId == 0 {
		if profile.User.IsAdmin == true {
			return profile, VerifyResult{nil, OK}
		} else {
			return profile, VerifyResult{nil, OrganizationIdMisMatch}
		}
	}

	if orgDeleted {
		return profile, VerifyResult{nil, OrganizationIdRearranged}
	}

	return profile, VerifyResult{nil, OK}
}

func requestUserProfileWithJWT(jwToken string) (ProfileResponse, error) {
	res, err := getRequest(ctxAuth.authServer+ctxAuth.authUrl, jwToken)
	if err != nil {
		log.WithError(err).Error("auth/requestUserProfileWithJWT")
		return ProfileResponse{}, err
	}

	errInfo := errStruct{}
	err = json.Unmarshal(*res, &errInfo)
	if err != nil {
		log.WithError(err).Error("auth/requestUserProfileWithJWT")
		return ProfileResponse{}, err
	}

	if errInfo.Error != "" {
		err := errors.New(errInfo.Error)
		log.WithError(err).Error("auth/requestUserProfileWithJWT")
		return ProfileResponse{}, err
	}

	userProfile := ProfileResponse{}
	err = json.Unmarshal(*res, &userProfile)
	if err != nil {
		log.WithError(err).Error("auth/requestUserProfileWithJWT")
		return ProfileResponse{}, err
	}
	return userProfile, nil
}

var validAuthorizationRegexp = regexp.MustCompile(`(?i)^bearer (.*)$`)

//get jwt token from ctx
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
	log.WithField("body", string(body)).Info("getRequest response")
	return &body, nil
}

type InternalServerAPI struct {
	serviceName string
}

func NewInternalServerAPI() *InternalServerAPI {
	return &InternalServerAPI{serviceName: "internal get jwt"}
}

func requestJWTWithUsernamePass(username string, password string) (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"password": password,
		"username": username,
	})

	if err != nil {
		log.Warn(err)
		return "", err
	}

	request, err := http.NewRequest("POST", ctxAuth.authServer+"/api/internal/login", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	if err != nil {
		log.Warn(err)
		return "", err
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// parse response
	errInfo := errStruct{}
	err = json.Unmarshal(body, &errInfo)
	if err != nil {
		fmt.Println("unmarshal err", err)
	}

	if errInfo.Error != "" {
		return "", err
	}

	var output map[string]string
	err = json.Unmarshal(body, &output)
	if err != nil {
		fmt.Println("unmarshal response", err)
	}
	return output["jwt"], nil
}

func (s *InternalServerAPI) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	jwt, err := requestJWTWithUsernamePass(req.Username, req.Password)
	if err != nil {
		return &api.LoginResponse{}, err
	}
	return &api.LoginResponse{Jwt: jwt}, nil
}

func (s *InternalServerAPI) GetUserOrganizationList(ctx context.Context, req *api.GetUserOrganizationListRequest) (*api.GetUserOrganizationListResponse, error) {
	userProfile, res := VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case AuthFailed:
		fallthrough
	case JsonParseError:
		fallthrough
	case OrganizationIdMisMatch:

		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case OrganizationIdRearranged:
		fallthrough
	case OK:
		orgList := api.GetUserOrganizationListResponse{}
		orgList.Organizations = userProfile.Organizations

		if userProfile.User.IsAdmin == true {
			org := api.OrganizationLink{
				OrganizationId:   0,
				OrganizationName: "Super_admin",
				IsAdmin:          true,
			}
			orgList.Organizations = append(orgList.Organizations, &org)
		}

		return &orgList, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}
