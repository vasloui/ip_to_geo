## IP Geolocation CLI
A minimal Go command-line tool that looks up the geolocation details for an IPv4 or IPv6 address using the IPGeolocation.io API.

It reads an API key from a local .env file and prints the raw JSON response returned by the API.

Features
Query IP geolocation data by IP address.
Loads the API key from a local .env file:
Inside:
```text
api_key=<...>
```

Uses the IPGeolocation.io v3 ipgeo endpoint.
Prints the complete JSON response.
https://ipgeolocation.io/documentation/ip-location-api.html

Simple CLI usage: 
Download dependancies (if any):
```
go mod tidy
```

```
go run . -i <ip-address>.
```

Or:
```
go build
.\ip_to_geo.exe -i <IP>
```

