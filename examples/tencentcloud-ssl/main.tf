resource "cloud_ssl_certificate" "ca" {
  name = "ssl-ca"
  type = "CA"
  cert = var.ca
}

resource "cloud_ssl_certificate" "svr" {
  name = "ssl-svr"
  type = "SVR"
  cert = var.cert
  key  = var.key
}

data "cloud_ssl_certificates" "ca" {
  name = cloud_ssl_certificate.ca.name
}

data "cloud_ssl_certificates" "svr" {
  type = cloud_ssl_certificate.svr.type
}