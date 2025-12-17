package structures

import (
	"time"
)

type Logs struct {
	userInput      string
	userInputError string
	userInputTime  time.Time
}

func (a *Logs) GetUserInput() string {
	return a.userInput
}

func (a *Logs) GetUserInputError() string {
	return a.userInputError
}

func (a *Logs) GetUserInputTime() time.Time {
	return a.userInputTime
}

func (a *Logs) ChangeError(x string) {
	a.userInputError = x
}

func MakeLogs(
	userInput string,
	userInputError string,
	userInputTime time.Time,
) *Logs {
	return &Logs{
		userInput:      userInput,
		userInputError: userInputError,
		userInputTime:  userInputTime,
	}
}

func AddLog(logs *[]*Logs, textSlice []string, isError string) {
	cmd := ""
	for i := range textSlice {
		cmd += textSlice[i]

		if i != len(textSlice)-1 {
			cmd += " "
		}
	}
	*logs = append(*logs, MakeLogs(
		cmd,
		isError,
		time.Now(),
	))

}
