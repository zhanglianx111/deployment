package main

import (
  "os"
  "fmt"
  "net/http"
  "net"
  "log"
)

func main() {
    count := 0
    addrsList := []string{}
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
		addrsList = append(addrsList, ipnet.IP.String())
            }
        }
    }

    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	count++
        fmt.Fprintf(os.Stdout, "I'm %s, ip: %v, count: %d\n", hostname, addrsList, count)
 	fmt.Fprintf(w, "I'm %s, ip: %v, count: %d\n", hostname, addrsList, count)
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}
