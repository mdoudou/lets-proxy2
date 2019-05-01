//nolint:golint
package cert_manager

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/rekby/lets-proxy2/internal/cache"

	"go.uber.org/zap"

	zc "github.com/rekby/zapcontext"

	"github.com/gojuno/minimock"

	"github.com/maxatome/go-testdeep"
	"github.com/rekby/lets-proxy2/internal/th"

	"golang.org/x/crypto/acme"
)

const testACMEServer = "http://localhost:4000/directory"
const rsaKeyLength = 2048

type contextConnection struct {
	net.Conn
	context.Context
}

func (c contextConnection) GetContext() context.Context {
	return c.Context
}

//nolint:gochecknoinits
func init() {
	zc.SetDefaultLogger(zap.NewNop())
}

func createTestClient(t *testing.T) *acme.Client {
	_, err := http.Get(testACMEServer)
	if err != nil {
		t.Fatalf("Can't connect to buoulder server: %q", err)
	}

	client := acme.Client{}
	client.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				//nolint:gosec
				InsecureSkipVerify: true,
			},
		},
	}

	client.DirectoryURL = testACMEServer
	client.Key, _ = rsa.GenerateKey(rand.Reader, rsaKeyLength)
	_, err = client.Register(context.Background(), &acme.Account{}, func(tosURL string) bool {
		return true
	})

	if err != nil {
		t.Fatal("Can't initialize acme client.")
	}
	return &client
}

func TestStoreCertificate(t *testing.T) {
	ctx, flush := th.TestContext()
	defer flush()

	//nolint:gosec
	key, _ := rsa.GenerateKey(rand.Reader, 512)

	cert := &tls.Certificate{Certificate: [][]byte{
		{1, 2, 3},
		{4, 5, 6},
	},
		PrivateKey: key,
	}

	mc := minimock.NewController(t)
	cacheMock := NewCacheMock(mc).PutMock.Set(func(ctx context.Context, key string, data []byte) (err error) {
		fmt.Println(key)
		fmt.Println(string(data))
		return nil
	})
	cacheMock.GetMock.Return(nil, cache.ErrCacheMiss)

	storeCertificate(ctx, cacheMock, "asd", cert)
}

func TestIsNeedRenew(t *testing.T) {
	td := testdeep.NewT(t)
	var cert = &tls.Certificate{}
	cert.Leaf = &x509.Certificate{NotAfter: time.Date(2000, 7, 31, 0, 0, 0, 0, time.UTC)}
	td.True(isNeedRenew(cert, time.Date(2000, 7, 31, 0, 0, 0, 1, time.UTC)))
	td.True(isNeedRenew(cert, time.Date(2000, 7, 1, 0, 0, 0, 1, time.UTC)))
	td.False(isNeedRenew(cert, time.Date(2000, 7, 1, 0, 0, 0, 0, time.UTC)))
	td.False(isNeedRenew(cert, time.Date(2000, 6, 30, 0, 0, 0, 0, time.UTC)))
}

type testManagerContext struct {
	ctx context.Context

	manager                *Manager
	connContext            contextConnection
	conn                   *ConnMock
	cache                  *CacheMock
	certForDomainAuthorize *ValueMock
	certState              *ValueMock
	client                 *AcmeClientMock
	domainChecker          *DomainCheckerMock
	httpTokens             *CacheMock
}

func TestManager_CertForLockedDomain(t *testing.T) {
	td := testdeep.NewT(t)
	c, cancel := createManager(t)
	defer cancel()

	c.certState.GetMock.Return(&certState{}, nil)
	c.cache.GetMock.Set(func(ctx context.Context, key string) (ba1 []byte, err error) {
		if key == "test.ru.lock" {
			return []byte{}, nil
		}
		return nil, cache.ErrCacheMiss
	})

	res, err := c.manager.GetCertificate(&tls.ClientHelloInfo{Conn: c.connContext, ServerName: "test.ru"})
	td.Nil(res)
	td.CmpError(err)
}

func TestManager_CertForDenied(t *testing.T) {
	td := testdeep.NewT(t)
	c, cancel := createManager(t)
	defer cancel()

	c.certState.GetMock.Return(&certState{}, nil)
	c.cache.GetMock.Return(nil, cache.ErrCacheMiss)
	c.domainChecker.IsDomainAllowedMock.Return(false, nil)

	res, err := c.manager.GetCertificate(&tls.ClientHelloInfo{Conn: c.connContext, ServerName: "test.ru"})
	td.Nil(res)
	td.CmpError(err)
}

func createManager(t *testing.T) (res testManagerContext, cancel func()) {
	ctx, ctxCancel := th.TestContext()
	mc := minimock.NewController(t)

	res.ctx = ctx
	res.conn = NewConnMock(mc)
	res.connContext = contextConnection{
		Conn:    res.conn,
		Context: zc.WithLogger(context.Background(), zap.NewNop()),
	}
	res.cache = NewCacheMock(mc)
	res.client = NewAcmeClientMock(mc)
	res.certForDomainAuthorize = NewValueMock(mc)
	res.certState = NewValueMock(mc)
	res.domainChecker = NewDomainCheckerMock(mc)
	res.httpTokens = NewCacheMock(mc)

	res.manager = &Manager{
		CertificateIssueTimeout: time.Second,
		Cache:                   res.cache,
		Client:                  res.client,
		DomainChecker:           res.domainChecker,
		EnableHTTPValidation:    true,
		EnableTLSValidation:     true,
		certForDomainAuthorize:  res.certForDomainAuthorize,
		certState:               res.certState,
		httpTokens:              res.httpTokens,
	}
	return res, func() {
		mc.Finish()
		ctxCancel()
	}
}
