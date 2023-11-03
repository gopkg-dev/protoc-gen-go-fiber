package main

import (
	"context"
	"encoding/json"
	v1 "example/gen/blog"

	"github.com/bufbuild/protovalidate-go"

	"example/proto/errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type service struct {
	v1.UnimplementedBlogServiceServer
}

func (s service) Test(ctx context.Context, request *v1.TestRequest) (*v1.TestResponse, error) {

	marshal, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return &v1.TestResponse{Message: string(marshal)}, nil
}

func (s service) CreateArticle(ctx context.Context, article *v1.Article) (*v1.Article, error) {
	if article.AuthorId < 1 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "author id must > 0")
	}
	return article, nil
}

func (s service) GetArticles(ctx context.Context, req *v1.GetArticlesReq) (*v1.GetArticlesResp, error) {
	if req.AuthorId < 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "author id must >= 0")
	}
	return &v1.GetArticlesResp{
		Total: 1,
		Articles: []*v1.Article{
			{
				Title:    "test article: " + req.Title,
				Content:  "test",
				AuthorId: req.AuthorId,
			},
		},
	}, nil
}

func main() {
	app := fiber.New(fiber.Config{
		Prefork:      false,
		ErrorHandler: defaultErrorHandler,
	})

	app.Use(logger.New(), recover.New())

	v1.RegisterBlogServiceHTTPServer(app, &service{})

	app.Use(defaultNotFoundHandler)

	for _, route := range app.GetRoutes() {
		fmt.Println(route.Name, route.Method, route.Path)
	}

	app.Listen(":3000")
}

func defaultNotFoundHandler(c *fiber.Ctx) error {
	return errors.NotFound("NOT_FOUND", "Cannot %s %s", c.Method(), c.Path())
}

func defaultErrorHandler(c *fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.Status(fiberErr.Code).JSON(errors.Error{
			Code:    fiberErr.Code,
			Reason:  errors.UnknownReason,
			Message: fiberErr.Message,
		})
	}
	var validationErr *protovalidate.ValidationError
	if ok := errors.As(err, &validationErr); ok {
		metadata := make(map[string]interface{}, len(validationErr.Violations))
		for _, v := range validationErr.Violations {
			metadata[v.FieldPath] = v.Message
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error{
			Code:     fiber.StatusBadRequest,
			Reason:   "VALIDATION_FAILED",
			Message:  fmt.Sprintf("%s: %s [%s]", validationErr.Violations[0].FieldPath, validationErr.Violations[0].Message, validationErr.Violations[0].ConstraintId),
			Metadata: metadata,
		})
	}
	e := errors.FromError(err)
	return c.Status(e.Code).JSON(e)
}
