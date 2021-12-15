package main

type syncCommandExecutor struct {
	sshRunner SshRunner
}

func NewSyncCommandExecutor(sshRunner SshRunner) CommandExecutor {
	return &syncCommandExecutor{sshRunner: sshRunner}
}

func (s syncCommandExecutor) RunCommands(hosts []string, command string, outputSupplier OutputSupplier) {
	// run commands on all host synchronously
	for i := range hosts {
		output := outputSupplier(hosts[i], command)
		s.sshRunner.RunCommand(hosts[i], command, output)
	}
}
