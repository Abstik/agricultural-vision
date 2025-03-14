package redis

// redis key 注意使用命名空间
const (
	Prefix           = "agricultural_vision" // 项目key前缀
	KeyPostTimeZSet  = "post:time"           // zset;成员名:帖子(postID),分数:发帖时间
	KeyPostScoreZSet = "post:score"          // zset;成员名:帖子(postID),分数:投票分数

	KeyPostVotedZSetPF = "post:voted:" // zset;记录某一个帖子投票的用户及投票类型，+postId构成完整键名（指定哪个帖子）
	//成员名:用户id(userID),分数:投票类型(1,-1)

	KeyCommunitySetPF = "community:" //set;保存每个分区下所有帖子的id，+communityid构成完整键名

	KeyPostCommentNumZSet     = "post:comment_num"    // zset;成员名:帖子id,分数:评论数
	KeyPostCommentTimeZSetPF  = "post:comment_time:"  // zset;+帖子(postID)构成键名，成员名:评论id，分数:评论时间
	KeyPostCommentScoreZSetPF = "post:comment_score:" // zset;+帖子(postID)构成键名，成员名:评论id，分数:投票分数
)

func getRedisKey(key string) string {
	return Prefix + key
}
