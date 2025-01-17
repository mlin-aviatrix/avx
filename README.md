`avx` Aviatrix CLI
=

Disclaimer: This is a personal project and is in no way affiliated with Aviatrix Systems Inc.

`avx` is a CLI for interacting with the Aviatrix RPC API https://api.aviatrix.com/

This is done through the Aviatrix SDK that is maintained for use in the Aviatrix
terraform provider https://github.com/AviatrixSystems/terraform-provider-aviatrix

Build & Install
-

### Install via Homebrew on MacOS
```shell script
$ brew tap CyrusJavan/tap
$ brew install avx
```

### Build from source on MacOS
```shell script
$ git clone git@github.com:CyrusJavan/avx.git
$ cd avx
$ go build -o /usr/local/bin/avx
$ chmod +x /usr/local/bin/avx
```

Required Configuration
-

Required environment variables:

- `AVIATRIX_CONTROLLER_IP`
- `AVIATRIX_USERNAME`
- `AVIATRIX_PASSWORD`

Usage
-

### `avx login`

`avx login` will attempt to login with the provided credentials. If
successful, the CID will be printed out.
```shell script
$ avx login
CID: "MMUyqYcNOjaWUWIFHmYA"
```

---

### `avx api <method> <path>`

`avx api <method> <path>` will attempt to login and send an HTTP request based on the
provided method `method` to the provided API `path`. `avx` prints out debug information
like the controller IP, request URL, request body and response latency. `avx` then prints out
the response status code and body.
```shell script
$ avx api GET app-domains
controller IP: 127.0.0.1
request url: https://127.0.0.1/v2.5/api/app-domains
request body:
null
latency: 11ms
response status code: 200
response body:
{
  "app_domains": [
    {
      ...
```

---

### `avx rpc <action>`

`avx rpc <action>` will attempt to login and send a POST request to
the API with the provided `action`. `avx` prints out debug information like the
controller IP, request body and response latency. `avx` then prints out the 
response body.
```shell script
$ avx rpc list_accounts
controller IP: 127.0.0.1
request body:
{
  "CID": "soPEtEopZlkC1Vwwzdl4",
  "action": "list_accounts"
}
latency: 153ms
response body:
{
  "return": true,
  "results": {
    "account_list": [
      {
        ...
```

---

### `avx rpc <action> <key>=<value> [<key>=<value>...]`

In this form `avx rpc` will send a POST request with the given action and any extra
params that were provided.
```shell script
$ avx rpc delete_account_profile account_name=john-gcloud
controller IP: 127.0.0.1
request body:
{
  "CID": "CgRVzRukvCtUGLwp80lw",
  "account_name": "tfa-byl0f",
  "action": "delete_account_profile"
}
latency: 113ms
response body:
{
  "return": true,
  "results": "Account deleted successfully."
} 
```

---

### `avx export <resource_name>`

Exports the Terraform configuration for the provided resource name.
```shell script
$ avx export transit_gateway
resource "aviatrix_transit_gateway" "transit_gateway_1" {
    gw_name = "transit-gateway-bezukhov"
    vpc_id = "vpc-0a00b42135b095d6c"
    cloud_type = 1
    vpc_reg = "us-east-1"
    connected_transit = true
    enable_active_mesh = true
    gw_size = "c4.4xlarge"
    account_name = "transit-gateway-acc-bezukhov"
    enable_hybrid_connection = true
    subnet = "10.0.0.64/28"
    learned_cidrs_approval_mode = "connection"
}
```

---
