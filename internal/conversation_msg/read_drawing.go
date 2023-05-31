package conversation_msg

import (
	"context"
	"open_im_sdk/pkg/common"
	"open_im_sdk/pkg/constant"
	"open_im_sdk/pkg/db/model_struct"

	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/log"
	pbConversation "github.com/OpenIMSDK/Open-IM-Server/pkg/proto/conversation"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/proto/wrapperspb"
)

func (c *Conversation) setHasReadSeq(ctx context.Context, conversationID string, conversationType int32, maxSeq int64, m map[string][]string) error {
	return c.newSetConversation(ctx, &pbConversation.SetConversationsReq{Conversation: &pbConversation.ConversationReq{
		ConversationID:   conversationID,
		ConversationType: conversationType,
		HasReadSeq:       wrapperspb.Int64(maxSeq),
		// AsReadMsgMap:     m,
	}})
}

// mark a conversation's all message as read
func (c *Conversation) markConversationMessageAsRead(ctx context.Context, conversationID string, conversationType int32) error {
	c.markAsReadLock.Lock()
	defer c.markAsReadLock.Unlock()
	peerUserMaxSeq, err := c.db.GetConversationPeerNormalMsgSeq(ctx, conversationID)
	if err != nil {
		return err
	}
	maxSeq, err := c.db.GetConversationNormalMsgSeq(ctx, conversationID)
	if err != nil {
		return err
	}
	msgs, err := c.db.GetUnreadMessage(ctx, conversationID)
	if err != nil {
		return err
	}
	msgIDs, asReadMsgMap := c.getAsReadMsgMapAndList(ctx, msgs)
	if err := c.setHasReadSeq(ctx, conversationID, conversationType, peerUserMaxSeq, asReadMsgMap); err != nil {
		return err
	}
	_, err = c.db.MarkConversationMessageAsRead(ctx, conversationID, msgIDs)
	if err != nil {
		return err
	}
	if err := c.db.UpdateColumnsConversation(ctx, conversationID, map[string]interface{}{"unread_count": 0}); err != nil {
		log.ZError(ctx, "UpdateColumnsConversation err", err, "conversationID", conversationID)
	}
	c.unreadChangeTrigger(ctx, conversationID, peerUserMaxSeq == maxSeq)
	return nil
}

// mark a conversation's message as read by seqs
func (c *Conversation) markConversationMessageAsReadBySeqs(ctx context.Context, conversationID string, conversationType int32, msgIDs []string) error {
	c.markAsReadLock.Lock()
	defer c.markAsReadLock.Unlock()
	msgs, err := c.db.GetMessagesByClientMsgIDs(ctx, conversationID, msgIDs)
	if err != nil {
		return err
	}
	if len(msgs) == 0 {
		return nil
	}
	asReadMsgIDs, asReadMsgMap := c.getAsReadMsgMapAndList(ctx, msgs)
	var hasReadSeq = msgs[0].Seq
	maxSeq, err := c.db.GetConversationNormalMsgSeq(ctx, conversationID)
	if err != nil {
		return err
	}
	if err := c.setHasReadSeq(ctx, conversationID, conversationType, hasReadSeq, asReadMsgMap); err != nil {
		return err
	}

	decrCount, err := c.db.MarkConversationMessageAsRead(ctx, conversationID, asReadMsgIDs)
	if err != nil {
		return err
	}
	if err := c.db.DecrConversationUnreadCount(ctx, conversationID, decrCount); err != nil {
		log.ZError(ctx, "decrConversationUnreadCount err", err, "conversationID", conversationID, "decrCount", decrCount)
	}
	c.unreadChangeTrigger(ctx, conversationID, hasReadSeq == maxSeq && msgs[0].SendID != c.loginUserID)
	return nil
}

func (c *Conversation) getAsReadMsgMapAndList(ctx context.Context, msgs []*model_struct.LocalChatLog) ([]string, map[string][]string) {
	var asReadMsgIDs []string
	var asReadMsgMap = make(map[string][]string)
	for _, msg := range msgs {
		if !msg.IsRead && msg.ContentType < constant.NotificationBegin && msg.SendID != c.loginUserID {
			asReadMsgIDs = append(asReadMsgIDs, msg.ClientMsgID)
			if v, ok := asReadMsgMap[msg.SendID]; ok {
				v = append(v, msg.ClientMsgID)
				asReadMsgMap[msg.SendID] = v
			} else {
				asReadMsgMap[msg.SendID] = []string{msg.ClientMsgID}
			}
		} else {
			log.ZWarn(ctx, "msg can't marked as read", nil, "msg", msg)
		}
	}
	return asReadMsgIDs, asReadMsgMap
}

