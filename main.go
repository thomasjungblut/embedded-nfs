package main

import (
	"github.com/smallfz/libnfs-go/backend"
	"github.com/smallfz/libnfs-go/server"
	"github.com/thomasjungblut/embeded-nfs/pkg/simpledbfs"
	"log"
	"os"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current directory: %v", err)
		return
	}

	dbfs, err := simpledbfs.NewSimpleDBFS(workingDir)
	if err != nil {
		log.Fatalf("simpledbfs.NewSimpleDBFS: %v", err)
	}

	svr, err := server.NewServerTCP(":2049", backend.New(dbfs))
	if err != nil {
		log.Fatalf("server.NewServerTCP: %v", err)
		return
	}

	if err := svr.Serve(); err != nil {
		log.Fatalf("svr.Serve: %v", err)
	}
}
