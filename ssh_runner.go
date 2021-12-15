package main

import "io"

type baseSshRunner struct {
	creds *Credentials
}

func NewBaseSshRunner(creds *Credentials) SshRunner {
	return &baseSshRunner{
		creds: creds,
	}
}

func (b baseSshRunner) RunCommand(host string, command string, output io.Writer) {
	conn, err := MakeSshConnection(host, b.creds.Username, b.creds.Password)
	if err != nil {
		panic(err)
	}

	conn.SendCommands(command, output)
}
