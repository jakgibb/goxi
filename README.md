Go XI
=========

Go XI is a Golang library for interacting with the Nagios XI REST API allowing retrieval of information of hosts and services within the system

Usage
-----
Clone the repository or use `go get`:

    $ go get https://github.com/jakgibb/goxi

Import the library into your project:

```go
import "github.com/jakgibb/goxi"
```

Examples
--------
The library allows retreiving information of any host or service; optional filters can be configured to only query a subset of data

Example to retrieve informaion on specific hosts:
```go
// Create a client object passing into the URL of the Nagios instance and API key
// The client is stateless and so can be reused
client := goxi.NewClient(
	"https://nagios.example/nagiosxi",
	"apiKeyFromUserDashboard",
)

// Set up a Host Filter to retrieve information for hosts: server-a and server-b
// See filter.go for a list of all available filters available
hostFilter := goxi.HostFilter{
	Name:    []string{"server-a", "server-b"},
}

// Pass the filter to GetHosts which will make a API call and retrieve all information for the two hosts
// and return a slice of Host structs and an error value (nil for no error)
hosts, err := client.GetHosts(hostFilter)

if err != nil {
	fmt.Println(err)
}

// Range over the results and print the name of the host, address and alias
// See host.go for a list of all available fields that can be accessed
for _, v := range *hosts {
	fmt.Printf("Host: %s - Address: %s - Alias: %s", v.Name, v.Address, v.Alias)
}
````
Example to retrieve information on specific services:
````go
// Retrieve information for all services which are assigned the 'Linux Disk space' check
// Limit the number of results returned to 10
serviceFilter := goxi.ServiceFilter{
	Name:    []string{"Linux Disk Space"},
	Records: "10",
}

services, err := client.GetServices(serviceFilter)

if err != nil {
	fmt.Println(err)
}

for _, v := range *services {
	fmt.Printf("Service: %s - Check Command: %s - Status Output: %s", v.Name, v.CheckCommand, v.StatusText)
}
````

To retireve information on all hosts or services, pass in an empty filter
````go
// To return the status information for all hosts or services, pass an empty filter
serviceFilter := goxi.ServiceFilter{}
hostFilter := goxi.HostFilter{}
````

More information on the Nagios API can be found within the 'Help' section of the dashboard

License
-------
```
GNU GENERAL PUBLIC LICENSE
Version 3, 29 June 2007
```
