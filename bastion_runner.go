package main

import "io"

type bastionRunner struct {
	creds *Credentials
}

func NewBastionSshRunner(creds *Credentials) SshRunner {
	return &bastionRunner{
		creds: creds,
	}
}

func (b bastionRunner) RunCommand(host string, command string, output io.Writer) {
	//TODO implement me
	panic("implement me")
}
