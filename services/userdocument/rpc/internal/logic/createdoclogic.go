package logic

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"chongsheng.art/wesearch/internal/message"
	"chongsheng.art/wesearch/services/userdocument/model"
	"chongsheng.art/wesearch/services/userdocument/rpc/internal/svc"
	"chongsheng.art/wesearch/services/userdocument/rpc/pb"
)

type CreateDocLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDocLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDocLogic {
	return &CreateDocLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDocLogic) CreateDoc(in *pb.DocumentCollectReq) (*pb.DocumentCollectResp, error) {
	err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		user, err := l.findOrCreateUser(in.GetWxUID(), session)
		if err != nil {
			return err
		}

		// 更新 updated_at
		if err = l.updateTime(user, session); err != nil {
			return err
		}

		// 创建document并绑定关系
		docParam := &model.Documents{
			Url:  in.GetURL(),
			Hash: calculateMD5(in.GetURL()),
		}
		newDoc, err := l.findOrCreateDoc(user.Id, session, docParam)
		if err != nil {
			return err
		}

		// 写入消息队列，数据库中不保存content
		msg, err := message.BuildDocMsg(&message.DocCollection{DocID: newDoc.Id, URL: in.GetURL()})
		if err != nil {
			return err
		}
		return l.svcCtx.Producer.Send(l.svcCtx.Config.Kafka.TopicCreateDoc, msg)
	},
	)

	if err != nil {
		return nil, err
	}
	return &pb.DocumentCollectResp{}, nil
}

func (l *CreateDocLogic) findOrCreateUser(openID string, session sqlx.Session) (*model.Users, error) {
	user, err := l.svcCtx.UserModel.FindOneByOpenid(l.ctx, session, openID)
	if err == model.ErrNotFound {
		user = buildNewUser(openID)
		res, err := l.svcCtx.UserModel.Insert(l.ctx, session, user)
		if err != nil {
			return nil, err
		}
		user.Id, _ = res.LastInsertId()
		return user, nil
	}
	if err != nil {
		return nil, err
	}
	return user, err
}

func (l *CreateDocLogic) updateTime(user *model.Users, session sqlx.Session) error {
	user.UpdatedAt = time.Now()
	_, err := l.svcCtx.UserModel.UpdateTimeByID(l.ctx, session, user)
	return err
}

func (l *CreateDocLogic) findOrCreateDoc(userID int64, session sqlx.Session, docParam *model.Documents) (*model.Documents, error) {
	docs, err := l.svcCtx.DocModel.FindByUrlHash(l.ctx, session, docParam.Hash)
	if err != nil {
		return nil, err
	}

	var create bool

	if len(docs) == 0 {
		create = true
	} else {
		docParam, err = l.svcCtx.DocModel.FindOneByUrl(l.ctx, session, docParam.Url)
		if err != nil && err != model.ErrNotFound {
			return nil, err
		}
		if err == model.ErrNotFound {
			create = true
		}
	}

	if create {
		docParam.CreatedAt = time.Now()
		rs, err := l.svcCtx.DocModel.Insert(l.ctx, session, docParam)
		if err != nil {
			return nil, err
		}
		docParam.Id, _ = rs.LastInsertId()
	}

	// 该用户已收藏该文档
	oldDoc, err := l.svcCtx.DocModel.FindOneByUIDAndDocID(l.ctx, session, userID, docParam.Id)
	if err != nil && err != sqlx.ErrNotFound {
		log.Printf("query error, %v\n", err)
	} else if err == nil {
		log.Printf("user %d already collected doc %d\n", userID, docParam.Id)
		return oldDoc, nil
	}
	// not found

	userDoc := &model.UserDocs{
		Uid:       userID,
		DocId:     docParam.Id,
		CreatedAt: time.Now(),
	}
	_, err = l.svcCtx.UserDocModel.Insert(l.ctx, session, userDoc)
	if err != nil {
		return nil, err
	}
	return docParam, nil
}

func buildNewUser(openID string) *model.Users {
	now := time.Now()
	return &model.Users{
		Openid:    openID,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func calculateMD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum(hash.Md5([]byte(str))))
}
