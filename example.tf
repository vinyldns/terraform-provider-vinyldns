provider "vinyldns" {
  host       = "https://dev-api.vinyldns.comcast.net:9443"
  access_key = "ptHljFIk44mCsyHimys3"
  secret_key = "3GxCeAviZAJiJ0E7KtVE"
}

resource "vinyldns_group" "test_group" {
  name       = "terraform-provider-test-group"
  member_ids = ["123"]
  admin_ids  = ["123"]
}

resource "vinyldns_zone" "test_zone" {
  name           = "system-test."
  email          = "foo@bar.com"
  admin_group_id = "${vinyldns_group.test_group.id}"

  zone_connection {
    name           = "vinyldns."
    key_name       = "vinyldns."
    key            = "123"
    primary_server = "127.0.0.1"
  }
}

resource "vinyldns_record_set" "test_record_set" {
  name             = "terraformtestrecordset"
  zone_id          = "${vinyldns_zone.test_zone.id}"
  type             = "A"
  ttl              = 6000
  record_addresses = ["127.0.0.1"]
}

resource "vinyldns_record_set" "another_test_record_set" {
  name         = "another-terraformtestrecordset"
  zone_id      = "${vinyldns_zone.test_zone.id}"
  type         = "CNAME"
  ttl          = 6000
  record_cname = "foo-bar.com."
}

data "vinyldns_group" "test_group_datasource" {
  id = vinyldns_group.test_group.id
}

output "name" {
  value = data.vinyldns_group.test_group_datasource.name
}

output "email" {
  value = data.vinyldns_group.test_group_datasource.email
}

output "description" {
  value = data.vinyldns_group.test_group_datasource.description
}

data "vinyldns_record_set" "test_recordset_datasource" {
  zoneid   = "zoneid"
}
 output "recordset"{
  value = data.vinyldns_record_set.test_recordset_datasource.recordset
 }