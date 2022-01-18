package group

import (
	"open_im_sdk/pkg/common"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/sdk_params_callback"
	"open_im_sdk/pkg/utils"
)

func (g *Group) SetGroupListener(callback OnGroupListener) {
	if callback == nil {
		return
	}
	g.listener = callback
}

func (g *Group) CreateGroup(callback common.Base, groupBaseInfo string, memberList string, operationID string) {
	if callback == nil {
		return
	}
	go func() {
		log.NewInfo(operationID, utils.RunFuncName(), groupBaseInfo, memberList)
		var unmarshalCreateGroupBaseInfoParam sdk_params_callback.CreateGroupBaseInfoParam
		common.JsonUnmarshalAndArgsValidate(groupBaseInfo, &unmarshalCreateGroupBaseInfoParam, callback, operationID)
		var unmarshalCreateGroupMemberRoleParam sdk_params_callback.CreateGroupMemberRoleParam
		common.JsonUnmarshalAndArgsValidate(memberList, &unmarshalCreateGroupMemberRoleParam, callback, operationID)
		result := g.createGroup(callback, unmarshalCreateGroupBaseInfoParam, unmarshalCreateGroupMemberRoleParam, operationID)
		callback.OnSuccess(utils.StructToJsonString(result))
		log.NewInfo(operationID, "CreateGroup callback: ", utils.StructToJsonString(result))
	}()
}

func (g *Group) JoinGroup(callback common.Base, groupID, reqMsg string, operationID string) {
	if callback == nil {
		return
	}
	//go func() {
	//	log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, reqMsg)
	//	g.joinGroup(groupID, reqMsg, callback, operationID)
	//	callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.JoinGroupCallback))
	//	log.NewInfo(operationID, "JoinGroup callback: ", utils.StructToJsonString(sdk_params_callback.JoinGroupCallback))
	//}()
}

func (g *Group) QuitGroup(callback common.Base, groupID string, operationID string) {
	if callback == nil {
		return
	}
	//go func() {
	//	log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID)
	//	g.quitGroup(groupID, callback, operationID)
	//	callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.QuitGroupCallback))
	//	log.NewInfo(operationID, "QuitGroup callback: ", utils.StructToJsonString(sdk_params_callback.QuitGroupCallback))
	//}()
}

func (g *Group) GetJoinedGroupList(callback common.Base, operationID string) {
	if callback == nil {
		return
	}
	//go func() {
	//	log.NewInfo(operationID, utils.RunFuncName(), "args: ")
	//	groupList := g.getJoinedGroupList(callback, operationID)
	//	callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(groupList)))
	//	log.NewInfo(operationID, "QuitGroup callback: ", utils.StructToJsonString(utils.StructToJsonString(groupList)))
	//}()
}

func (g *Group) GetGroupsInfo(callback common.Base, groupIDList string, operationID string) {
	if callback == nil {
		return
	}
	//go func() {
	//	log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupIDList)
	//	var unmarshalGetGroupsInfoParam sdk_params_callback.GetGroupsInfoParam
	//	common.JsonUnmarshalAndArgsValidate(groupIDList, &unmarshalGetGroupsInfoParam, callback, operationID)
	//	groupsInfoList := g.getGroupsInfo(unmarshalGetGroupsInfoParam, callback, operationID)
	//	callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(groupsInfoList)))
	//	log.NewInfo(operationID, "GetGroupsInfo callback: ", utils.StructToJsonString(utils.StructToJsonString(groupsInfoList)))
	//
	//}()
}

