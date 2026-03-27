package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"strings"
	"testing"
	"time"
)

func freeTCPAddress(t *testing.T) string {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Could not allocate TCP port: %s", err)
	}
	addr := listener.Addr().String()
	if err := listener.Close(); err != nil {
		t.Fatalf("Could not release TCP port: %s", err)
	}
	return addr
}

func writeTempTLSCertFiles(t *testing.T) (string, string) {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Could not generate private key: %s", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(time.Now().UnixNano()),
		Subject: pkix.Name{
			CommonName: "localhost",
		},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              []string{"localhost"},
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		t.Fatalf("Could not generate certificate: %s", err)
	}

	tmpDir := t.TempDir()
	certPath := tmpDir + "/smtp-test-cert.pem"
	keyPath := tmpDir + "/smtp-test-key.pem"

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})

	if err := os.WriteFile(certPath, certPEM, 0600); err != nil {
		t.Fatalf("Could not write certificate file: %s", err)
	}
	if err := os.WriteFile(keyPath, keyPEM, 0600); err != nil {
		t.Fatalf("Could not write key file: %s", err)
	}

	return certPath, keyPath
}

func TestSMTPLibRoundTrip(t *testing.T) {
	address := freeTCPAddress(t)
	program := fmt.Sprintf(`
		smtp := import('smtp')
		{ wait: wait } := import('std')

		received := ''
		closeServer := smtp.listen('%s', {
			onMessage: fn(message) {
				received <- message.body
				{ ok?: true, message: 'queued' }
			}
		})

		wait(0.05)
		result := smtp.connect('%s')
		client := result.client
		client.ehlo('localhost')
		client.send({
			from: 'sender@example.com'
			to: ['rcpt@example.com']
			raw: 'Subject: Test\r\n\r\nhello smtp'
		})
		client.quit()
		wait(0.05)
		closeServer()
		received
	`, address, address)

	expectProgramToReturn(t, program, MakeString("Subject: Test\r\n\r\nhello smtp"))
}

func TestSMTPLibStartTLSRoundTrip(t *testing.T) {
	address := freeTCPAddress(t)
	certPath, keyPath := writeTempTLSCertFiles(t)
	certPath = strings.ReplaceAll(certPath, "\\", "/")
	keyPath = strings.ReplaceAll(keyPath, "\\", "/")

	program := fmt.Sprintf(`
		smtp := import('smtp')
		{ wait: wait } := import('std')

		receivedTls := false
		closeServer := smtp.listen('%s', {
			onMessage: fn(message) {
				receivedTls <- message.tls?
				{ ok?: true, message: 'queued' }
			}
		}, {
			startTLS?: true
			certFile: '%s'
			keyFile: '%s'
		})

		wait(0.05)
		result := smtp.connect('%s')
		client := result.client
		client.ehlo('localhost')
		client.startTLS({
			insecureSkipVerify: true
			serverName: 'localhost'
		})
		client.send({
			from: 'sender@example.com'
			to: ['rcpt@example.com']
			raw: 'Subject: TLS\r\n\r\nhello tls'
		})
		client.quit()
		wait(0.05)
		closeServer()
		receivedTls
	`, address, certPath, keyPath, address)

	expectProgramToReturn(t, program, oakTrue)
}

func TestPOPLibRetrieve(t *testing.T) {
	address := freeTCPAddress(t)
	program := fmt.Sprintf(`
		pop := import('pop')
		{ wait: wait } := import('std')
		{ join: join } := import('str')

		closeServer := pop.listen('%s', {
			auth: fn(user, pass) user = 'demo' & pass = 'secret'
			messages: fn(_) [{ uid: 'm1', data: 'Subject: Test\r\n\r\nhello pop' }]
		})

		wait(0.05)
		result := pop.connect('%s')
		client := result.client
		client.login('demo', 'secret')
		msg := client.retr(1)
		client.quit()
		wait(0.05)
		closeServer()
		join(msg.lines, '\n')
	`, address, address)

	expectProgramToReturn(t, program, MakeString("Subject: Test\n\nhello pop"))
}

