provider "vinyldns" {
  host       = "https://dev-api.vinyldns.comcast.net:9443"
  access_key = "ptHljFIk44mCsyHimys3"
  secret_key = "3GxCeAviZAJiJ0E7KtVE"
}
resource "vinyldns_group" "test_groups" {
  name       = "terraform-provider-group4"
  member_ids = ["6e07183f-a68e-42cf-acce-044296ede753"]
  admin_ids  = ["6e07183f-a68e-42cf-acce-044296ede753"]
  email      = "foo@bar.com"
}

data "vinyldns_group" "current_group" {
  id= vinyldns_group.test_groups.id
}
output "name" {
  value = data.vinyldns_group.current_group.name
}

output "email" {
  value = data.vinyldns_group.current_group.email
}

output "description" {
  value = data.vinyldns_group.current_group.description
}