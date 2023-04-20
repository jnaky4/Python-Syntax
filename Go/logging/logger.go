package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

const (
	LOGLEVEL = "loglevel"
)

//	TraceLevel	-1 	"trace"
//	DebugLevel 	0  	"debug"
//	InfoLevel 	1	"info"
//	WarnLevel 	2	"warn"
//	ErrorLevel 	3	"error"
//	FatalLevel 	4	"fatal"
//	PanicLevel	5	"panic"
//	NoLevel		6
//	Disabled	7

func NewLogger() zerolog.Logger {
	var level zerolog.Level
	var err error
	if viper.IsSet(LOGLEVEL) {
		if level, err = zerolog.ParseLevel(viper.GetString(LOGLEVEL)); err != nil {
			log.Warn().Msgf("error creating logger: %s", err.Error())
			level = zerolog.WarnLevel
		}
	}

	lg := zerolog.New(os.Stderr).
		Level(level).
		With().
		Timestamp().
		Logger().
		With().Str("contextMap", `{}`).Logger()
	zerolog.TimeFieldFormat = ""
	zerolog.TimestampFieldName = "timeMillis"

	log.Output(lg)

	lg.Debug().Msgf("loglevel set to %v", lg.GetLevel())
	return lg
}
