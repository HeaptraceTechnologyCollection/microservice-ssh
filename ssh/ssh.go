package ssh

import (
	"encoding/json"
	"net/http"
	"os"

	result "github.com/heaptracetechnology/microservice-ssh/result"
	"golang.org/x/crypto/ssh"
)

type SSHArguments struct {
	Command string `json:"command,omitempty"`
}

var pwd string

//SSH
func SSH(responseWriter http.ResponseWriter, request *http.Request) {

	var host = os.Getenv("HOST")
	var port = os.Getenv("PORT")
	var username = os.Getenv("USRNAME")
	var password = os.Getenv("PASSWORD")
	pwd = password

	decoder := json.NewDecoder(request.Body)

	var param SSHArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	var hostname string
	if port != "" {
		hostname = host + ":" + port
	} else {
		hostname = host + ":22"
	}

	client, session, err := connectToHost(username, hostname)
	if err != nil {
		panic(err)
	}
	out, err := session.CombinedOutput(param.Command)
	if err != nil {
		panic(err)
	}

	bytes, _ := json.Marshal(string(out))
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
	client.Close()

}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {

	pass := pwd

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}

	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
