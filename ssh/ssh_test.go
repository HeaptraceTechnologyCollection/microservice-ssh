package ssh

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var _ = Describe("SSH", func() {

	os.Setenv("HOST", "192.168.1.88:22")
	os.Setenv("USERNAME", "admin1")
	os.Setenv("PASSWORD", "admin")

	ssh := SSHArguments{Command: "ifconfig"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(ssh)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/ssh", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SSH)
	handler.ServeHTTP(recorder, request)

	Describe("Run command on SSH server", func() {
		Context("SSH server", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
