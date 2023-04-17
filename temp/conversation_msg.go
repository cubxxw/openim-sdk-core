package open_im_sdk

import "open_im_sdk/open_im_sdk_callback"

func GetAllConversationList(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Conversation().GetAllConversationList)
}

func GetConversationListSplit(callback open_im_sdk_callback.Base, operationID string, offset int, count int) {
	call(callback, operationID, UserForSDK.Conversation().GetConversationListSplit, offset, count)
}

func GetOneConversation(callback open_im_sdk_callback.Base, operationID string, sessionType int, sourceID string) {
	call(callback, operationID, UserForSDK.Conversation().GetOneConversation, sessionType, sourceID)
}

func GetMultipleConversation(callback open_im_sdk_callback.Base, operationID string, conversationIDList string) {
	call(callback, operationID, UserForSDK.Conversation().GetMultipleConversation, conversationIDList)
}

func SetConversationRecvMessageOpt(callback open_im_sdk_callback.Base, operationID string, conversationIDList string, opt int) {
	call(callback, operationID, UserForSDK.Conversation().SetConversationRecvMessageOpt, conversationIDList, opt)
}

func SetGlobalRecvMessageOpt(callback open_im_sdk_callback.Base, operationID string, opt int) {
	call(callback, operationID, UserForSDK.Conversation().SetGlobalRecvMessageOpt, opt)
}

func GetConversationRecvMessageOpt(callback open_im_sdk_callback.Base, operationID string, conversationIDList string) {
	call(callback, operationID, UserForSDK.Conversation().GetConversationRecvMessageOpt, conversationIDList)
}

func DeleteConversation(callback open_im_sdk_callback.Base, operationID string, conversationID string) {
	call(callback, operationID, UserForSDK.Conversation().DeleteConversation, conversationID)
}

func DeleteAllConversationFromLocal(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Conversation().DeleteAllConversationFromLocal)
}

func SetConversationDraft(callback open_im_sdk_callback.Base, operationID string, conversationID string, draftText string) {
	call(callback, operationID, UserForSDK.Conversation().SetConversationDraft, conversationID, draftText)
}

func ResetConversationGroupAtType(callback open_im_sdk_callback.Base, operationID string, conversationID string) {
	call(callback, operationID, UserForSDK.Conversation().ResetConversationGroupAtType, conversationID)
}

func PinConversation(callback open_im_sdk_callback.Base, operationID string, conversationID string, isPinned bool) {
	call(callback, operationID, UserForSDK.Conversation().PinConversation, conversationID, isPinned)
}

func GetTotalUnreadMsgCount(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Conversation().GetTotalUnreadMsgCount)
}

func FindMessageList(callback open_im_sdk_callback.Base, operationID string, findMessageOptions string) {
	call(callback, operationID, UserForSDK.Conversation().FindMessageList, findMessageOptions)
}

func GetHistoryMessageList(callback open_im_sdk_callback.Base, operationID string, getMessageOptions string) {
	call(callback, operationID, UserForSDK.Conversation().GetHistoryMessageList, getMessageOptions)
}

func GetAdvancedHistoryMessageList(callback open_im_sdk_callback.Base, operationID string, getMessageOptions string) {
	call(callback, operationID, UserForSDK.Conversation().GetAdvancedHistoryMessageList, getMessageOptions)
}

func GetAdvancedHistoryMessageListReverse(callback open_im_sdk_callback.Base, operationID string, getMessageOptions string) {
	call(callback, operationID, UserForSDK.Conversation().GetAdvancedHistoryMessageListReverse, getMessageOptions)
}

func GetHistoryMessageListReverse(callback open_im_sdk_callback.Base, operationID string, getMessageOptions string) {
	call(callback, operationID, UserForSDK.Conversation().GetHistoryMessageListReverse, getMessageOptions)
}

func RevokeMessage(callback open_im_sdk_callback.Base, operationID string, message string) {
	call(callback, operationID, UserForSDK.Conversation().RevokeMessage, message)
}

