package net

import ()

type APiUser struct {
}

func (u APiUser) Default() {

}

func (u APiUser) Login(packet []byte) bool {

	return false
}

func (u APiUser) ReLogin() bool {
	return false
}

func (u APiUser) Logout() bool {
	return false
}

func (u APiUser) Signup() bool {
	return false
}

func init() {
	//user := APiUser{}
	//api["user"] = user
}

//func Login(packet []byte) {
//
//	req := &pb.LoginReq{}
//	err := proto.Unmarshal(packet, req)
//	if err != nil {
//		util.LogErr(err, zap.String("reason", "proto.Unmarshal Login"))
//		return
//	}
//	util.LogInfo("Loginreq", zap.String("proto", req.String()))
//	//log2.app.Login(req)
//}
//
//func Logout(packet []byte) {
//	log.Printf("Logout req")
//	req := &pb.LogoutReq{}
//	err := proto.Unmarshal(packet, req)
//	if err != nil {
//		util.LogErr(err, zap.String("reason", "proto.Unmarshal Logout"))
//	}
//	util.LogInfo("Logoutreq", zap.String("proto", req.String()))
//	//log2.app.Logout(req)
//}
