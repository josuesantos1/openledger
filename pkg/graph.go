package pkg

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Graph interface {
	Connect() error
	Close(ctx context.Context) error
	CreateNode(data interface{}) (string, error)
}

type neo4jGraph struct {
	uri      string
	username string
	password string
	driver   neo4j.DriverWithContext
}

func NewGraph(uri, username, password string) *neo4jGraph {
	return &neo4jGraph{
		uri:      uri,
		username: username,
		password: password,
	}
}

func (g *neo4jGraph) Connect() error {
	driver, err := neo4j.NewDriverWithContext(g.uri, neo4j.BasicAuth(g.username, g.password, ""))
	if err != nil {
		return err
	}

	g.driver = driver
	return nil
}

func (g *neo4jGraph) Close(ctx context.Context) error {
	if g.driver != nil {
		return g.driver.Close(ctx)
	}

	return nil
}
