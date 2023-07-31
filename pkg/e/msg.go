package e

var MsgMap = map[int]string{
	SUCCESS:      "操作成功",
	ERROR:        "操作失败",
	InvalidParam: "参数绑定失败",

	CreateTopicError: "创造话题失败",
	CreatePostError:  "创建回复失败",
	TopicNotExist:    "话题不存在",
	TopicIdIsNull:    "TopicId缺失",
}

func GetMsg(code int) string {
	msg, ok := MsgMap[code]
	if ok {
		return msg
	}
	return MsgMap[ERROR]
}
