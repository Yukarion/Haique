package models

type InlineObject5 struct {

	SessionId string `json:"session_id"`

	// 最新のhaikuからN番目を起点とする（デフォルト0）
	Start int64 `json:"start,omitempty"`

	// 最新のhaikuからM番目を終点とする（デフォルト30）
	Stop int64 `json:"stop,omitempty"`
}
