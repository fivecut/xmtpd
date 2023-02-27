package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jessevdk/go-flags"
	"github.com/xmtp/xmtpd/pkg/crdt"
	memstore "github.com/xmtp/xmtpd/pkg/crdt/stores/mem"
	"github.com/xmtp/xmtpd/pkg/node"
	postgresstore "github.com/xmtp/xmtpd/pkg/store/postgres"
	"github.com/xmtp/xmtpd/pkg/zap"
)

// GitCommit should be included in the binary via -ldflags=-X ${COMMIT}
var GitCommit string

var opts node.Options

func main() {
	// Initialize options.
	_, err := flags.Parse(&opts)
	if err != nil {
		if err, ok := err.(*flags.Error); !ok || err.Type != flags.ErrHelp {
			fatal("error parsing options: %s", err)
		}
		return
	}

	ctx := context.Background()

	// Initialize logger.
	log, err := zap.NewLogger(&opts.Log)
	if err != nil {
		fatal("error building logger: %s", err)
	}
	log.Info("running", zap.String("git-commit", GitCommit))

	// Initialize datastore.
	var storeMaker node.StoreMakerFunc
	if opts.Store.Postgres.DSN != "" {
		log.Info("using postgres store")
		db, err := postgresstore.NewDB(opts.Store.Postgres.DSN)
		if err != nil {
			fatal("error initializing postgres: %s", err)
		}
		store, err := postgresstore.New(ctx, log, db)
		if err != nil {
			fatal("error initializing postgres: %s", err)
		}
		storeMaker = func(topic string) (crdt.Store, error) {
			return store.Scoped(topic), nil
		}
	} else {
		log.Info("using memory store")
		storeMaker = func(topic string) (crdt.Store, error) {
			return memstore.New(log), nil
		}
	}

	// Initialize node.
	node, err := node.New(ctx, log, storeMaker, &opts)
	if err != nil {
		log.Fatal("error initializing node", zap.Error(err))
	}
	defer node.Close()

	// Wait for shutdown.
	sig := waitForEndSignal()
	log.Info("ending", zap.String("signal", sig.String()))
}

func waitForEndSignal() os.Signal {
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	return <-sigC
}

func fatal(msg string, args ...any) {
	log.Fatalf(msg, args...)
}
