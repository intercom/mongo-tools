package mongoreplay

import "sync"

var skipCommands map[string]bool
var skipCommandsMutex sync.RWMutex

var skipDatabases map[string]bool
var skipDatabasesMutex sync.RWMutex

func initiailizeSkipCommands(commandsToSkip []string) {
	skipCommandsMutex.Lock()
	defer skipCommandsMutex.Unlock()
	skipCommands = make(map[string]bool)
	for _, cmd := range commandsToSkip {
		skipCommands[cmd] = true
	}
}

func shouldSkipCommandByOpType(cmd string) bool {
	skipCommandsMutex.RLock()
	defer skipCommandsMutex.RUnlock()
	return skipCommands[cmd]
}

func initiailizeSkipDatabases(databasesToSkip []string) {
	skipDatabasesMutex.Lock()
	defer skipDatabasesMutex.Unlock()
	skipDatabases = make(map[string]bool)
	for _, db := range databasesToSkip {
		skipDatabases[db] = true
	}
}

func shouldSkipCommandByDatabaseName(database string) bool {
	skipDatabasesMutex.RLock()
	defer skipDatabasesMutex.RUnlock()
	return skipDatabases[database]
}
