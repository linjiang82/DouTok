// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/TremblingV5/DouTok/applications/api/biz/handler"
	"github.com/TremblingV5/DouTok/applications/api/initialize"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/relation"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/hertz-contrib/jwt"

	api "github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
)

// RelationAction .
//
//	@Tags		Relation关注
//
//	@Summary	关注或取消关注
//	@Description
//	@Param		req		body		api.DouyinRelationActionRequest	true	"关注操作信息"
//	@Success	200		{object}	relation.DouyinRelationActionResponse
//	@Failure	default	{object}	api.DouyinRelationActionResponse
//	@router		/douyin/relation/action [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildRelationActionResp(errno.ErrBind))
		return
	}

	userId := int64(jwt.ExtractClaims(ctx, c)[initialize.AuthMiddleware.IdentityKey].(float64))

	rpcReq := relation.DouyinRelationActionRequest{
		UserId:     userId,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	}
	resp, err := rpc.RelationAction(ctx, rpc.RelationClient, &rpcReq)
	if err != nil {
		handler.SendResponse(c, handler.BuildRelationActionResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}

// RelationFollowList .
//
//	@Tags		Relation关注
//
//	@Summary	获取已关注用户的列表
//	@Description
//	@Param		req		query		api.DouyinRelationFollowListRequest	true	"获取操作信息"
//	@Success	200		{object}	relation.DouyinRelationFollowListResponse
//	@Failure	default	{object}	api.DouyinRelationFollowListResponse
//
//	@router		/douyin/relation/follow/list [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildRelationFollowListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.RelationFollowList(ctx, rpc.RelationClient, &relation.DouyinRelationFollowListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildRelationFollowListResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}

// RelationFollowerList .
//
//	@Tags		Relation关注
//
//	@Summary	获取粉丝用户列表
//	@Description
//	@Param		req		query		api.DouyinRelationFollowerListRequest	true	"获取操作信息"
//	@Success	200		{object}	relation.DouyinRelationFollowerListResponse
//	@Failure	default	{object}	api.DouyinRelationFollowerListResponse
//	@router		/douyin/relation/follower/list [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowerListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildRelationFollowerListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.RelationFollowerList(ctx, rpc.RelationClient, &relation.DouyinRelationFollowerListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildRelationFollowerListResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}

// RelationFriendList .
//
//	@Tags		Relation关注
//
//	@Summary	获取好友列表
//	@Description
//	@Param		req		query		api.DouyinRelationFollowerListRequest	true	"获取操作信息"
//	@Success	200		{object}	relation.DouyinRelationFollowerListResponse
//	@Failure	default	{object}	api.DouyinRelationFollowerListResponse
//	@router		/douyin/relation/friend/list [GET]
//
// 内部为调用获取粉丝列表
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFriendListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		handler.SendResponse(c, handler.BuildRelationFriendListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.RelationFriendList(ctx, rpc.RelationClient, &relation.DouyinRelationFriendListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		handler.SendResponse(c, handler.BuildRelationFriendListResp(errno.ConvertErr(err)))
		return
	}
	// TODO 此处直接返回了 rpc 的 resp
	handler.SendResponse(c, resp)
}
