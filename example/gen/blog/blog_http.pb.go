// Code generated by protoc-gen-go-fiber. DO NOT EDIT.
// versions:
// - protoc-gen-go-fiber v0.0.1
// - protoc             (unknown)
// source: blog/blog.proto

package v1

import (
	"context"
	"github.com/bufbuild/protovalidate-go"
	"github.com/gofiber/fiber/v2"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-fiber package it is being compiled against.

const (
	OperationBlogServiceCreateArticle = "/blog.app.v1.BlogService/CreateArticle"
	OperationBlogServiceGetArticles   = "/blog.app.v1.BlogService/GetArticles"
	OperationBlogServiceTest          = "/blog.app.v1.BlogService/Test"
)

type BlogServiceHTTPServer interface {
	CreateArticle(context.Context, *Article) (*Article, error)
	GetArticles(context.Context, *GetArticlesReq) (*GetArticlesResp, error)
	Test(context.Context, *TestRequest) (*TestResponse, error)
}

func RegisterBlogServiceHTTPServer(app *fiber.App, srv BlogServiceHTTPServer) {
	r := app.Group("/")
	r.Get("/v1/author/:author_id/articles", _BlogService_GetArticles0_HTTP_Handler(srv))
	r.Get("/v1/articles", _BlogService_GetArticles1_HTTP_Handler(srv))
	r.Post("/v1/author/:author_id/articles", _BlogService_CreateArticle0_HTTP_Handler(srv))
	r.Post("/v1/test/:id?", _BlogService_Test0_HTTP_Handler(srv))
}

func _BlogService_GetArticles0_HTTP_Handler(srv BlogServiceHTTPServer) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var in GetArticlesReq
		if err := c.QueryParser(&in); err != nil {
			return err
		}
		if err := c.ParamsParser(&in); err != nil {
			return err
		}
		v, err := protovalidate.New(
			protovalidate.WithFailFast(true),
		)
		if err != nil {
			return err
		}
		if err = v.Validate(&in); err != nil {
			return err
		}
		out, err := srv.GetArticles(c.Context(), &in)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "success",
			"data":    out,
		})
	}
}

func _BlogService_GetArticles1_HTTP_Handler(srv BlogServiceHTTPServer) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var in GetArticlesReq
		if err := c.QueryParser(&in); err != nil {
			return err
		}
		v, err := protovalidate.New(
			protovalidate.WithFailFast(true),
		)
		if err != nil {
			return err
		}
		if err = v.Validate(&in); err != nil {
			return err
		}
		out, err := srv.GetArticles(c.Context(), &in)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "success",
			"data":    out,
		})
	}
}

func _BlogService_CreateArticle0_HTTP_Handler(srv BlogServiceHTTPServer) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var in Article
		if err := c.BodyParser(&in); err != nil {
			return err
		}
		if err := c.QueryParser(&in); err != nil {
			return err
		}
		if err := c.ParamsParser(&in); err != nil {
			return err
		}
		v, err := protovalidate.New(
			protovalidate.WithFailFast(true),
		)
		if err != nil {
			return err
		}
		if err = v.Validate(&in); err != nil {
			return err
		}
		out, err := srv.CreateArticle(c.Context(), &in)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "success",
			"data":    out,
		})
	}
}

func _BlogService_Test0_HTTP_Handler(srv BlogServiceHTTPServer) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var in TestRequest
		if err := c.BodyParser(&in); err != nil {
			return err
		}
		if err := c.QueryParser(&in); err != nil {
			return err
		}
		if err := c.ParamsParser(&in); err != nil {
			return err
		}
		v, err := protovalidate.New(
			protovalidate.WithFailFast(true),
		)
		if err != nil {
			return err
		}
		if err = v.Validate(&in); err != nil {
			return err
		}
		out, err := srv.Test(c.Context(), &in)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "success",
			"data":    out,
		})
	}
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) CreateArticle(context.Context, *Article) (*Article, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "method CreateArticle not implemented")
}
func (UnimplementedBlogServiceServer) GetArticles(context.Context, *GetArticlesReq) (*GetArticlesResp, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "method GetArticles not implemented")
}
func (UnimplementedBlogServiceServer) Test(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, fiber.NewError(fiber.StatusNotImplemented, "method Test not implemented")
}
