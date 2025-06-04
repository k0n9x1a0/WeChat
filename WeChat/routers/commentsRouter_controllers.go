package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["wechatdll/controllers:FavorController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FavorController"],
		beego.ControllerComments{
			Method:           "Del",
			Router:           `/Del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FavorController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FavorController"],
		beego.ControllerComments{
			Method:           "GetFavInfo",
			Router:           `/GetFavInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FavorController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FavorController"],
		beego.ControllerComments{
			Method:           "GetFavItem",
			Router:           `/GetFavItem`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FavorController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FavorController"],
		beego.ControllerComments{
			Method:           "Sync",
			Router:           `/Sync`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FinderController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FinderController"],
		beego.ControllerComments{
			Method:           "UserPrepare",
			Router:           `/UserPrepare`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "Comment",
			Router:           `/Comment`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "GetDetail",
			Router:           `/GetDetail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "GetIdDetail",
			Router:           `/GetIdDetail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "GetList",
			Router:           `/GetList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "Messages",
			Router:           `/Messages`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "MmSnsSync",
			Router:           `/MmSnsSync`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "Operation",
			Router:           `/Operation`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "PrivacySettings",
			Router:           `/PrivacySettings`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendCircleController"],
		beego.ControllerComments{
			Method:           "Upload",
			Router:           `/Upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "Blacklist",
			Router:           `/Blacklist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/Delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "BatchDelFriend",
			Router:           `/BatchDelFriend`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "GetContractDetail",
			Router:           `/GetContractDetail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "GetContractList",
			Router:           `/GetContractList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "GetMFriend",
			Router:           `/GetMFriend`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// 检测单线好友
	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "CheckSingleFriend",
			Router:           `/CheckSingleFriend`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "PassVerify",
			Router:           `/PassVerify`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "Search",
			Router:           `/Search`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "GetFriendRelation",
			Router:           `/GetFriendRelation`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "SendRequest",
			Router:           `/SendRequest`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "SetRemarks",
			Router:           `/SetRemarks`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:FriendController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:FriendController"],
		beego.ControllerComments{
			Method:           "Upload",
			Router:           `/Upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "AddChatRoomMember",
			Router:           `/AddChatRoomMember`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "ConsentToJoin",
			Router:           `/ConsentToJoin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "CreateChatRoom",
			Router:           `/CreateChatRoom`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "FacingCreateChatRoom",
			Router:           `/FacingCreateChatRoom`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "DelChatRoomMember",
			Router:           `/DelChatRoomMember`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "GetChatRoomInfo",
			Router:           `/GetChatRoomInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "GetChatRoomInfoDetail",
			Router:           `/GetChatRoomInfoDetail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "GetChatRoomMemberDetail",
			Router:           `/GetChatRoomMemberDetail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "GetQRCode",
			Router:           `/GetQRCode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "InviteChatRoomMember",
			Router:           `/InviteChatRoomMember`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "MoveContractList",
			Router:           `/MoveContractList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "OperateChatRoomAdmin",
			Router:           `/OperateChatRoomAdmin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "Quit",
			Router:           `/Quit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "ScanIntoGroup",
			Router:           `/ScanIntoGroup`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "ScanIntoGroupEnterprise",
			Router:           `/ScanIntoGroupEnterprise`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "SetChatRoomAnnouncement",
			Router:           `/SetChatRoomAnnouncement`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "SetChatRoomName",
			Router:           `/SetChatRoomName`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:GroupController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:GroupController"],
		beego.ControllerComments{
			Method:           "SetChatRoomRemarks",
			Router:           `/SetChatRoomRemarks`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LabelController"],
		beego.ControllerComments{
			Method:           "Add",
			Router:           `/Add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LabelController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/Delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LabelController"],
		beego.ControllerComments{
			Method:           "GetList",
			Router:           `/GetList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LabelController"],
		beego.ControllerComments{
			Method:           "UpdateList",
			Router:           `/UpdateList`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LabelController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LabelController"],
		beego.ControllerComments{
			Method:           "UpdateName",
			Router:           `/UpdateName`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Data62Login",
			Router:           `/62data`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Data62SMSApply",
			Router:           `/62dataSMSApply`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Data62SMSAgain",
			Router:           `/62dataSMSAgain`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Data62SMSVerify",
			Router:           `/62dataSMSVerify`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Data62QRCodeApply",
			Router:           `/62dataQrApply`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "A16Data",
			Router:           `/A16Data`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LoginAwaken",
			Router:           `/Awaken`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// 二维码检测
	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LoginCheckQR",
			Router:           `/CheckQR`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "ExtDeviceLoginConfirmGet",
			Router:           `/ExtDeviceLoginConfirmGet`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "ExtDeviceLoginConfirmOk",
			Router:           `/ExtDeviceLoginConfirmOk`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Get62Data",
			Router:           `/Get62Data`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "GetCacheInfo",
			Router:           `/GetCacheInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// 获取二维码
	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LoginGetQR",
			Router:           `/GetQR`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "HeartBeat",
			Router:           `/HeartBeat`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	/*    beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
	      beego.ControllerComments{
	          Method: "HeartBeatLong",
	          Router: `/HeartBeatLong`,
	          AllowHTTPMethods: []string{"post"},
	          MethodParams: param.Make(),
	          Filters: nil,
	          Params: nil})*/

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LogOut",
			Router:           `/LogOut`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Newinit",
			Router:           `/Newinit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:LoginController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LoginTwiceAutoAuth",
			Router:           `/TwiceAutoAuth`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "Revoke",
			Router:           `/Revoke`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "SendApp",
			Router:           `/SendApp`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "SendCDNFile",
			Router:           `/SendCDNFile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "SendCDNImg",
			Router:           `/SendCDNImg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "SendCDNVideo",
			Router:           `/SendCDNVideo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "SendEmoji",
			Router:           `/SendEmoji`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "SendTxt",
			Router:           `/SendTxt`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "SendVideo",
			Router:           `/SendVideo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "SendVoice",
			Router:           `/SendVoice`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "ShareCard",
			Router:           `/ShareCard`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "ShareLink",
			Router:           `/ShareLink`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "ShareLocation",
			Router:           `/ShareLocation`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "ShareVideo",
			Router:           `/ShareVideo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "Sync",
			Router:           `/Sync`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:MsgController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:MsgController"],
		beego.ControllerComments{
			Method:           "UploadImg",
			Router:           `/UploadImg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// 关注公众号
	beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"],
		beego.ControllerComments{
			Method:           "Follow",
			Router:           `/Follow`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"],
		beego.ControllerComments{
			Method:           "GetAppMsgExt",
			Router:           `/GetAppMsgExt`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"],
		beego.ControllerComments{
			Method:           "GetAppMsgExtLike",
			Router:           `/GetAppMsgExtLike`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"],
		beego.ControllerComments{
			Method:           "JSAPIPreVerify",
			Router:           `/JSAPIPreVerify`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"],
		beego.ControllerComments{
			Method:           "MpGetA8Key",
			Router:           `/MpGetA8Key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"],
		beego.ControllerComments{
			Method:           "OauthAuthorize",
			Router:           `/OauthAuthorize`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:OfficialAccountsController"],
		beego.ControllerComments{
			Method:           "Quit",
			Router:           `/Quit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:SayHelloController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:SayHelloController"],
		beego.ControllerComments{
			Method:           "ModelV1",
			Router:           `/Modelv1`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:SayHelloController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:SayHelloController"],
		beego.ControllerComments{
			Method:           "Modelv2",
			Router:           `/Modelv2`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "SwitchAccountService",
			Router:           `/SwitchAccountService`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "DownloadFile",
			Router:           `/DownloadFile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "DownloadImg",
			Router:           `/DownloadImg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "CdnDownloadImage",
			Router:           `/CdnDownloadImage`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "DownloadVideo",
			Router:           `/DownloadVideo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "DownloadVoice",
			Router:           `/DownloadVoice`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "DownloadVoicePlus",
			Router:           `/DownloadVoicePlus`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "GeneratePayQCode",
			Router:           `/GeneratePayQCode`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "GetA8Key",
			Router:           `/GetA8Key`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "GetCdnDns",
			Router:           `/GetCdnDns`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "GetBandCardList",
			Router:           `/GetBandCardList`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// 文件上传
	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "UploadFile",
			Router:           `/UploadFile`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "GetAccountInfo",
			Router:           `/GetAccountInfo`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "GetBoundHardDevices",
			Router:           `/GetBoundHardDevices`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "ThirdAppGrant",
			Router:           `/ThirdAppGrant`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// 修改微信步数
	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "UpdateStepNumberApi",
			Router:           `/UpdateStep`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:ToolsController"],
		beego.ControllerComments{
			Method:           "SetProxy",
			Router:           `/setproxy`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "BindingEmail",
			Router:           `/BindingEmail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "BindingMobile",
			Router:           `/BindingMobile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "BindQQ",
			Router:           `/BindQQ`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "DelSafetyInfo",
			Router:           `/DelSafetyInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetContractProfile",
			Router:           `/GetContractProfile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetQRCode",
			Router:           `/GetQRCode`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetSafetyInfo",
			Router:           `/GetSafetyInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "PrivacySettings",
			Router:           `/PrivacySettings`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "ReviseMotion",
			Router:           `/ReportMotion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SendVerifyMobile",
			Router:           `/SendVerifyMobile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SetAlisa",
			Router:           `/SetAlisa`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SetPasswd",
			Router:           `/SetPasswd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdateProfile",
			Router:           `/UpdateProfile`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UploadHeadImage",
			Router:           `/UploadHeadImage`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:UserController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:UserController"],
		beego.ControllerComments{
			Method:           "VerifyPasswd",
			Router:           `/VerifyPasswd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:WxappController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:WxappController"],
		beego.ControllerComments{
			Method:           "Updatewxausagerecord",
			Router:           `/updatewxausagerecord`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
	beego.GlobalControllerRouter["wechatdll/controllers:WxappController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:WxappController"],
		beego.ControllerComments{
			Method:           "JSLogin",
			Router:           `/JSLogin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:WxappController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:WxappController"],
		beego.ControllerComments{
			Method:           "JSOperateWxData",
			Router:           `/JSOperateWxData`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"],
		beego.ControllerComments{
			Method:           "Receivewxhb",
			Router:           `/Receivewxhb`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"],
		beego.ControllerComments{
			Method:           "Openwxhb",
			Router:           `/Openwxhb`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// 查看微信红包
	beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"],
		beego.ControllerComments{
			Method:           "Qrydetailwxhb",
			Router:           `/Qrydetailwxhb`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	//自动抢红包
	beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"],
		beego.ControllerComments{
			Method:           "AutoHongBao",
			Router:           `/AutoHongBao`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// 红包列表
	beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"] = append(beego.GlobalControllerRouter["wechatdll/controllers:TenPayController"],
		beego.ControllerComments{
			Method:           "GetRedPacketListApi",
			Router:           `/GetRedPacketListApi`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
