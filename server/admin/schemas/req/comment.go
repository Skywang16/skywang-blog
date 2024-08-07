package req

import "time"

// commentsReq 登录日志列表参数
type commentsReq struct {
	Username  string    `form:"username"`                           // 登录账号
	Content   string    `form:"content"`                            // 评论内容
	StartTime time.Time `form:"startTime" time_format:"2006-01-02"` // 评论时间
}
