package service

import (
	"AlittleRequire/pkg/e"
	"AlittleRequire/pkg/logger"
	"AlittleRequire/repository/dao"
	"AlittleRequire/repository/model"
	"AlittleRequire/types"
	"context"
	"errors"
	"fmt"
	"sync"
)

var TaskSrvIns *TaskSrv
var taskOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	taskOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (*TaskSrv) CreateTopic(ctx context.Context, req *types.CreateTopicReq) (interface{}, int, error) {
	code := e.SUCCESS
	taskDao := dao.NewTaskDao(ctx)
	topic := &model.Topic{
		Title:   req.Title,
		Content: req.Content,
	}
	err := taskDao.CreateTopic(topic)
	if err != nil {
		code = e.CreateTopicError
		return nil, code, err
	}
	return nil, code, nil
}

func (*TaskSrv) CreatePost(ctx context.Context, req *types.CreatePostReq) (interface{}, int, error) {
	code := e.SUCCESS
	taskDao := dao.NewTaskDao(ctx)
	if req.TopicId <= 0 {
		code = e.TopicIdIsNull
		return nil, code, errors.New("参数缺失")
	}
	_, err := taskDao.FindTopicByTopicId(req.TopicId)
	if err != nil {
		code = e.TopicNotExist
		return nil, code, err
	}
	post := &model.Post{
		Content: req.Content,
		TopicId: req.TopicId,
	}
	err = taskDao.CreatePost(post)
	if err != nil {
		code = e.CreatePostError
		return nil, code, err
	}
	return nil, code, nil
}

func (*TaskSrv) GetTopicInfo(ctx context.Context, req *types.GetTopicInfoReq) (interface{}, int, error) {
	code := e.SUCCESS
	taskDao := dao.NewTaskDao(ctx)
	if req.TopicId <= 0 {
		code = e.TopicIdIsNull
		return nil, code, fmt.Errorf("get topicInfo failed,err %s", e.GetMsg(code))
	}
	topic, err := taskDao.FindTopicByTopicId(req.TopicId)
	if err != nil {
		code = e.TopicNotExist
		return nil, code, err
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	posts, count, err := taskDao.GetPostList(req.TopicId, req.PageSize, req.PageNum)
	if err != nil {
		logger.LogrusObj.Error(err)
		code = e.ERROR
		return nil, code, err
	}
	postList := types.BuildPostList(posts, count)
	return &types.TopicInfoResp{
		Topic: topic,
		Posts: postList,
	}, code, nil
}
