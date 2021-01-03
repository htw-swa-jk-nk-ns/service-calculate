package cmd

import (
	"github.com/htw-swa-jk-nk-ns/service-calculate/api"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func init() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})

	cobra.OnInitialize(func() {
		rep := strings.NewReplacer(".", "_")
		viper.SetEnvKeyReplacer(rep)

		viper.SetEnvPrefix("SERVICE_CALCULATE")
		viper.AutomaticEnv()

		_ = viper.ReadInConfig()
	})

	//api
	rootCMD.PersistentFlags().String("api-format", "json", "json format ('json' or 'xml')")
	rootCMD.PersistentFlags().Int("api-port", 8889, "api port")

	//api
	err := viper.BindPFlag("api.format", rootCMD.PersistentFlags().Lookup("api-format"))
	if err != nil {
		log.Error().
			AnErr("Error", err).
			Msg("Can't bind flag api format")
		return
	}

	err = viper.BindPFlag("api.port", rootCMD.PersistentFlags().Lookup("api-port"))
	if err != nil {
		log.Error().
			AnErr("Error", err).
			Msg("Can't bind flag api port")
		return
	}

}

var rootCMD = &cobra.Command{
	Use:   "service-calculate",
	Short: "This tool is used to calculate results based on a given list of votes.",
	Long: "It starts an API which mainly offers a few functionalities to calculate specific values for given votes.\n" +
		"Each endpoint accepts a list (an array) of votes and returns a results based on these votes.",
	DisableFlagsInUseLine: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !(viper.GetString("api.format") == "json" || viper.GetString("format") == "xml") {
			return errors.New("invalid api format set")
		}
		if viper.GetString("api.username") != "" && viper.GetString("api.password") == "" {
			return errors.New("username but no password for api authorization set")
		}
		if viper.GetString("api.username") == "" && viper.GetString("api.password") != "" {
			return errors.New("password but no username for api authorization set")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		api.StartAPI()
	},
}

// Execute is the entrypoint for the CLI interface.
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		os.Exit(1)
	}
}
