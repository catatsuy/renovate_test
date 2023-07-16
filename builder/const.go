package builder

// nginx
const (
	NginxVersion           = "1.22.0"
	NginxDownloadURLPrefix = "https://nginx.org/download"
)

// pcre
const (
	// datasource=github-tags depName=PCRE2Project/pcre2 versioning=regex:^(?<compatibility>.*)-(?<version>[0-9.]+)$
	PcreVersion           = "10.40"
	PcreDownloadURLPrefix = "https://github.com/PhilipHazel/pcre2/releases/download"
)

// openssl
const (
	OpenSSLVersion           = "1.1.1o"
	OpenSSLDownloadURLPrefix = "https://www.openssl.org/source"
)

// libressl
const (
	LibreSSLVersion           = "3.5.3"
	LibreSSLDownloadURLPrefix = "https://ftp.openbsd.org/pub/OpenBSD/LibreSSL"
)

// datasource=github-tags depName=libressl/portable
 "v3.5.3"

// datasource=github-tags depName=libressl/portable
 "3.5.3"

// datasource=golang-version depName=golang
 "1.17.1"

// zlib
const (
	// datasource=github-tags depName=madler/zlib
	ZlibVersion           = "1.2.12"
	ZlibDownloadURLPrefix = "https://zlib.net"
)

// openResty
const (
	OpenRestyVersion           = "1.21.4.1"
	OpenRestyDownloadURLPrefix = "https://openresty.org/download"
)

// tengine
const (
	TengineVersion           = "2.3.3"
	TengineDownloadURLPrefix = "https://tengine.taobao.org/download"
)

// component enumerations
const (
	ComponentNginx = iota
	ComponentOpenResty
	ComponentTengine
	ComponentPcre
	ComponentOpenSSL
	ComponentLibreSSL
	ComponentZlib
	ComponentMax
)