func NewRevokeMessage(callback open_im_sdk_callback.Base, operationID string, message string) {
	call(callback, operationID, UserForSDK.Conversation().NewRevokeMessage, message)
}

func TypingStatusUpdate(callback open_im_sdk_callback.Base, operationID string, recvID string, msgTip string) {
	call(callback, operationID, UserForSDK.Conversation().TypingStatusUpdate, recvID, msgTip)
}

func MarkC2CMessageAsRead(callback open_im_sdk_callback.Base, operationID string, userID string, msgIDList string) {
	call(callback, operationID, UserForSDK.Conversation().MarkC2CMessageAsRead, userID, msgIDList)
}

func MarkMessageAsReadByConID(callback open_im_sdk_callback.Base, operationID string, conversationID string, msgIDList string) {
	call(callback, operationID, UserForSDK.Conversation().MarkMessageAsReadByConID, conversationID, msgIDList)
}

func MarkGroupMessageHasRead(callback open_im_sdk_callback.Base, operationID string, groupID string) {
	call(callback, operationID, UserForSDK.Conversation().MarkGroupMessageHasRead, groupID)
}

func MarkGroupMessageAsRead(callback open_im_sdk_callback.Base, operationID string, groupID string, msgIDList string) {
	call(callback, operationID, UserForSDK.Conversation().MarkGroupMessageAsRead, groupID, msgIDList)
}

func DeleteMessageFromLocalStorage(callback open_im_sdk_callback.Base, operationID string, message string) {
	call(callback, operationID, UserForSDK.Conversation().DeleteMessageFromLocalStorage, message)
}

func DeleteMessageFromLocalAndSvr(callback open_im_sdk_callback.Base, operationID string, message string) {
	call(callback, operationID, UserForSDK.Conversation().DeleteMessageFromLocalAndSvr, message)
}

func DeleteConversationFromLocalAndSvr(callback open_im_sdk_callback.Base, operationID string, conversationID string) {
	call(callback, operationID, UserForSDK.Conversation().DeleteConversationFromLocalAndSvr, conversationID)
}

func DeleteAllMsgFromLocalAndSvr(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Conversation().DeleteAllMsgFromLocalAndSvr)
}

func DeleteAllMsgFromLocal(callback open_im_sdk_callback.Base, operationID string) {
	call(callback, operationID, UserForSDK.Conversation().DeleteAllMsgFromLocal)
}

func ClearC2CHistoryMessage(callback open_im_sdk_callback.Base, operationID string, userID string) {
	call(callback, operationID, UserForSDK.Conversation().ClearC2CHistoryMessage, userID)
}

func ClearC2CHistoryMessageFromLocalAndSvr(callback open_im_sdk_callback.Base, operationID string, userID string) {
	call(callback, operationID, UserForSDK.Conversation().ClearC2CHistoryMessageFromLocalAndSvr, userID)
}

func ClearGroupHistoryMessage(callback open_im_sdk_callback.Base, operationID string, groupID string) {
	call(callback, operationID, UserForSDK.Conversation().ClearGroupHistoryMessage, groupID)
}

func ClearGroupHistoryMessageFromLocalAndSvr(callback open_im_sdk_callback.Base, operationID string, groupID string) {
	call(callback, operationID, UserForSDK.Conversation().ClearGroupHistoryMessageFromLocalAndSvr, groupID)
}

func InsertSingleMessageToLocalStorage(callback open_im_sdk_callback.Base, operationID string, message string, recvID string, sendID string) {
	call(callback, operationID, UserForSDK.Conversation().InsertSingleMessageToLocalStorage, message, recvID, sendID)
}

func InsertGroupMessageToLocalStorage(callback open_im_sdk_callback.Base, operationID string, message string, groupID string, sendID string) {
	call(callback, operationID, UserForSDK.Conversation().InsertGroupMessageToLocalStorage, message, groupID, sendID)
}

func SearchLocalMessages(callback open_im_sdk_callback.Base, operationID string, searchParam string) {
	call(callback, operationID, UserForSDK.Conversation().SearchLocalMessages, searchParam)
}
