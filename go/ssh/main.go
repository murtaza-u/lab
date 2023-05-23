package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

var sigmap = map[string]ssh.Signal{
	"aborted":                  "ABRT",
	"alarm clock":              "ALRM",
	"floating point exception": "FPE",
	"hangup":                   "HUP",
	"illegal instruction":      "ILL",
	"interrupt":                "INT",
	"killed":                   "KILL",
	"broken pipe":              "PIPE",
	"quit":                     "QUIT",
	"segmentation fault":       "SEGV",
	"terminated":               "TERM",
	"user defined signal 1":    "USR1",
	"user defined signal 2":    "USR2",
}

func interactive() {
	key, err := os.ReadFile("key")
	if err != nil {
		log.Fatal(err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal(err)
	}

	callback := ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", "3.109.50.164:22", &ssh.ClientConfig{
		User:            "addmb",
		HostKeyCallback: callback,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	sess, err := client.NewSession()
	if err != nil {
		return
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	term := os.Getenv("TERM")
	if term == "" {
		term = "xterm-256color"
	}

	width, height, err := terminal.GetSize(0)
	if err != nil {
		height = 80
		width = 40
	}

	if err := sess.RequestPty(term, height, width, modes); err != nil {
		sess.Close()
		log.Fatalf("request for pseudo terminal failed: %s", err.Error())
	}

	stdin, err := sess.StdinPipe()
	if err != nil {
		sess.Close()
		log.Fatalf("unable to setup stdin for session: %s", err.Error())
	}

	stdout, err := sess.StdoutPipe()
	if err != nil {
		sess.Close()
		log.Fatalf("unable to setup stdout for session: %s", err.Error())
	}

	stderr, err := sess.StderrPipe()
	if err != nil {
		sess.Close()
		log.Fatalf("unable to setup stderr for session: %s", err.Error())
	}

	go io.Copy(stdin, os.Stdin)

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go copy(wg, os.Stdout, stdout)
	go copy(wg, os.Stderr, stderr)

	err = sess.Shell()
	if err != nil {
		fmt.Println(err.Error())
	}

	sigssh := make(chan os.Signal)
	signal.Notify(
		sigssh,
		syscall.SIGABRT, syscall.SIGALRM, syscall.SIGFPE,
		syscall.SIGHUP, syscall.SIGILL, syscall.SIGINT, syscall.SIGKILL,
		syscall.SIGPIPE, syscall.SIGQUIT, syscall.SIGSEGV,
		syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2,
	)

	go func() {
		for {
			s := <-sigssh
			sess.Signal(sigmap[s.String()])
		}
	}()

	wg.Wait()
	sess.Close()
}

func main() {
	interactive()
}

func copy(wg *sync.WaitGroup, dst io.Writer, src io.Reader) {
	io.Copy(dst, src)
	wg.Done()
}
