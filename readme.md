# Shrun

This application is meant for running commands on a remote host or multiple hosts via an ssh connection. **When prompted for a password on the remote server, the password will be supplied from the environment**

## Usage

To run the application you have two arguments to pass in. First you can pass in a command
string `"sudo puppet agent -t"` and a comma separated list of hosts
`"dev-test-app-app200.myhost.com,`dev-test-app-app201.myhost.com"`.

For this to work you will need to set the `"password"` environment variable to your password. The cli will try to get
the username from the system and if that fails it will try to read in the `"username"` environment variable.

```shell
shrun "sudo puppet agent -t" dev-test-app-app200.myhost.com
```

### Advanced flags

this is the output of `shrun --help`

```shell
Usage:
  shrun [flags]

Examples:
shrun "sudo run puppet agent -t" my.host.com,my.host1.com,my.second.host.com

Flags:
  -e, --explode-logs     would you like to create a dir and files for each host logs
  -h, --help             help for shrun
  -n, --no-prefix        if flag is set will not add prefix to the log
  -c, --run-concurrent   would you like to run commands on hosts concurrently
```

`-e` or `--explode-logs` will create a directory with the current time and then a log file per host you are executing
the command on.

`-n` or `--no-prefix` will remove the normal prefix that is added to log messages. This is useful if you want a clean
output for the `cat` command or something of that nature.
*the prefix helps distinguish which host the log is coming from and the command being executed.*

`-c` or `--run-concurrent` setting this flag will execute all ssh commands remote host concurrently.

# Building

to build a binary you will have to `Go` installed.

*execute all commands in the root dir*

**Build on Windows**

```shell
go build -o shrun.exe .
```

**Build on Mac/linux**

```shell
go build -o shrun .
```

**build_binaries.sh**
this will build cross-platform binaries for some of the most common configurations. Binaries are located in
the `/binaries` directory.

```shell
% ls binaries 
shrun_darwin_amd64      shrun_darwin_arm64      shrun_linux_386         shrun_linux_amd64       shrun_linux_arm64       shrun_windows_386.exe   shrun_windows_amd64.exe

```
