package service

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/TremblingV5/DouTok/applications/relation/pack"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{ctx: ctx}
}

// 关注 || 取关动作
func (s *RelationActionService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	// 在 SafeMap 中更新局部关注数和粉丝数
	followKey := constants.FollowCount + string(req.UserId)
	followerKey := constants.FollowerCount + string((req.ToUserId))
	follow, ok := ConcurrentMap.Get(followKey)
	if !ok {
		klog.Errorf("get follow count from concurrentMap error")
	}
	follower, ok := ConcurrentMap.Get(followerKey)
	if !ok {
		klog.Errorf("get follow count from concurrentMap error")
	}
	mu.Lock()
	ConcurrentMap.Set(followKey, follow.(int64)+1)
	ConcurrentMap.Set(followerKey, follower.(int64)+1)
	mu.Unlock()
	// 使用同步producer，将消息存入 kafka
	defer func() {
		_ = SyncProducer.Close()
	}()
	// 构建消息
	val, err := json.Marshal(pack.NewRelation(req))
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: ViperConfig.Viper.GetStringSlice("Kafka.Topics")[0],
		Value: sarama.StringEncoder(val),
	}
	partition, offset, err := SyncProducer.SendMessage(msg)

	if err == nil {
		klog.Infof("produce success, partition:", partition, ",offset:", offset)
	} else {
		return err
	}

	return nil
}