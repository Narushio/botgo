package apitest

import (
	"log"
	"testing"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/websocket"
)

func Test_websocket(t *testing.T) {
	ws, err := api.WS(ctx, nil, "")
	log.Printf("%+v, err:%v", ws, err)

	t.Run(
		"at message", func(t *testing.T) {
			var message websocket.ATMessageEventHandler = func(event *dto.WSPayload, data *dto.WSATMessageData) error {
				log.Println(event, data)
				return nil
			}
			intent := websocket.RegisterHandlers(message)
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
	t.Run(
		"at message assign shard to 2", func(t *testing.T) {
			var message websocket.ATMessageEventHandler = func(event *dto.WSPayload, data *dto.WSATMessageData) error {
				log.Println(event, data)
				return nil
			}
			ws.Shards = 2
			intent := websocket.RegisterHandlers(message)
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
	t.Run(
		"at message and guild event", func(t *testing.T) {
			var message websocket.ATMessageEventHandler = func(event *dto.WSPayload, data *dto.WSATMessageData) error {
				log.Println(event, data)
				return nil
			}
			var guildEvent websocket.GuildEventHandler = func(event *dto.WSPayload, data *dto.WSGuildData) error {
				log.Println(event, data)
				return nil
			}
			intent := websocket.RegisterHandlers(message, guildEvent)
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
	t.Run(
		"message reaction", func(t *testing.T) {
			var message websocket.MessageReactionEventHandler = func(
				event *dto.WSPayload, data *dto.WSMessageReactionData,
			) error {
				log.Println(event, data)
				return nil
			}
			intent := websocket.RegisterHandlers(message)
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
	t.Run(
		"thread event", func(t *testing.T) {
			var message websocket.ThreadEventHandler = func(
				event *dto.WSPayload, data *dto.WSThreadData,
			) error {
				log.Println(event, data)
				return nil
			}
			intent := websocket.RegisterHandlers(message)
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
	t.Run(
		"post event", func(t *testing.T) {
			var message websocket.PostEventHandler = func(
				event *dto.WSPayload, data *dto.WSPostData,
			) error {
				log.Println(event, data)
				return nil
			}
			intent := websocket.RegisterHandlers(message)
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
	t.Run(
		"Reply event", func(t *testing.T) {
			var message websocket.ReplyEventHandler = func(
				event *dto.WSPayload, data *dto.WSReplyData,
			) error {
				log.Println(event, data)
				return nil
			}
			intent := websocket.RegisterHandlers(message)
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
	t.Run(
		"Forum audit event", func(t *testing.T) {
			var message websocket.ForumAuditEventHandler = func(
				event *dto.WSPayload, data *dto.WSForumAuditData,
			) error {
				log.Println(event, data)
				return nil
			}
			intent := websocket.RegisterHandlers(message)
			botgo.NewSessionManager().Start(ws, botToken, &intent)
		},
	)
}
