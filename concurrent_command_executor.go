package main

import "sync"

type concurrentCommandExecutor struct {
	sshRunner SshRunner
}

func NewConcurrentCommandExecutor(sshRunner SshRunner) CommandExecutor {
	return &concurrentCommandExecutor{sshRunner: sshRunner}
}

func (c concurrentCommandExecutor) RunCommands(hosts []string, command string, outputSupplier OutputSupplier) {
	wg := &sync.WaitGroup{}
	for i := range hosts {
		// run command executor async
		wg.Add(1)
		go func(host string) {
			output := outputSupplier(host, command)
			c.sshRunner.RunCommand(host, command, output)
			wg.Done()
		}(hosts[i])
	}
	wg.Wait()
}
