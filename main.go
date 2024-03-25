package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	GET = iota
	SET = iota
)

type FSMLog struct {
	Optype int
	Key    string
	Value  string
}

type NonVolatileState struct {
	CurrentTerm int
	VotedFor    string
	Log         []Log
}

type VolatileState struct {
	CommitIndex int
	LastApplied int
}

type VolatileStateLeader struct {
	NextIndex  []int
	MatchIndex []int
}

type RaftGrpcServer struct{}

func (s *RaftGrpcServer) AppendEntries(ctx context.Context, req *AppendEntriesRequest) (*AppendEntriesResponse, error) {
	res := AppendEntriesResponse{Term: 1, Success: true}
	return &res, nil
}

func (s *RaftGrpcServer) mustEmbedUnimplementedRaftServiceServer() {}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	RegisterRaftServiceServer(s, &RaftGrpcServer{})
	log.Printf("Server listening on port %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
