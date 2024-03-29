package ssh

import (
	"encoding/base64"
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-ssh/result"
	"golang.org/x/crypto/ssh"
	"net/http"
	"os"
)

type SSHArguments struct {
	Command  string `json:"command,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
}

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

type Output struct {
	Output     string `json:"standard_output"`
	Error      string `json:"standard_error"`
	ReturnCode int    `json:"return_code"`
}

//SSH
func SSH(responseWriter http.ResponseWriter, request *http.Request) {

	var privateKey = os.Getenv("PRIVATE_KEY")
	decoder := json.NewDecoder(request.Body)

	var param SSHArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	if param.Password != "" && privateKey != "" {
		message := Message{"false", "Please provide either password or private key", http.StatusBadRequest}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	} else if param.Password == "" && privateKey == "" {
		message := Message{"false", "Please provide password or private key", http.StatusBadRequest}
		bytes, _ := json.Marshal(message)
		result.WriteJsonResponse(responseWriter, bytes, http.StatusBadRequest)
		return
	}

	var hostname string
	if param.Port != "" {
		hostname = param.Host + ":" + param.Port
	} else {
		hostname = param.Host + ":22"
	}

	client, session, err := connectToHost(param.Username, hostname, param.Password, privateKey)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	out, outputErr := session.CombinedOutput(param.Command)
	if outputErr != nil {
		result.WriteErrorResponse(responseWriter, outputErr)
		return
	}

	client.Close()
	message := Output{string(out), "", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

func connectToHost(user, host, password, privateKey string) (*ssh.Client, *ssh.Session, error) {

	var sshConfig *ssh.ClientConfig

	if password != "" {
		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{ssh.Password(password)},
		}
		sshConfig = config

	} else if privateKey != "" {
		pemBytes, err := base64.StdEncoding.DecodeString(privateKey)
		if err != nil {
			return nil, nil, err
		}

		signer, err := ssh.ParsePrivateKey(pemBytes)
		if err != nil {
			return nil, nil, err
		}

		config := &ssh.ClientConfig{
			User: user,
			Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
		}
		sshConfig = config
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
