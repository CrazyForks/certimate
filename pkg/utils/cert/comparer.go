package cert

import (
	"crypto/x509"
)

// 比较两个 x509.Certificate 对象，判断它们是否是同一张证书。
//
// 入参:
//   - a: 待比较的第一个 x509.Certificate 对象。
//   - b: 待比较的第二个 x509.Certificate 对象。
//
// 出参:
//   - 是否相同。
func EqualCertificates(a, b *x509.Certificate) bool {
	if a == nil || b == nil {
		return false
	}

	return a.Equal(b)
}

// 与 [EqualCertificates] 方法类似，但入参是 PEM 编码的证书字符串。
//
// 入参:
//   - a: 待比较的第一个证书 PEM 内容。
//   - b: 待比较的第二个证书 PEM 内容。
//
// 出参:
//   - 是否相同。
func EqualCertificatesFromPEM(a, b string) bool {
	aCert, _ := ParseCertificateFromPEM(a)
	bCert, _ := ParseCertificateFromPEM(b)
	return EqualCertificates(aCert, bCert)
}
