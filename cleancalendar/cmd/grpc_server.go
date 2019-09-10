package cmd

import (
	"github.com/otusteam/go/cleancalendar/internal/domain/services"
	"github.com/otusteam/go/cleancalendar/internal/grpc/api"
	"github.com/otusteam/go/cleancalendar/internal/maindb"
	"github.com/spf13/cobra"
	"log"
)

// TODO: dependency injection, orchestrator
func construct(dsn string) (*api.CalendarServer, error) {
	eventStorage, err := maindb.NewPgEventStorage(dsn)
	if err != nil {
		return nil, err
	}
	eventService := &services.EventService{
		EventStorage: eventStorage,
	}
	server := &api.CalendarServer{
		EventService: eventService,
	}
	return server, nil
}

var addr string
var dsn string

var GrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "Run grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := construct(dsn)
		if err != nil {
			log.Fatal(err)
		}
		err = server.Serve(addr)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	GrpcServerCmd.Flags().StringVar(&addr, "addr", "localhost:8080", "host:port to listen")
	GrpcServerCmd.Flags().StringVar(&dsn, "dsn", "host=127.0.0.1 user=event_user password=event_pwd dbname=event_db", "database connection string")
}
