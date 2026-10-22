package main

import (
	"context"
	"log"
	"net"
	"sort"
	"strings"
	"sync"

	pb "example.com/grpc-review/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type todoServer struct {
	pb.UnimplementedTodoServiceServer

	mu     sync.RWMutex
	nextID int64
	todos  map[int64]*pb.Todo
}

func newTodoServer() *todoServer {
	return &todoServer{
		todos: make(map[int64]*pb.Todo),
	}
}

func cloneTodo(todo *pb.Todo) *pb.Todo {
	return &pb.Todo{
		Id:        todo.GetId(),
		Title:     todo.GetTitle(),
		Completed: todo.GetCompleted(),
	}
}

func (s *todoServer) CreateTodo(_ context.Context, req *pb.CreateTodoRequest) (*pb.Todo, error) {
	title := strings.TrimSpace(req.GetTitle())
	if title == "" {
		return nil, status.Error(codes.InvalidArgument, "title must not be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	todo := &pb.Todo{
		Id:        s.nextID,
		Title:     title,
		Completed: false,
	}
	s.todos[todo.Id] = todo

	return cloneTodo(todo), nil
}

func (s *todoServer) GetTodo(_ context.Context, req *pb.GetTodoRequest) (*pb.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, ok := s.todos[req.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "todo %d not found", req.GetId())
	}

	return cloneTodo(todo), nil
}

func (s *todoServer) ListTodos(_ context.Context, _ *pb.ListTodosRequest) (*pb.ListTodosResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*pb.Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, cloneTodo(todo))
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].GetId() < todos[j].GetId()
	})

	return &pb.ListTodosResponse{Todos: todos}, nil
}

func (s *todoServer) CompleteTodo(_ context.Context, req *pb.CompleteTodoRequest) (*pb.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.todos[req.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "todo %d not found", req.GetId())
	}

	todo.Completed = true
	return cloneTodo(todo), nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, newTodoServer())

	log.Println("gRPC server listening on localhost:50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
