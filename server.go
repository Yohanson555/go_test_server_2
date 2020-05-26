package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/valyala/fasthttp"
)

const DefaultPort = 8080

func main() {
	// The server will listen for incoming requests on this address.

	args := mapArgs(os.Args)

	var port int64 = 8080

	if p, e := getServicePort(args); e == nil {
		port = p
	}

	// This function will be called by the server for each incoming request.
	//
	// RequestCtx provides a lot of functionality related to http request
	// processing. See RequestCtx docs for details.
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "I am FasttHTTP server: \n")
		fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
		fmt.Fprintf(ctx, "Addr is %q\n", ctx.LocalAddr())
		fmt.Fprintf(ctx, "IP is %q\n", ctx.LocalIP())

		//fmt.Fprintf(ctx, "Hello, world! Requested path is %q", ctx.Path())
	}

	// Start the server with default settings.
	// Create Server instance for adjusting server settings.
	//
	// ListenAndServe returns only on error, so usually it blocks forever.
	if err := fasthttp.ListenAndServe(fmt.Sprintf(":%v", port), requestHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}

func mapArgs(args []string) map[string]string {
	var p interface{} = nil
	mappedArgs := map[string]string{}

	for _, s := range args {
		if p == nil {
			p = s
		} else {
			mappedArgs[p.(string)] = s
			p = nil
		}
	}

	return mappedArgs
}

func getServicePort(args map[string]string) (int64, error) {
	var port int64

	if val, ok := args["-p"]; ok {
		if i, err := strconv.ParseInt(val, 10, 64); err == nil {
			if i > 0 && i < 65535 {
				port = i
			} else {
				return 0, fmt.Errorf("port value should be greater than 0 and less than 65535")
			}
		} else {
			return 0, err
		}
	} else {
		port = DefaultPort
	}

	return port, nil
}
