package main

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"io"
	"strings"
)

type SshConnection struct {
	*ssh.Client
	password string
}

func MakeSshConnection(host string, username, password string) (*SshConnection, error) {

	// ssh client password and username
	sshConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// make the ssh connection
	// defaulting  the :22 port now
	// should probably be smarter about that
	conn, err := ssh.Dial("tcp", host+":22", sshConfig)
	if err != nil {
		return nil, err
	}

	return &SshConnection{
		Client:   conn,
		password: password,
	}, nil
}

func (this SshConnection) SendCommands(cmd string, output io.Writer) {
	session, err := this.Client.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		panic(err)
	}

	// RUN COMMAND
	stdoutB := new(bytes.Buffer)
	newStdOut := io.MultiWriter(stdoutB, output)
	session.Stdout = newStdOut
	in, _ := session.StdinPipe()

	// if we execute a sudo command and there is a request for password send IN password
	go func(in io.Writer, output *bytes.Buffer) {
		for {
			if strings.Contains(string(output.Bytes()), "[sudo] password for ") {
				_, err = in.Write([]byte(this.password + "\n"))
				if err != nil {
					println(err.Error())
					break
				}
				break
			}
		}
	}(in, stdoutB)

	err = session.Run(cmd)
	if err != nil {
		panic(err)
	}
}
