package cmd

import (
	"os"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gitlab.com/MXCFoundation/cloud/mxprotocol-server/m2m-wallet/pkg/config"
)

// when updating this template, don't forget to update config.md!
const configTemplate = `[general]
# Log level
#
# debug=5, info=4, warning=3, error=2, fatal=1, panic=0
log_level={{ .General.LogLevel }}

# The number of times passwords must be hashed. A higher number is safer as
# an attack takes more time to perform.
password_hash_iterations={{ .General.PasswordHashIterations }}

host_server={{ .General.HostServer }}
auth_server={{ .General.AuthServer }}
auth_url={{ .General.AuthServer }}


# PostgreSQL settings.
#
# Please note that PostgreSQL 9.5+ is required.
[postgresql]
# PostgreSQL dsn (e.g.: postgres://user:password@hostname/database?sslmode=disable).
#
# Besides using an URL (e.g. 'postgres://user:password@hostname/database?sslmode=disable')
# it is also possible to use the following format:
# 'user=loraserver dbname=loraserver sslmode=disable'.
#
# The following connection parameters are supported:
#
# * dbname - The name of the database to connect to
# * user - The user to sign in as
# * password - The user's password
# * host - The host to connect to. Values that start with / are for unix domain sockets. (default is localhost)
# * port - The port to bind to. (default is 5432)
# * sslmode - Whether or not to use SSL (default is require, this is not the default for libpq)
# * fallback_application_name - An application_name to fall back to if one isn't provided.
# * connect_timeout - Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
# * sslcert - Cert file location. The file must contain PEM encoded data.
# * sslkey - Key file location. The file must contain PEM encoded data.
# * sslrootcert - The location of the root certificate file. The file must contain PEM encoded data.
#
# Valid values for sslmode are:
#
# * disable - No SSL
# * require - Always SSL (skip verification)
# * verify-ca - Always SSL (verify that the certificate presented by the server was signed by a trusted CA)
# * verify-full - Always SSL (verify that the certification presented by the server was signed by a trusted CA and the server host name matches the one in the certificate)
dsn="{{ .PostgreSQL.DSN }}"

# Automatically apply database migrations.
#
# It is possible to apply the database-migrations by hand
# (see https://github.com/brocaar/lora-app-server/tree/master/migrations)
# or let LoRa App Server migrate to the latest state automatically, by using
# this setting. Make sure that you always make a backup when upgrading Lora
# App Server and / or applying migrations.
automigrate={{ .PostgreSQL.Automigrate }}

# This is the API and web-interface exposed to the end-user.
[application_server.http_server]
# ip:port to bind the (user facing) http server to (web-interface and REST / gRPC api)
bind="{{ .ApplicationServer.HttpServer.Bind }}"

# http server TLS certificate (optional)
tls_cert="{{ .ApplicationServer.HttpServer.TLSCert }}"

# http server TLS key (optional)
tls_key="{{ .ApplicationServer.HttpServer.TLSKey }}"

# JWT secret used for api authentication / authorization
# You could generate this by executing 'openssl rand -base64 32' for example
jwt_secret="{{ .ApplicationServer.HttpServer.JWTSecret }}"

# Allow origin header (CORS).
#
# Set this to allows cross-domain communication from the browser (CORS).
# Example value: https://example.com.
# When left blank (default), CORS will not be used.
cors_allow_origin="{{ .ApplicationServer.HttpServer.CORSAllowOrigin }}"

# when set, existing users can't be re-assigned (to avoid exposure of all users to an organization admin)"
disable_assign_existing_users={{ .ApplicationServer.HttpServer.DisableAssignExistingUsers }}

[supernode]
contract_address={{ .SuperNode.ContractAddress }}
supernode_address={{ .SuperNode.SuperNodeAddress }}
api_key={{ .SuperNode.APIKey }}
check_account_seconds={{ .SuperNode.CheckAccountSeconds }}
check_payment_seconds={{ .SuperNode.CheckPaymentSecond }}

[paymentserver]
payment_service_address={{ .PaymentServer.PaymentServiceAddress }}
payment_service_port={{ .PaymentServer.PaymentServicePort }}
`

var cmdConfig = &cobra.Command{
	Use:   "configfile",
	Short: "Print mxp-server configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, config.Cstruct)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
