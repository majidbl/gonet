package main

import (
    "flag"
    "fmt"
    "net"
)

func main() {
    // Register Int flag.
    port := flag.String("port", "", "check specific port")
    url := flag.String("url", "", "LookupIP for specific url")
    ipValue := flag.String("ip", "", "LookupAddr for specific ip address")
    ns := flag.String("ns", "", "LookupNS for specific ip address")
    // Parse the flags.
    flag.Parse()

    // Print the argument.
    //fmt.Println("Argument", *port)

    // Get int from the Int pointer.
    //value := *port
    //fmt.Println("Value", value)
    //fmt.Println(flag.Args())
    if *port != ""{
      _, err := net.Listen("tcp", ":" + *port)
      if err != nil {
        // Log or report the error here
        fmt.Printf("Port %s unavailable!\n", *port)
      }else{
        fmt.Printf("Port %s available!\n", *port)
      }
      //defer l.Close()
    }
    if *url != ""{
      iprecords, _ := net.LookupIP(*url)
      for _, ip := range iprecords {
        fmt.Println(ip)
      }
    }
    if *ipValue != ""{
      ptr, _ := net.LookupAddr(*ipValue)
      for _, ptrvalue := range ptr {
        fmt.Println(ptrvalue)
      }
    }
    if *ns != "" {
      nameserver, _ := net.LookupNS(*ns)
      for _, ns := range nameserver {
        fmt.Println(ns)
        
      }
    }
}
