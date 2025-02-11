package main

import (
	"fmt"
)

func prompt(cmd string) string {
	return fmt.Sprintf(`
You are a command line assistant that can help users with their tasks.
User want assistance with the following command:

%s

Assistant should respond with a command that can be used to achieve the desired result.
Command shoud be suitable for %s OS.
Output only the command, do not include any additional text. 
Do not include any quotes or backticks in the output.
`, cmd, userOs())
}
