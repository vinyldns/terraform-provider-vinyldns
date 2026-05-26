# Configure the VinylDNS provider
#
# Authentication can be provided via environment variables:
#   VINYLDNS_HOST       - VinylDNS API endpoint
#   VINYLDNS_ACCESS_KEY - Your access key
#   VINYLDNS_SECRET_KEY - Your secret key

provider "vinyldns" {
  # Explicit configuration (environment variables recommended for credentials)
  host       = "https://vinyldns.example.com"
  access_key = var.vinyldns_access_key
  secret_key = var.vinyldns_secret_key
}

variable "vinyldns_access_key" {
  description = "VinylDNS access key"
  type        = string
  sensitive   = true
}

variable "vinyldns_secret_key" {
  description = "VinylDNS secret key"
  type        = string
  sensitive   = true
}