func (c *Conversation) unreadChangeTrigger(ctx context.Context, conversationID string, latestMsgIsRead bool) {
	if latestMsgIsRead {
		_ = common.TriggerCmdUpdateConversation(ctx, common.UpdateConNode{ConID: conversationID, Action: constant.UpdateLatestMessageChange}, c.GetCh())
	}
	_ = common.TriggerCmdUpdateConversation(ctx, common.UpdateConNode{ConID: conversationID, Action: constant.ConChange, Args: []string{conversationID}}, c.GetCh())
	_ = common.TriggerCmdUpdateConversation(ctx, common.UpdateConNode{ConID: conversationID, Action: constant.TotalUnreadMessageChanged}, c.GetCh())
}

func (c *Conversation) doUnreadCount(ctx context.Context, conversationID string, hasReadSeq int64) {
	c.markAsReadLock.Lock()
	defer c.markAsReadLock.Unlock()
	conversation, err := c.db.GetConversation(ctx, conversationID)
	if err != nil {
		log.ZError(ctx, "GetConversation err", err, "conversationID", conversationID)
		return
	}
	var seqs []int64
	if hasReadSeq > conversation.HasReadSeq {
		for i := conversation.HasReadSeq + 1; i <= hasReadSeq; i++ {
			seqs = append(seqs, i)
		}
		decrCount, err := c.db.MarkConversationMessageAsReadBySeqs(ctx, conversationID, seqs)
		if err != nil {
			log.ZError(ctx, "MarkConversationMessageAsReadBySeqs err", err, "conversationID", conversationID, "seqs", seqs)
			return
		}
		if err := c.db.DecrConversationUnreadCount(ctx, conversationID, decrCount); err != nil {
			log.ZError(ctx, "decrConversationUnreadCount err", err, "conversationID", conversationID, "decrCount", decrCount)
		}
		if err := c.db.UpdateColumnsConversation(ctx, conversationID, map[string]interface{}{"has_read_seq": hasReadSeq}); err != nil {
			log.ZError(ctx, "UpdateColumnsConversation err", err, "conversationID", conversationID)
		}
	} else {
		log.ZWarn(ctx, "hasReadSeq <= conversation.HasReadSeq", nil, "hasReadSeq", hasReadSeq, "conversation.HasReadSeq", conversation.HasReadSeq)
	}
}

func (c *Conversation) doReadDrawing(ctx context.Context, conversationID, userID string, clientMsgIDs []string) {
	// c.msgListener.OnRecvGroupReadReceipt(utils.StructToJsonString(groupMessageReceiptResp))
	// c.db.UpdateColumnsConversation(ctx, conversationID, map[string]interface{}{"has_read_seq": hasReadSeq})
	// messages, err := c.db.GetMultipleMessageController(newMsgID, groupID, sessionType)
	// if err != nil {
	// 	log.Error("internal", "GetMessage err:", err.Error(), "ClientMsgID", newMsgID)
	// 	continue
	// }
	// msgRt := new(sdk_struct.MessageReceipt)
	// msgRt.UserID = userID
	// msgRt.GroupID = groupID
	// msgRt.SessionType = sessionType
	// msgRt.ContentType = constant.GroupHasReadReceipt

	// for _, message := range messages {
	// 	t := new(model_struct.LocalChatLog)
	// 	if userID != c.loginUserID {
	// 		attachInfo := sdk_struct.AttachedInfoElem{}
	// 		_ = utils.JsonStringToStruct(message.AttachedInfo, &attachInfo)
	// 		attachInfo.GroupHasReadInfo.HasReadUserIDList = utils.RemoveRepeatedStringInList(append(attachInfo.GroupHasReadInfo.HasReadUserIDList, userID))
	// 		attachInfo.GroupHasReadInfo.HasReadCount = int32(len(attachInfo.GroupHasReadInfo.HasReadUserIDList))
	// 		t.AttachedInfo = utils.StructToJsonString(attachInfo)
	// 	}
	// 	t.ClientMsgID = message.ClientMsgID
	// 	t.IsRead = true
	// 	t.SessionType = message.SessionType
	// 	t.RecvID = message.RecvID
	// 	err1 := c.db.UpdateMessageController(t)
	// 	if err1 != nil {
	// 		log.Error("internal", "setGroupMessageHasReadByMsgID err:", err1, "ClientMsgID", t, message)
	// 		continue
	// 	}
	// 	successMsgIDlist = append(successMsgIDlist, message.ClientMsgID)
	// }
}