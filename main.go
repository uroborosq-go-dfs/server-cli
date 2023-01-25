package main

import (
	"fmt"
	guid "github.com/google/uuid"
	"github.com/uroborosq-go-dfs/server/connector"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
	"github.com/uroborosq-go-dfs/server/server"
)

var serv *server.Server

func main() {
	host := os.Getenv("GODFS_DB_HOST")
	port := os.Getenv("GODFS_DB_PORT")
	user := os.Getenv("GODFS_DB_USER")
	password := os.Getenv("GODFS_DB_PASSWORD")
	dbname := os.Getenv("GODFS_DB_DBNAME")
	
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	serv, err := server.CreateServer("pgx", conn)
	if err != nil {
		log.Fatal(err.Error())
	}
	app := &cli.App{
		Name:    "UroborosQ's simple distributed file system - server",
		Usage:   "If you want to, you can use it",
		Version: "v0.0.1",
		Commands: []*cli.Command{
			{
				Name:     "add-file",
				Usage:    "Add file to the volume",
				Category: "file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "source",
						Usage: "Path to your local file",
					},
					&cli.StringFlag{
						Name:  "target",
						Usage: "Partial path on the volume.",
					},
				},
				Action: func(context *cli.Context) error {
					return serv.AddFile(context.Args().First(), context.Args().Get(1))
				},
			},
			{
				Name:     "remove-file",
				Usage:    "Remove file from the volume",
				Category: "file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "target",
						Usage: "Partial path on the volume",
					},
				},
				Action: func(context *cli.Context) error {
					return serv.RemoveFile(context.Args().First())
				},
			},
			{
				Name:     "get-file",
				Usage:    "Download file to given path",
				Category: "file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "source",
						Usage: "Partial path on the volume",
					},
					&cli.StringFlag{
						Name:  "target",
						Usage: "Path to download on the host",
					},
				},
				Action: func(context *cli.Context) error {
					return serv.GetFile(context.Args().Get(0), context.Args().Get(1))
				},
			},
			{
				Name:     "add-node",
				Usage:    "Add node to the system",
				Category: "node",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "ip",
						Usage: "ip address of the node",
					},
					&cli.StringFlag{
						Name:  "port",
						Usage: "port of the node",
					},
					&cli.Int64Flag{
						Name:  "size",
						Usage: "Max amount of bytes, which client can use on the node",
					},
					&cli.StringFlag{
						Name:  "type",
						Usage: "Type of protocol which will be used to communicate with node",
					},
				},
				Action: func(context *cli.Context) error {
					connectionTypeInt := -1
					connectionType := context.Args().Get(3)
					if connectionType == "tcp" {
						connectionTypeInt = 1
					} else if connectionType == "http" {
						connectionTypeInt = 2
					}
					maxSize, err := strconv.ParseInt(context.Args().Get(2), 0, 64)
					if err != nil {
						return err
					}
					id, err := serv.AddNode(context.Args().Get(0), context.Args().Get(1), maxSize, connector.NetConnectorType(connectionTypeInt))
					if err != nil {
						return err
					}
					fmt.Printf("Node added! Its id is %s", id.String())
					return nil
				},
			},
			{
				Name:     "remove-node",
				Usage:    "Remove node from the system",
				Category: "node",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "node",
						Usage: "uuid of node to remove",
					},
				},
				Action: func(context *cli.Context) error {
					id, err := guid.Parse(context.Args().First())
					if err != nil {
						return err
					}
					return serv.RemoveNode(id)
				},
			},
			{
				Name:     "clean-node",
				Usage:    "Delete all files from node",
				Category: "node",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "node",
						Usage: "uuid of node to be cleaned",
					},
				},
				Action: func(context *cli.Context) error {
					id, err := guid.Parse(context.Args().First())
					if err != nil {
						return err
					}
					return serv.CleanNode(id)
				},
			},
			{
				Name:     "node-list",
				Usage:    "Show all files on this node",
				Category: "node",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "node",
						Usage: "uuid of node to show files on",
					},
				},
				Action: func(context *cli.Context) error {
					id, err := guid.Parse(context.Args().First())
					if err != nil {
						return err
					}
					paths, sizes, err := serv.ListOfNodeFiles(id)
					if err != nil {
						return err
					}
					fmt.Println("Files on node uuid " + id.String() + ":")
					for i := 0; i < len(paths); i++ {
						fmt.Printf("%s - %d b\n", paths[i], sizes[i])
					}
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "Show all files",
				Action: func(context *cli.Context) error {
					paths, sizes, err := serv.ListOfAllFiles()
					if err != nil {
						return err
					}
					fmt.Println("All files:")
					for i := 0; i < len(paths); i++ {
						fmt.Printf("%s - %d b\n", paths[i], sizes[i])
					}
					return nil
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Welcome to the distributed file system client!")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
