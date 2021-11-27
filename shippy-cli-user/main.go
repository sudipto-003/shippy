package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	pb "github.com/sudipto-003/shippy/shippy-service-user/proto/user"
)

func createUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService("shippy.service.user", service.Client())
	users, err := client.GetAll(ctx, &pb.Request{})
	if err != nil {
		return err
	}
	for _, v := range users.Users {
		log.Println(v)
	}

	rsp, err := client.Create(ctx, user)

	if err != nil {
		return err
	}

	fmt.Println("Response: ", rsp.User)
	return nil
}

func authUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService("shippy.service.user", service.Client())
	token, err := client.Auth(ctx, user)

	if err != nil {
		return err
	}

	fmt.Println("Token: ", token.Token)
	valid, err := client.ValidateToken(ctx, &pb.Token{
		Token: token.Token,
	})
	if err != nil {
		return err
	}
	if valid.Valide == true {
		fmt.Println("Token Validated")
	}

	return nil
}

func main() {
	service := micro.NewService(
		micro.Flags(
			&cli.StringFlag{
				Name:  "name",
				Usage: "Your Name",
			},
			&cli.StringFlag{
				Name:  "email",
				Usage: "Your Email",
			},
			&cli.StringFlag{
				Name:  "company",
				Usage: "Your Company",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "Password",
			},
		),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			name := c.String("name")
			email := c.String("email")
			company := c.String("company")
			password := c.String("password")

			log.Println("Test: ", name, email, company, password)

			ctx := context.Background()
			user := &pb.User{
				Name:     name,
				Email:    email,
				Company:  company,
				Password: password,
			}

			if err := createUser(ctx, service, user); err != nil {
				log.Println("error creating user: ", err.Error())
				return err
			}

			if err := authUser(ctx, service, &pb.User{
				Email:    email,
				Password: password,
			}); err != nil {
				log.Println("error authenticate user: ", err.Error())
				return err
			}

			return nil
		}),
	)
}
