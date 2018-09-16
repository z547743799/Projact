package data

type TMineCaveConfig struct {
	Id          int32   //矿洞Id
	CaveQuality int32   //矿洞品质
	CaveName    int32   //名称
	CaveCount   int32   //资源矿星数
	MineId      []int32 //资源矿Id
	CaveRegion  int32   //矿洞区域
	CaveBonus   int32   //矿洞奖励
	Man         int32   //当前占矿人数
	Guild       int32   //公会
}
