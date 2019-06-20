package cmd

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"mxprotocol-server/m2m-wallet/pkg/config"
)

var cfgFile string
var version string

var cmdMXP = &cobra.Command{
	Use:   "mxp-server",
	Short: "MXProtocol server",
	RunE:  run,
}

func init() {
	// after main function start running, initializers registered here will be called one after another
	cobra.OnInitialize(initConfig)

	// settings before the main function starts
	cmdMXP.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "path to configuration file (optional)")
	cmdMXP.PersistentFlags().Int("log-level", 4, "debug=5, info=4, error=2, fatal=1, panic=0")

	// bind flag to config vars
	viper.BindPFlag("general.log_level", cmdMXP.PersistentFlags().Lookup("log-level"))

	// defaults
	viper.SetDefault("general.password_hash_iterations", 100000)
	viper.SetDefault("general.host_server", "mxp-server-domain")

	viper.SetDefault("postgresql.dsn", "postgres://localhost/mxp-server?sslmode=disable")
	viper.SetDefault("postgresql.automigrate", true)

	viper.SetDefault("application_server.http_server.bind", "0.0.0.0:5420")
	viper.SetDefault("application_server.http_server.tls_cert", "")
	viper.SetDefault("application_server.http_server.tls_key", "")
	viper.SetDefault("application_server.http_server.jwt_secret", "DOE1KiNzpQ82elRQ9HMWyxmADQ5f2B2TBAgOjL7ZZWA=")
	viper.SetDefault("application_server.http_server.cors_allow_origin", "http://localhost:5420")
	viper.SetDefault("application_server.http_server.disable_assign_existing_users", false)

	cmdMXP.AddCommand(cmdConfig)
	cmdMXP.AddCommand(cmdVersion)
}

func Execute(v string) {
	version = v
	if err := cmdMXP.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	if cfgFile != "" {
		b, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("error loading config file")
		}
		viper.SetConfigType("toml")
		if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("error loading config file")
		}
	} else {
		viper.SetConfigName("mxp-server")
		// search in order: "."  "$HOME/.config/mxp-server"  "/etc/mxp-server"
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/mxp-server")
		viper.AddConfigPath("/etc/mxp-server")
		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				log.Warning("No configuration file found, using defaults.")
			default:
				log.WithError(err).Fatal("read configuration file error")
			}
		}
	}

	viperBindEnvs(config.Cstruct)

	if err := viper.Unmarshal(&config.Cstruct); err != nil {
		log.WithError(err).Fatal("unmarshal config error")
	}
}

func viperBindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			tv = strings.ToLower(t.Name)
		}
		if tv == "-" {
			continue
		}

		switch v.Kind() {
		case reflect.Struct:
			viperBindEnvs(v.Interface(), append(parts, tv)...)
		default:
			key := strings.Join(append(parts, tv), ".")
			viper.BindEnv(key)
		}
	}
}
