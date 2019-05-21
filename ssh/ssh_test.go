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

var _ = Describe("SSH without private key and password", func() {

	ssh := SSHArguments{Command: "ifconfig", Username: "admin1", Host: "192.168.1.88"}
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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("SSH with password", func() {

	ssh := SSHArguments{Command: "ifconfig", Username: "admin1", Password: "admin", Host: "192.168.1.88"}
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
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("SSH with private key and password", func() {

	privateKey := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBcSt5cVhscDJrQk9USVFHZ01WYVhHdVFWVE5uUmtJNlJPdzJuV1ZkdTZQblRBL01SCkw3UTZwRzVBMkQyR2hENkswUS9sbUdSYk9GMkgzaTkrN1ZmNTNtWGxuVGdLZm5jWTRpZFkwL0xFUkVZNUJJMGcKR1dGaStyUzhoZDRZOURtK2hVcE0zVTMybzN0Ym52UHBTa1N4cVlwSjF0RTZ3VjNXWnRpd3RiYzFsenVuU0lITgpDUlVMT2kydUhiVzJ1bEFiaktUcnBqbENWRDY4SG5laGxWUEs5c0FDcFh2VTg2UjdSekJZUnRVbldETXdQUHVUCmxXa0E1L0xnMFYxRjhremw2akgzckhiKzY3Y2tJcjRqL0d2MVJWUFB3UFdNRUhLWGtZT3kxYkp1K0h4YmhkY2IKNEhyTThLZXhTc0lpd0NFa2NVcGl2YVN6K2QvWWcwS2VtOEh3VlFJREFRQUJBb0lCQUJUZkFqKzZFN0toN2JhUQpEL1p0WUtLdkZiYmlxb0kyRElOeXdPSWpyeDh4Vk9DaDNYQkJITzFoUlJYN1FoMUR5bTVlMDZ5UVlsS1JhREVjCjZ2ZTlPbVE1VW9xbmh2NmJUcllGdU03aUpkbHovUEFFZ2VVUzRZVTE4N3o4bldMb3I2eFd2TVlROU9RYjBaK1cKVkxmalp2NEIvN3pJWWEyZnlxdHBtdW8wZDVrSVBlcWFYWDBDMXVGSUo4ckhsbTVhTFFSZStFaXBtcHlzdlN4TgplQVgvRDFIUm4ydGMwYTQwRnpwN0JUOVR5bTFkM2w3YStqUkNEdWdWUGVMMUF1c2s3NDNQRTlQOEFvb3NaREZJCitLSEllT2N2Zkc2NS9DY3R0S1lrVWJpQ2dzZnZHNlVROVJMVjhQc2lURTNzdDZJRDB5djRPdlhyWUtzdXhrNSsKNGxpZVlIRUNnWUVBMlZRcUxqZjdQcFlmeDZ0VkcxYmJMVDlVRXVyRmpiL3VWL2RzNmpES2xraWFDRlBGOHZoZgpqN3FjNmlpUzk0dHl6MG5DTnZ1NHkxc3hMYVlVWTRpcThUWlQzMTIzY1ZYQVJTSENGa1ZHcDFnUzErdUpqL1kyCjlwTjUwVStRUk1acE94ZHFiVXJZVENNY3ZYRWxZL1RLSjdqbk5vbkRQNzc1MXdIN0lvR1dFcnNDZ1lFQXlvUTYKaDRwdCtsUUhiaU9neXhPamFpcHZ0eWE3emg1d1FGYVVLbm1MZXUwdlZXbWtlVDZudU1GbkNPRFUyM1Y2WHIwZgpsZ2RXeUxrSEFwTlE3c041dGxIV0Qvakd6dVl2Q2xNUUVnQzdLYVJrZ1BjUXhGV2kyN0Fqb0gveXNjWnFCSi8xCkxCVFVjcHhGVnZLdGJWenRCYnlHZVJuT2MyQWM2bzNhc3hPVmdDOENnWUVBckp1SElNdy9sTmJCQ09HUUo3V1YKUUZ1aTE1OTFKZjhCT3daOWo0Y083OHRiNHk4OThacklzeXZnd3ExVkJKelJvOGNPSklOS291Q2JyNGpQZXJvcQpJb2dtbHlva3J6UVBFQmtld0hkbkJUUVRTMEI4TWtXNEk2Qy95TGtyZVNRb29kRVlLeE9kdE9MU1NiZmFuWWZuCkl5TmRKOWpFcFJWMTh3bFV2M1F5a0U4Q2dZQjIwSkZPU0VjeDRPN2pEWUFlNVF5eEV1aXNPY3RocUxZTzZUelEKbHJMZ2todDlMeGZTRXBKd2NQZTBXOFJHWld3Ly9SRjFBaVZHYWxmVWlQMm90NExIRnNoU1lwQ3hmcGNHcGFqKwpCdlBJQUt6K2hQV1BXdmJMa1ZHMXJwdUM5WGZwOHJieS85Mk15R1plRnM3dEpPSGl4YkxYaGU1Ny9sMjR0elVpCmIxRDgvUUtCZ1FEUWNQNnV3Wkx5dFdUV05qRDFJbU4vZVRDUU0yLzdyR1ZUUU5pb1g2RExSQk43aktVaFhrUmEKU0ZoclFMU1dsYWtPSjY0TFd1TXV1a1V3RE5laTNkK2UvazlCYUhRTStrU25NR1FISkhsZU5ncmU1RWJBVmppNwpHaEJ5aFlUVlh0NW40bG5hZEFMdzh4bzlXRGxlM1h3RlFuQjdyWkxNTHkyYWw5WCt0eDRPVXc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="

	os.Setenv("PRIVATE_KEY", privateKey)
	ssh := SSHArguments{Command: "ifconfig", Username: "admin1", Host: "192.168.1.88", Password: "admin"}
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
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("SSH with private key", func() {

	privateKey := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBcSt5cVhscDJrQk9USVFHZ01WYVhHdVFWVE5uUmtJNlJPdzJuV1ZkdTZQblRBL01SCkw3UTZwRzVBMkQyR2hENkswUS9sbUdSYk9GMkgzaTkrN1ZmNTNtWGxuVGdLZm5jWTRpZFkwL0xFUkVZNUJJMGcKR1dGaStyUzhoZDRZOURtK2hVcE0zVTMybzN0Ym52UHBTa1N4cVlwSjF0RTZ3VjNXWnRpd3RiYzFsenVuU0lITgpDUlVMT2kydUhiVzJ1bEFiaktUcnBqbENWRDY4SG5laGxWUEs5c0FDcFh2VTg2UjdSekJZUnRVbldETXdQUHVUCmxXa0E1L0xnMFYxRjhremw2akgzckhiKzY3Y2tJcjRqL0d2MVJWUFB3UFdNRUhLWGtZT3kxYkp1K0h4YmhkY2IKNEhyTThLZXhTc0lpd0NFa2NVcGl2YVN6K2QvWWcwS2VtOEh3VlFJREFRQUJBb0lCQUJUZkFqKzZFN0toN2JhUQpEL1p0WUtLdkZiYmlxb0kyRElOeXdPSWpyeDh4Vk9DaDNYQkJITzFoUlJYN1FoMUR5bTVlMDZ5UVlsS1JhREVjCjZ2ZTlPbVE1VW9xbmh2NmJUcllGdU03aUpkbHovUEFFZ2VVUzRZVTE4N3o4bldMb3I2eFd2TVlROU9RYjBaK1cKVkxmalp2NEIvN3pJWWEyZnlxdHBtdW8wZDVrSVBlcWFYWDBDMXVGSUo4ckhsbTVhTFFSZStFaXBtcHlzdlN4TgplQVgvRDFIUm4ydGMwYTQwRnpwN0JUOVR5bTFkM2w3YStqUkNEdWdWUGVMMUF1c2s3NDNQRTlQOEFvb3NaREZJCitLSEllT2N2Zkc2NS9DY3R0S1lrVWJpQ2dzZnZHNlVROVJMVjhQc2lURTNzdDZJRDB5djRPdlhyWUtzdXhrNSsKNGxpZVlIRUNnWUVBMlZRcUxqZjdQcFlmeDZ0VkcxYmJMVDlVRXVyRmpiL3VWL2RzNmpES2xraWFDRlBGOHZoZgpqN3FjNmlpUzk0dHl6MG5DTnZ1NHkxc3hMYVlVWTRpcThUWlQzMTIzY1ZYQVJTSENGa1ZHcDFnUzErdUpqL1kyCjlwTjUwVStRUk1acE94ZHFiVXJZVENNY3ZYRWxZL1RLSjdqbk5vbkRQNzc1MXdIN0lvR1dFcnNDZ1lFQXlvUTYKaDRwdCtsUUhiaU9neXhPamFpcHZ0eWE3emg1d1FGYVVLbm1MZXUwdlZXbWtlVDZudU1GbkNPRFUyM1Y2WHIwZgpsZ2RXeUxrSEFwTlE3c041dGxIV0Qvakd6dVl2Q2xNUUVnQzdLYVJrZ1BjUXhGV2kyN0Fqb0gveXNjWnFCSi8xCkxCVFVjcHhGVnZLdGJWenRCYnlHZVJuT2MyQWM2bzNhc3hPVmdDOENnWUVBckp1SElNdy9sTmJCQ09HUUo3V1YKUUZ1aTE1OTFKZjhCT3daOWo0Y083OHRiNHk4OThacklzeXZnd3ExVkJKelJvOGNPSklOS291Q2JyNGpQZXJvcQpJb2dtbHlva3J6UVBFQmtld0hkbkJUUVRTMEI4TWtXNEk2Qy95TGtyZVNRb29kRVlLeE9kdE9MU1NiZmFuWWZuCkl5TmRKOWpFcFJWMTh3bFV2M1F5a0U4Q2dZQjIwSkZPU0VjeDRPN2pEWUFlNVF5eEV1aXNPY3RocUxZTzZUelEKbHJMZ2todDlMeGZTRXBKd2NQZTBXOFJHWld3Ly9SRjFBaVZHYWxmVWlQMm90NExIRnNoU1lwQ3hmcGNHcGFqKwpCdlBJQUt6K2hQV1BXdmJMa1ZHMXJwdUM5WGZwOHJieS85Mk15R1plRnM3dEpPSGl4YkxYaGU1Ny9sMjR0elVpCmIxRDgvUUtCZ1FEUWNQNnV3Wkx5dFdUV05qRDFJbU4vZVRDUU0yLzdyR1ZUUU5pb1g2RExSQk43aktVaFhrUmEKU0ZoclFMU1dsYWtPSjY0TFd1TXV1a1V3RE5laTNkK2UvazlCYUhRTStrU25NR1FISkhsZU5ncmU1RWJBVmppNwpHaEJ5aFlUVlh0NW40bG5hZEFMdzh4bzlXRGxlM1h3RlFuQjdyWkxNTHkyYWw5WCt0eDRPVXc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="

	os.Setenv("PRIVATE_KEY", privateKey)
	ssh := SSHArguments{Command: "ifconfig", Username: "admin1", Host: "192.168.1.88"}
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
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
