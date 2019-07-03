package auth

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Err struct {
	Error   string `json:"error,omitempty"`
	Message string    `json:"message,omitempty"`
	Code    int64 `json:"code,omitempty"`
	Details []byte   `json:"details,omitempty"`
}

type ProfileResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Organizations to which the user is associated.
	Organizations []*OrganizationLink `protobuf:"bytes,3,rep,name=organizations,proto3" json:"organizations,omitempty"`
	// Profile settings.
	Settings *ProfileSettings `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings,omitempty"`
}

type User struct {
	// User ID.
	// Will be set automatically on create.
	Id string `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Username of the user.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// The session timeout, in minutes.
	SessionTtl int32 `protobuf:"varint,3,opt,name=session_ttl,json=sessionTTL,proto3" json:"session_ttl,omitempty"`
	// Set to true to make the user a global administrator.
	IsAdmin bool `protobuf:"varint,4,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// Set to false to disable the user.
	IsActive bool `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// E-mail of the user.
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	// Optional note to store with the user.
	Note string `protobuf:"bytes,7,opt,name=note,proto3" json:"note,omitempty"`
}

type OrganizationLink struct {
	// Organization ID.
	OrganizationId int64 `protobuf:"varint,1,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Organization name.
	OrganizationName string `protobuf:"bytes,2,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty"`
	// User is admin within the context of this organization.
	IsAdmin bool `protobuf:"varint,3,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

type ProfileSettings struct {
	// Existing users in the system can not be assigned to organizations and
	// application and can not be listed by non global admin users.
	DisableAssignExistingUsers bool `protobuf:"varint,1,opt,name=disable_assign_existing_users,json=disableAssignExistingUsers,proto3" json:"disable_assign_existing_users,omitempty"`
}

var validAuthorizationRegexp = regexp.MustCompile(`(?i)^bearer (.*)$`)

func TokenMiddleware(ctx context.Context, url string) (*[]byte, error) {
	tokenStr, err := getTokenFromContext(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "get token from context error")
	} else {
		/*//only if we need to verify jwt in M2M!!!
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("not authorization")
			}
			return []byte("secret"), nil
		})
		if !token.Valid {
		} else {
		}
		*/

		res, err := getRequest(url, tokenStr)
		if err != nil {
			return nil, errors.Wrap(err, "no response from lora app server")
		}
		return res, nil
	}
}

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