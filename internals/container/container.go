package container

import (
	"fmt"

	"go-boilerplate/internals/controller"
	"go-boilerplate/internals/infrastructure/database"
	grpcserver "go-boilerplate/internals/infrastructure/grpc_server"
	"go-boilerplate/internals/repository"
	"go-boilerplate/internals/service"
	"go-boilerplate/internals/utils"

	"go-boilerplate/internals/config"

	"go.uber.org/dig"
)

// Container ...
type Container struct {
	container *dig.Container
}

// Configure ...
func (c *Container) Configure() error {
	// config
	if err := c.container.Provide(config.NewConfiguration); err != nil {
		return err
	}

	// server
	if err := c.container.Provide(grpcserver.NewServer); err != nil {
		return err
	}

	// controller
	if err := c.container.Provide(controller.NewHealthZController); err != nil {
		return err
	}
	if err := c.container.Provide(controller.NewPingPongController); err != nil {
		return err
	}

	// database
	if err := c.container.Provide(database.NewServerBase); err != nil {
		return err
	}

	// repository
	if err := c.container.Provide(repository.NewRepository); err != nil {
		return err
	}

	// service
	if err := c.container.Provide(service.NewService); err != nil {
		return err
	}

	// utils
	if err := c.container.Provide(utils.NewUtils); err != nil {
		return err
	}
	if err := c.container.Provide(utils.NewCustomValidator); err != nil {
		return err
	}

	// If have new dependency should be set here
	return nil
}

// Start ...
func (c *Container) Start() error {
	fmt.Println("Start Container")
	if err := c.container.Invoke(func(s *grpcserver.Server) {
		s.Start()
	}); err != nil {
		fmt.Printf("%s", err)
		return err
	}
	return nil
}

// MigrateDB ...
// func (c *Container) MigrateDB() error {
// 	fmt.Println("Start Container DB")
// 	if err := c.container.Invoke(func(d *database.DB) {
// 		d.MigrateDB()
// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

// // SeederDb ...
// func (c *Container) SeederDb() error {
// 	if err := c.container.Invoke(func(d *database.DB) {
// 		d.Seeder()
// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

// NewContainer ...
func NewContainer() (*Container, error) {
	fmt.Println("this file should be set dependency injection by using uber-dig")
	d := dig.New()
	container := &Container{
		container: d,
	}
	if err := container.Configure(); err != nil {
		return nil, err
	}
	return container, nil
}
