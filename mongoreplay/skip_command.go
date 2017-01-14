package mongoreplay

import "sync"

var skipCommands map[string]bool
var skipCommandsMutex sync.RWMutex

func initiailizeSkipCommands(commandsToSkip []string) {
	skipCommandsMutex.Lock()
	defer skipCommandsMutex.Unlock()
	skipCommands = make(map[string]bool)
	for _, cmd := range commandsToSkip {
		skipCommands[cmd] = true
	}
}

func shouldSkipCommand(cmd string) bool {
	skipCommandsMutex.RLock()
	defer skipCommandsMutex.RUnlock()
	return skipCommands[cmd]
}
