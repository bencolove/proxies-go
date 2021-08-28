package httproxies

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

func ProxiedClient(proxyURL string) (*http.Client, error) {
	tr := http.DefaultTransport.(*http.Transport).Clone()

	if proxyUrl, err := url.Parse(proxyURL); err != nil {
		return nil, err
	} else {
		/* Transport tls config
		// Load client cert
		cert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
		if err != nil {
		return &tlsConfig, err
		}
		tlsConfig.Certificates = []tls.Certificate{cert}

		// Load CA cert
		caCert, err := ioutil.ReadFile(caCertFile)
		if err != nil {
		return &tlsConfig, err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig.RootCAs = caCertPool

		tlsConfig.BuildNameToCertificate()
		*/
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		tr.Proxy = http.ProxyURL(proxyUrl)

		return &http.Client{Transport: tr}, nil
	}
}
