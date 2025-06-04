package Friend

type DefaultParam struct {
	Wxid   string
	ToWxid string
}

type BlacklistParam struct {
	Wxid   string
	ToWxid string
	Val    uint32
}

type BatchDelFriendParam struct {
	Wxid    string
	ToWxids []string
}
