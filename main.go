package main

import (
	"github.com/spf13/cobra"
	"io"
	"strings"
)

type OutputSupplier func(string, string) io.Writer

type SshRunner interface {
	RunCommand(host string, command string, writer io.Writer)
}

type CommandExecutor interface {
	RunCommands(hosts []string, command string, writerSupplier OutputSupplier)
}

var (
	RunConcurrent = false
	ExplodeLogs   = false
	NoPrefix      = false
	rootCmd       = &cobra.Command{
		Use:     "shrun",
		Short:   "ssh runner",
		Example: `shrun "sudo run puppet agent -t" my.host.com,my.host1.com,my.second.host.com`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				panic("must past command and host(csv) to program")
			}
			runCommandOnHosts(NoPrefix, RunConcurrent, ExplodeLogs, args[0], args[1])
		},
	}
)

func main() {
	rootCmd.PersistentFlags().BoolVarP(&NoPrefix, "no-prefix", "n", false, "if flag is set will not add prefix to the log")
	rootCmd.PersistentFlags().BoolVarP(&RunConcurrent, "run-concurrent", "c", false, "would you like to run commands on hosts concurrently")
	rootCmd.PersistentFlags().BoolVarP(&ExplodeLogs, "explode-logs", "e", false, "would you like to create a dir and files for each host logs")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func runCommandOnHosts(noPrefix, runConcurrent, explodeLogs bool, command string, hostCsv string) {
	creds := GetCredentialsFromEnvironmentVariables()

	// how are we going to connect to remote hosts
	var sshRunner SshRunner
	sshRunner = NewBaseSshRunner(creds)

	// commands executor
	var commandExecutor CommandExecutor
	if runConcurrent {
		commandExecutor = NewConcurrentCommandExecutor(sshRunner)
	} else {
		commandExecutor = NewSyncCommandExecutor(sshRunner)
	}

	// output handler
	// if explode we're going to write output per fill
	var outputSupplier OutputSupplier
	if explodeLogs {
		outputSupplier = ExplodedLogOutputSupplier
	} else {
		outputSupplier = getOutputSupplier(noPrefix)
	}

	// split hosts
	hosts := strings.Split(hostCsv, ",")
	commandExecutor.RunCommands(hosts, command, outputSupplier)
}
