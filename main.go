package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {

	//////////Normal Logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log.Debug().
		Str("Name", "Tom").
		Send()

	////////////////////////LOG LEVEL
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Debug().Msg("This message appears only when log level set to Debug")
	log.Info().Msg("This message appears when log level set to Debug or Info")

	if e := log.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}
	//////////Create logger instance to manage different outputs

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	logger.Info().Str("foo", "bar").Msg("hello world")

	//Pretty logging
	//To log a human-friendly, colorized output, use zerolog.ConsoleWriter:
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Str("foo", "bar").Msg("Hello world")

}

//////////Normal Logging
// Output: {"level":"debug","Scale":"833 cents","Interval":833.09,"time":1562212768,"message":"Fibonacci is everywhere"}
// Output: {"level":"debug","Name":"Tom","time":1562212768}

///////////LOG LVL
/*
Info Output (no flag)

$ ./logLevelExample
{"time":1516387492,"level":"info","message":"This message appears when log level set to Debug or Info"}
Debug Output (debug flag set)

$ ./logLevelExample -debug
{"time":1516387573,"level":"debug","message":"This message appears only when log level set to Debug"}
{"time":1516387573,"level":"info","message":"This message appears when log level set to Debug or Info"}
{"time":1516387573,"level":"debug","foo":"bar","message":"some debug message"}
*/

//////////Create logger instance to manage different outputs
// Output: {"level":"info","time":1494567715,"message":"hello world","foo":"bar"}

//Pretty logging
//To log a human-friendly, colorized output, use zerolog.ConsoleWriter:
// Output: 3:04PM INF Hello World foo=bar
