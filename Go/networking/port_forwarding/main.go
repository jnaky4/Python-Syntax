package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	//"net"
	"os"
)

//p 8080:80
//Map TCP port 80 in the container to port 8080 on the Docker host.
//p 192.168.1.100:8080:80
//Map TCP port 80 in the container to port 8080 on the Docker host for connections to host IP 192.168.1.100.
//p 8080:80/udp
//Map UDP port 80 in the container to port 8080 on the Docker host.
//p 8080:80/tcp -p 8080:80/udp
//Map TCP port 80 in the container to TCP port 8080 on the Docker host,
//and map UDP port 80 in the container to UDP port 8080 on the Docker host.
func main() {
	fmt.Println(os.Args[2])
	if len(os.Args) < 3 {
		fmt.Println("Too Few Arguments")
	} else {
		pPtr := flag.String("p", "p", "port forwarding")
		flag.Parse()
		fmt.Println("Port PTR: ", *pPtr)
		args := flag.Args()
		fmt.Println("args: ", args)

		portArg := args[1]
		proto := "tcp"

		//pro, por := nat.SplitProtoPort("1234/tcp")
		//println("proto: ", pro, "por: ", por)
		//
		//portMappings, err := nat.ParsePortSpec("0.0.0.0:1234-1235:3333-3334/tcp")
		//if err != nil{
		//	println("ERR: ", err)
		//}

		//fmt.Printf("PortMapping: %+v ->", portMappings)
		//println("PortMapping: ", portMappings)
		//println("bindings: ", bindings)


		parts := strings.SplitN(portArg, "/", 2)
		if len(parts) == 2 && len(parts[1]) != 0 {
			portArg = parts[0]
			proto = parts[1]
		}

		//natPort := portArg + "/" + proto
		//newP, err := nat.NewPort(proto, portArg)
		//if err != nil {
		//	println("ERR: ", err)
		//}

		//if frontends, exists := c.NetworkSettings.Ports[newP]; exists && frontends != nil {
		//	for _, frontend := range frontends {
		//		fmt.Fprintln(dockerCli.Out(), net.JoinHostPort(frontend.HostIP, frontend.HostPort))
		//	}
		//}
		//println("Error: No public port '%s' published for %s", natPort, opts.container)


		ports := strings.Split(portArg, ":")

		argsIndex := 0

		var localIP string
		var port int
		var portForward int

		if len(ports) == 3 {
			//Local IP specified
			localIP = ports[0]
			argsIndex = 1
		}
		if i, err := strconv.Atoi(ports[argsIndex]); err == nil {
			port = i
		} else {
			println(err)
		}

		if i, err := strconv.Atoi(ports[argsIndex+1]); err == nil {
			portForward = i
		} else {
			println(err)

			//newP, err := nat.NewPort(proto, portArg)
			//if err != nil {
			//	//return err
			//	fmt.Println("ERR: ", err)
			//}

			fmt.Println("LocalIp: ", localIP)
			fmt.Println("Port: ", port)
			fmt.Println("PortForward: ", portForward)
			fmt.Println("Connection: ", proto)

			//fmt.Println("tail: ", flag.Args())
		}
	}
}
