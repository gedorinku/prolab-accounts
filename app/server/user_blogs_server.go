package server

import (
	"context"
	"database/sql"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "github.com/ProgrammingLab/prolab-accounts/api"
	"github.com/ProgrammingLab/prolab-accounts/app/di"
	"github.com/ProgrammingLab/prolab-accounts/app/interceptor"
	"github.com/ProgrammingLab/prolab-accounts/app/util"
	"github.com/ProgrammingLab/prolab-accounts/infra/record"
	"github.com/ProgrammingLab/prolab-accounts/model"
)

// UserBlogServiceServer is a composite interface of api_pb.UserBlogServiceServer and grapiserver.Server.
type UserBlogServiceServer interface {
	api_pb.UserBlogServiceServer
	grapiserver.Server
}

// NewUserBlogServiceServer creates a new UserBlogServiceServer instance.
func NewUserBlogServiceServer(store di.StoreComponent) UserBlogServiceServer {
	return &userBlogServiceServerImpl{
		StoreComponent: store,
	}
}

type userBlogServiceServerImpl struct {
	di.StoreComponent
}

var (
	// ErrFeedURLDetectAutomatically returns when feed url could not be found automatically
	ErrFeedURLDetectAutomatically = status.Error(codes.InvalidArgument, "feed url could not be found automatically")
	// ErrInvalidFeedURL returns when feed url is invalid
	ErrInvalidFeedURL = status.Error(codes.InvalidArgument, "feed url is invalid")
)

func (s *userBlogServiceServerImpl) CreateUserBlog(ctx context.Context, req *api_pb.CreateUserBlogRequest) (*api_pb.Blog, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	blog := req.GetBlog()
	feedURL, err := getFeedURL(ctx, s, req)
	if err != nil {
		return nil, err
	}

	b := &record.Blog{
		URL:     blog.GetUrl(),
		FeedURL: feedURL,
		UserID:  int64(userID),
	}

	bs := s.UserBlogStore(ctx)
	err = bs.CreateUserBlog(b)
	if err != nil {
		return nil, err
	}

	return blogToResponse(b), nil
}

func (s *userBlogServiceServerImpl) UpdateUserBlog(ctx context.Context, req *api_pb.UpdateUserBlogRequest) (*api_pb.Blog, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	blog := req.GetBlog()
	feedURL, err := getFeedURL(ctx, s, req)
	if err != nil {
		return nil, err
	}

	b := &record.Blog{
		ID:      int64(blog.GetBlogId()),
		URL:     blog.GetUrl(),
		FeedURL: feedURL,
		UserID:  int64(userID),
	}

	bs := s.UserBlogStore(ctx)

	if err := s.canWrite(ctx, userID, b.ID); err != nil {
		return nil, err
	}

	err = bs.UpdateUserBlog(b)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, util.ErrNotFound
		}
		return nil, err
	}

	return blogToResponse(b), nil
}

func (s *userBlogServiceServerImpl) DeleteUserBlog(ctx context.Context, req *api_pb.DeleteUserBlogRequest) (*empty.Empty, error) {
	userID, ok := interceptor.GetCurrentUserID(ctx)
	if !ok {
		return nil, util.ErrUnauthenticated
	}

	blogID := int64(req.GetBlogId())
	bs := s.UserBlogStore(ctx)

	if err := s.canWrite(ctx, userID, blogID); err != nil {
		return nil, err
	}

	err := bs.DeleteUserBlog(blogID)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (s *userBlogServiceServerImpl) canWrite(ctx context.Context, userID model.UserID, blogID int64) error {
	bs := s.UserBlogStore(ctx)
	b, err := bs.GetUserBlog(int64(blogID))
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return util.ErrNotFound
		}
		return err
	}
	if b.UserID != int64(userID) {
		return util.ErrNotFound
	}
	return nil
}

type blogRequest interface {
	GetBlog() *api_pb.Blog
	GetAutoDetectFeed() bool
}

func getFeedURL(ctx context.Context, s di.StoreComponent, req blogRequest) (string, error) {
	blog := req.GetBlog()
	if req.GetAutoDetectFeed() {
		fs := s.FeedStore(ctx)
		u, err := fs.GetFeedURL(blog.GetUrl())
		if err != nil {
			return "", ErrFeedURLDetectAutomatically
		}
		return u, nil
	}

	u := blog.GetFeedUrl()
	fs := s.FeedStore(ctx)
	err := fs.IsValidFeedURL(u)
	if err != nil {
		return "", ErrInvalidFeedURL
	}
	return u, nil
}

func blogToResponse(blog *record.Blog) *api_pb.Blog {
	if blog == nil {
		return nil
	}

	return &api_pb.Blog{
		BlogId:  uint32(blog.ID),
		Url:     blog.URL,
		FeedUrl: blog.FeedURL,
	}
}