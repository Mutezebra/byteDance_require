package types

import (
	"AlittleRequire/repository/model"
)

type CreateTopicReq struct {
	Content string `json:"content" form:"content"`
	Title   string `json:"title" form:"title"`
}

type CreatePostReq struct {
	Content string `json:"content" form:"content"`
	TopicId uint   `json:"topic_id" form:"topic_id"`
}

type GetTopicInfoReq struct {
	TopicId  uint `json:"topic_id" form:"topic_id"`
	PageSize int  `json:"page_size" form:"page_size"`
	PageNum  int  ` json:"page_num" form:"page_num"`
}

type PostListResp struct {
	Posts []*model.Post
	Total int64
}

type TopicInfoResp struct {
	Topic *model.Topic
	Posts *PostListResp
}

func BuildPostList(posts []model.Post, total int64) *PostListResp {
	var items []*model.Post
	for _, it := range posts {
		items = append(items, &model.Post{
			Content: it.Content,
		})
	}
	return &PostListResp{
		Posts: items,
		Total: total,
	}
}