//
//func (g *Group) SetGroupInfo(callback common.Base, groupInfo string, groupID string, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupInfo, groupID)
//		var unmarshalSetGroupInfoParam sdk_params_callback.SetGroupInfoParam
//		common.JsonUnmarshalAndArgsValidate(groupInfo, &unmarshalSetGroupInfoParam, callback, operationID)
//		g.setGroupInfo(callback, unmarshalSetGroupInfoParam, groupID, operationID)
//		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(sdk_params_callback.SetGroupInfoCallback)))
//		log.NewInfo(operationID, "SetGroupInfo callback: ", utils.StructToJsonString(sdk_params_callback.SetGroupInfoCallback))
//	}()
//}
//
//func (g *Group) GetGroupMemberList(callback common.Base, groupID string, filter int32, next int32, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, filter, next)
//		groupMemberList := g.getGroupMemberList(callback, groupID, filter, next, operationID)
//		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(groupMemberList)))
//		log.NewInfo(operationID, "GetGroupMemberList callback: ", utils.StructToJsonString(groupMemberList))
//	}()
//}
//
//func (g *Group) GetGroupMembersInfo(callback common.Base, groupID string, userIDList string, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, userIDList)
//		var unmarshalParam sdk_params_callback.GetGroupMembersInfoParam
//		common.JsonUnmarshal(userIDList, &unmarshalParam, callback, operationID)
//		groupMemberList := g.getGroupMembersInfo(callback, groupID, unmarshalParam, operationID)
//		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(groupMemberList)))
//		log.NewInfo(operationID, "GetGroupMembersInfo callback: ", utils.StructToJsonString(groupMemberList))
//	}()
//}
//
//func (g *Group) KickGroupMember(callback common.Base, groupID string, reason string, userIDList string, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, reason, userIDList)
//		var unmarshalParam sdk_params_callback.KickGroupMemberParam
//		common.JsonUnmarshal(userIDList, &unmarshalParam, callback, operationID)
//		result := g.kickGroupMember(callback, groupID, unmarshalParam, reason, operationID)
//		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(result)))
//		log.NewInfo(operationID, "GetGroupMembersInfo callback: ", utils.StructToJsonString(result))
//	}()
//}
//
//func (g *Group) TransferGroupOwner(callback common.Base, groupID, newOwnerUserID string, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		g.transferGroupOwner(callback, groupID, newOwnerUserID, operationID)
//		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(sdk_params_callback.TransferGroupOwnerCallback)))
//	}()
//}
//
//func (g *Group) InviteUserToGroup(callback common.Base, groupID, reason string, userIDList string, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, reason, userIDList)
//		var unmarshalParam sdk_params_callback.InviteUserToGroupParam
//		common.JsonUnmarshalAndArgsValidate(userIDList, &unmarshalParam, callback, operationID)
//		result := g.inviteUserToGroup(callback, groupID, reason, unmarshalParam, operationID)
//		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(result)))
//		log.NewInfo(operationID, utils.RunFuncName(), "callback: ", utils.StructToJsonString(result))
//	}()
//}
//
//func (g *Group) GetGroupApplicationList(callback common.Base, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		log.NewInfo(operationID, utils.RunFuncName(), "args: ")
//		result := g.getGroupApplicationList(callback, operationID)
//		callback.OnSuccess(utils.StructToJsonString(utils.StructToJsonString(result)))
//		log.NewInfo(operationID, utils.RunFuncName(), "callback: ", utils.StructToJsonString(result))
//	}()
//}
//
//func (g *Group) AcceptGroupApplication(callback common.Base, groupID, fromUserID, handleMsg string, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, fromUserID, handleMsg)
//		g.processGroupApplication(callback, groupID, fromUserID, handleMsg, 1, operationID)
//		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.AcceptGroupApplicationCallback))
//		log.NewInfo(operationID, utils.RunFuncName(), "callback: ", utils.StructToJsonString(sdk_params_callback.AcceptGroupApplicationCallback))
//	}()
//}
//
//func (g *Group) RefuseGroupApplication(callback common.Base, groupID, fromUserID, handleMsg string, operationID string) {
//	if callback == nil {
//		return
//	}
//	go func() {
//		log.NewInfo(operationID, utils.RunFuncName(), "args: ", groupID, fromUserID, handleMsg)
//		g.processGroupApplication(callback, groupID, fromUserID, handleMsg, -1, operationID)
//		callback.OnSuccess(utils.StructToJsonString(sdk_params_callback.RefuseGroupApplicationCallback))
//		log.NewInfo(operationID, utils.RunFuncName(), "callback: ", utils.StructToJsonString(sdk_params_callback.RefuseGroupApplicationCallback))
//	}()
//}

/*
func (u *UserRelated) TsetGetGroupApplicationList(callback Base) string {
	if callback == nil {
		return ""
	}

	r, err := u.getGroupApplicationList()
	if err != nil {
		sdkLog("getGroupApplicationList faild, ", err.Error())
		callback.OnError(ErrCodeGroup, err.Error())
		return ""
	}
	jsonResult, err := json.Marshal(r)
	if err != nil {
		sdkLog("getGroupApplicationList faild, ", err.Error())
		callback.OnError(ErrCodeGroup, err.Error())
		return ""
	}
	callback.OnSuccess(string(jsonResult))
	return string(jsonResult)

}*/
