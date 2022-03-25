package packagemanager

import (
	"time"

	"github.com/theckman/yacspin"
)

func spinnerMessage(message string, spinner *yacspin.Spinner) {
	spinner.Message(" " + message)
}

func startSpinner(message string, spinner *yacspin.Spinner) {
	spinner.Start()
	spinnerMessage(message, spinner)
}

func stopSpinner(message string, spinner *yacspin.Spinner) {
	spinner.StopMessage(" " + message)
	spinner.Stop()
}

func failSpinner(message string, spinner *yacspin.Spinner) {
	spinner.StopFailMessage(" " + message)
	spinner.StopFail()
}

func initSpinners() {

	findSpinner, _ = yacspin.New(
		yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[59],
			Prefix:            "[gibson-cli] ",
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	)

	getSpinner, _ = yacspin.New(
		yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[59],
			Prefix:            "[gibson-cli] ",
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	)

	downloadSpinner, _ = yacspin.New(
		yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[59],
			Prefix:            "[gibson-cli] ",
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	)

	installSpinner, _ = yacspin.New(
		yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[59],
			Prefix:            "[gibson-cli] ",
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	)

	cacheSpinner, _ = yacspin.New(
		yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[59],
			Prefix:            "[gibson-cli] ",
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	)
}
