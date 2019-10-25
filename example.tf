data "vinyldns_group" "test_group" {
  id = "testgroupid"
}

resource "vinyldns_zone" "test_zone" {
  name           = "system-test."
  email          = "foo@bar.com"
  admin_group_id = "${data.vinyldns_group.test_group.id}"

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
