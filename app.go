package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	client, err := ssh.Dial("tcp", "lanks.top:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("SSH dial error: %s", err.Error())
	}

	// 建立新会话
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run("ls")
	if err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
