package service

import (
	"blog-kratos/models"
	"context"

	pb "blog-kratos/api/blog"
)

type ArticleService struct {
	pb.UnimplementedArticleServer
}

func NewArticleService() *ArticleService {
	return &ArticleService{}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	return &pb.CreateArticleReply{}, nil
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	return &pb.DeleteArticleReply{}, nil
}
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	return &pb.GetArticleReply{}, nil
}
func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	bz := make([]*models.BizArticle, 0)
	var total int64
	if err := models.DB.Model(new(models.BizArticle)).Count(&total).Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).Limit(int(req.GetPageSize())).
		Find(&bz).Error; err != nil {
		return nil, err
	}
	list := make([]*pb.ListRepoItem, 0, len(bz))
	for _, v := range bz {
		list = append(list, &pb.ListRepoItem{
			UserId:        v.UserId,
			SubjectId:     v.SubjectId,
			Title:         v.Title,
			Intro:         v.Intro,
			Cover:         v.Cover,
			Content:       v.Content,
			ContentMd:     v.ContentMd,
			Code:          v.Code,
			KeyWords:      v.KeyWords,
			Sort:          v.Sort,
			IsElite:       v.IsElite,
			Hits:          v.Hits,
			ArticleStatus: v.ArticleStatus,
		})
	}
	return &pb.ListArticleReply{
		List:  list,
		Total: total,
	}, nil
}
