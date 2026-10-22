package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	pb "example.com/grpc-review/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func usage() error {
	return errors.New("usage: client create <title> | get <id> | list | complete <id>")
}

func parseID(value string) (int64, error) {
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid todo id %q", value)
	}
	return id, nil
}

func printTodo(todo *pb.Todo) {
	fmt.Printf("id=%d title=%q completed=%t\n", todo.GetId(), todo.GetTitle(), todo.GetCompleted())
}

func run() error {
	if len(os.Args) < 2 {
		return usage()
	}

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("create gRPC client: %w", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch os.Args[1] {
	case "create":
		if len(os.Args) != 3 {
			return usage()
		}
		todo, err := client.CreateTodo(ctx, &pb.CreateTodoRequest{Title: os.Args[2]})
		if err != nil {
			return err
		}
		printTodo(todo)

	case "get":
		if len(os.Args) != 3 {
			return usage()
		}
		id, err := parseID(os.Args[2])
		if err != nil {
			return err
		}
		todo, err := client.GetTodo(ctx, &pb.GetTodoRequest{Id: id})
		if err != nil {
			return err
		}
		printTodo(todo)

	case "list":
		if len(os.Args) != 2 {
			return usage()
		}
		response, err := client.ListTodos(ctx, &pb.ListTodosRequest{})
		if err != nil {
			return err
		}
		for _, todo := range response.GetTodos() {
			printTodo(todo)
		}

	case "complete":
		if len(os.Args) != 3 {
			return usage()
		}
		id, err := parseID(os.Args[2])
		if err != nil {
			return err
		}
		todo, err := client.CompleteTodo(ctx, &pb.CompleteTodoRequest{Id: id})
		if err != nil {
			return err
		}
		printTodo(todo)

	default:
		return usage()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		if grpcStatus, ok := status.FromError(err); ok {
			fmt.Fprintf(os.Stderr, "gRPC error: code=%s message=%s\n", grpcStatus.Code(), grpcStatus.Message())
		} else {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		os.Exit(1)
	}
}
