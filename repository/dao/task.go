package dao

import (
	"AlittleRequire/pkg/logger"
	"AlittleRequire/repository/model"
	"context"
	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) FindTopicByTopicId(tid uint) (*model.Topic, error) {
	var topic *model.Topic
	err := dao.Model(&model.Topic{}).Where("id=?", tid).First(&topic).Error
	if err != nil {
		logger.LogrusObj.Error(err)
		return nil, err
	}
	return topic, err
}

func (dao *TaskDao) CreateTopic(topic *model.Topic) error {
	err := dao.Model(&model.Topic{}).Create(topic).Error
	if err != nil {
		logger.LogrusObj.Error(err)
		return err
	}
	return nil
}

func (dao *TaskDao) CreatePost(post *model.Post) error {
	err := dao.Model(&model.Post{}).Create(post).Error
	if err != nil {
		logger.LogrusObj.Error(err)
		return err
	}
	return nil
}

func (dao *TaskDao) GetPostList(tid uint, pageSize, pageNum int) ([]model.Post, int64, error) {
	var posts []model.Post
	var count int64
	err := dao.Model(&model.Post{}).Where("topic_id=?", tid).Count(&count).
		Limit(pageSize).Offset((pageNum - 1) * pageSize).
		Find(&posts).Error
	if err != nil {
		logger.LogrusObj.Error(err)
		return nil, 0, err
	}
	return posts, count, nil
}
