package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	pg_query "github.com/lfittl/pg_query_go"
	"github.com/rueian/pgbroker/backend"
	"github.com/rueian/pgbroker/message"
	"github.com/rueian/pgbroker/proxy"
)

func sanitizeQueryStr(queryString string) string {
	sanitizedQueryStr := strings.Replace(queryString, "\n", " ", -1)
	sanitizedQueryStr = strings.Replace(sanitizedQueryStr, "\r", "", -1)
	return sanitizedQueryStr
}

func generateErrorQuery(msg string) string {
	// return a query that triggers an error on the server that contains our
	// desired error message
	return fmt.Sprintf("'%s';", strings.Replace(msg, "'", "''", -1))
}

func main() {
	inShutdown := false

	var configPath string
	if len(os.Args) == 2 {
		configPath = os.Args[1]
	} else {
		configPath = "config.yaml"
	}

	if err := loadConfig(configPath); err != nil {
		panic(err)
	}

	ln, err := net.Listen("tcp", GetConfig().Listen)
	if err != nil {
		panic(err)
	}

	clientMessageHandlers := proxy.NewClientMessageHandlers()
	clientMessageHandlers.AddHandleQuery(func(ctx *proxy.Ctx, msg *message.Query) (query *message.Query, e error) {
		fingerprint, err := pg_query.FastFingerprint(msg.QueryString)
		if err != nil {
			fmt.Printf("failed to parse query: %v: %s\n", err, sanitizeQueryStr(msg.QueryString))

			msg.QueryString = generateErrorQuery(fmt.Sprintf("failed to parse query: %v", err))
			return msg, nil
		}

		_, ok := GetConfig().AllowedFingerprints[fingerprint]
		if !ok {
			fmt.Printf("query with finterprint %s not allowed: %s\n", fingerprint, sanitizeQueryStr(msg.QueryString))

			msg.QueryString = generateErrorQuery("query is not allowed")
			return msg, nil
		}

		return msg, nil
	})
	clientMessageHandlers.AddHandleClientOther(func(ctx *proxy.Ctx, msg *message.Raw) (raw *message.Raw, e error) {
		_, ok := GetConfig().AllowedCommands[msg.Type]
		if !ok {
			return nil, fmt.Errorf("disallowed client command %c", msg.Type)
		}
		return msg, nil
	})

	serverStreamCallbackFactories := proxy.NewStreamCallbackFactories()

	server := &proxy.Server{
		PGResolver:                    backend.NewStaticPGResolver(GetConfig().TargetServer),
		ConnInfoStore:                 backend.NewInMemoryConnInfoStore(),
		ClientMessageHandlers:         clientMessageHandlers,
		ServerStreamCallbackFactories: serverStreamCallbackFactories,
		OnHandleConnError: func(err error, ctx *proxy.Ctx, conn net.Conn) {
			if err == io.EOF {
				return
			}

			fmt.Println("OnHandleConnError", err)
		},
	}

	go func() {
		if err := server.Serve(ln); err != nil {
			if !inShutdown {
				panic(err)
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	//<-sigs
	for {
		signal := <-sigs
		if signal == syscall.SIGHUP {
			if err := loadConfig(configPath); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("reloaded config")
			}
		} else {
			break
		}
	}
	inShutdown = true
	server.Shutdown()
}
