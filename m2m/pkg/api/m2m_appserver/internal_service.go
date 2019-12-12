package appserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	api "gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/api/appserver"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func requestJWTWithUsernamePass(username string, password string) (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"password": password,
		"username": username,
	})

	if err != nil {
		log.Warn(err)
		return "", err
	}

	request, err := http.NewRequest("POST", auth.AuthServer+"/api/internal/login", bytes.NewBuffer(requestBody))
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
	errInfo := auth.ErrStruct{}
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

func (s *M2MServerAPI) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	jwt, err := requestJWTWithUsernamePass(req.Username, req.Password)
	if err != nil {
		return &api.LoginResponse{}, err
	}

	tokenStr, err := auth.GetTokenFromContext(ctx)
	if err != nil {
		if jwt != tokenStr {
			return &api.LoginResponse{}, err
		}
	}

	return &api.LoginResponse{Jwt: jwt}, nil
}

func (s *M2MServerAPI) GetUserOrganizationList(ctx context.Context, req *api.GetUserOrganizationListRequest) (*api.GetUserOrganizationListResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:

		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		fallthrough
	case auth.OK:
		orgList := api.GetUserOrganizationListResponse{}

		// users who are not super admin users
		orgList.Organizations = userProfile.Organizations

		/*		if userProfile.User.IsAdmin == true {
				org := api.OrganizationLink{
					OrganizationId:   0,
					OrganizationName: "Super_admin",
					IsAdmin:          true,
				}
				orgList.Organizations = append(orgList.Organizations, &org)
			}*/

		return &orgList, nil
	}

	return nil, status.Errorf(codes.Unknown, "")
}

func (s *M2MServerAPI) GetUserProfile(ctx context.Context, req *api.GetUserProfileRequest) (*api.GetUserProfileResponse, error) {
	userProfile, res := auth.VerifyRequestViaAuthServer(ctx, s.serviceName, req.OrgId)

	switch res.Type {
	case auth.AuthFailed:
		fallthrough
	case auth.JsonParseError:
		fallthrough
	case auth.OrganizationIdMisMatch:

		return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %s", res.Err)

	case auth.OrganizationIdRearranged:
		fallthrough
	case auth.OK:
		return &api.GetUserProfileResponse{UserProfile: &userProfile}, nil
	}
	return nil, status.Errorf(codes.Unknown, "")
}