func TestPOPLibSTLSRetrieve(t *testing.T) {
	address := freeTCPAddress(t)
	certPath, keyPath := writeTempTLSCertFiles(t)
	certPath = strings.ReplaceAll(certPath, "\\", "/")
	keyPath = strings.ReplaceAll(keyPath, "\\", "/")

	program := fmt.Sprintf(`
		pop := import('pop')
		{ wait: wait } := import('std')
		{ join: join } := import('str')

		closeServer := pop.listen('%s', {
			auth: fn(user, pass) user = 'demo' & pass = 'secret'
			messages: fn(_) [{ uid: 'm1', data: 'Subject: TLS\r\n\r\nhello pop tls' }]
		}, {
			startTLS?: true
			certFile: '%s'
			keyFile: '%s'
		})

		wait(0.05)
		result := pop.connect('%s')
		client := result.client
		stls := client.startTLS({
			insecureSkipVerify: true
			serverName: 'localhost'
		})
		if stls.type {
			:error -> {
				closeServer()
				stls.error
			}
			_ -> {
				client.login('demo', 'secret')
				msg := client.retr(1)
				client.quit()
				wait(0.05)
				closeServer()
				join(msg.lines, '\n')
			}
		}
	`, address, certPath, keyPath, address)

	expectProgramToReturn(t, program, MakeString("Subject: TLS\n\nhello pop tls"))
}

func TestIMAPLibFetchBody(t *testing.T) {
	address := freeTCPAddress(t)
	program := fmt.Sprintf(`
		imap := import('imap')
		{ wait: wait } := import('std')

		closeServer := imap.listen('%s', {
			auth: fn(user, pass) user = 'demo' & pass = 'secret'
			mailboxes: fn(_) [{
				name: 'INBOX'
				messages: [{
					header: 'Subject: Test\r\n\r\n'
					body: 'hello imap'
				}]
			}]
		})

		wait(0.05)
		result := imap.connect('%s')
		client := result.client
		client.capability()
		client.login('demo', 'secret')
		client.select('INBOX')
		fetched := client.fetch(1, 'BODY[]')
		client.logout()
		wait(0.05)
		closeServer()
		fetched.entries.0.literal
	`, address, address)

	expectProgramToReturn(t, program, MakeString("hello imap"))
}

func TestIMAPLibStartTLSFetchBody(t *testing.T) {
	address := freeTCPAddress(t)
	certPath, keyPath := writeTempTLSCertFiles(t)
	certPath = strings.ReplaceAll(certPath, "\\", "/")
	keyPath = strings.ReplaceAll(keyPath, "\\", "/")

	program := fmt.Sprintf(`
		imap := import('imap')
		{ wait: wait } := import('std')

		closeServer := imap.listen('%s', {
			auth: fn(user, pass) user = 'demo' & pass = 'secret'
			mailboxes: fn(_) [{
				name: 'INBOX'
				messages: [{
					header: 'Subject: TLS Test\r\n\r\n'
					body: 'hello imap tls'
				}]
			}]
		}, {
			startTLS?: true
			certFile: '%s'
			keyFile: '%s'
		})

		wait(0.05)
		result := imap.connect('%s')
		client := result.client
		tlsResult := client.startTLS({
			insecureSkipVerify: true
			serverName: 'localhost'
		})
		if tlsResult.type {
			:error -> {
				closeServer()
				tlsResult.error
			}
			_ -> {
				client.login('demo', 'secret')
				client.select('INBOX')
				fetched := client.fetch(1, 'BODY[]')
				client.logout()
				wait(0.05)
				closeServer()
				fetched.entries.0.literal
			}
		}
	`, address, certPath, keyPath, address)

	expectProgramToReturn(t, program, MakeString("hello imap tls"))
}

func TestMailLibSourcesLoad(t *testing.T) {
	for _, name := range []string{"socket", "smtp", "pop", "imap"} {
		if _, ok := stdlibs[name]; !ok {
			t.Fatalf("Expected stdlib %s to be registered", name)
		}
		if strings.TrimSpace(stdlibs[name]) == "" {
			t.Fatalf("Expected stdlib %s source to be non-empty", name)
		}
	}
}
