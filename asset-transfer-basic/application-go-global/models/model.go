package models

type Asset struct {
	CertNo    string `form:"certNo" json:"certNo" xml:"certNo"  binding:"required"`
	ID        string `form:"id" json:"id" xml:"color"  binding:"required"`
	Name      string `form:"name" json:"name" xml:"name"  binding:"required"`
	Brand     string `form:"brand" json:"brand" xml:"brand"  binding:"required"`
	NumOfDose string `form:"numOfDose" json:"numOfDose" xml:"numOfDose"  binding:"required"`
	Time      string `form:"time" json:"time" xml:"time"  binding:"required"`
	Issuer    string `form:"issuer" json:"issuer" xml:"issuer"  binding:"required"`
	Remark    string `form:"remark" json:"remark" xml:"remark"  binding:""`
}

type InputInfo struct {
	CertDetail Asset  `form:"CertDetail" json:"CertDetail" xml:"CertDetail"  binding:"required"`
	PersonHash string `form:"PersonHash" json:"PersonHash" xml:"PersonHash"  binding:"required"`
}

type VerifyInfo struct {
	CertNo       string `form:"certNo" json:"certNo" xml:"certNo"  binding:"required"`
	LocalChainID string `form:"localChainID" json:"localChainID" xml:"localChainID"  binding:"required"`
	PersonHash   string `form:"personHash" json:"personHash" xml:"personHash"  binding:"required"`
}

type LocalChainInfo struct {
	LocalChainID         string `form:"localChainID" json:"localChainID" xml:"localChainID"  binding:""`
	MerkelTreePathDetail string `form:"merkelTreePathDetail" json:"merkelTreePathDetail" xml:"merkelTreePathDetail"  binding:"required"`
	LocalChainTxHash     string `form:"localChainTxHash" json:"localChainTxHash" xml:"localChainTxHash"  binding:""`
	LocalChainBlockNum   int64  `form:"localChainBlockNum" json:"localChainBlockNum" xml:"localChainBlockNum"  binding:""`
	LocalChainTimeStamp  int64  `form:"localChainTimeStamp" json:"localChainTimeStamp" xml:"localChainTimeStamp"  binding:""`
}

type GlocalChainInfo struct {
	CertIDList           string `form:"certIDList" json:"certIDList" xml:"certIDList"  binding:"required"`
	MerkelTreeRoot       string `form:"merkelTreeRoot" json:"merkelTreeRoot" xml:"merkelTreeRoot"  binding:"required"`
	GlobalChainTxHash    string `form:"globalChainTxHash" json:"globalChainTxHash" xml:"globalChainTxHash"  binding:""`
	GlobalChainBlockNum  int64  `form:"globalChainBlockNum" json:"globalChainBlockNum" xml:"globalChainBlockNum"  binding:""`
	GlobalChainTimeStamp int64  `form:"globalChainTimeStamp" json:"globalChainTimeStamp" xml:"globalChainTimeStamp"  binding:""`
}
type MerkelTreePath struct {
	GlobalID    string
	CurrentHash string
	Path        []string
	Indexes     []int64
}

type GetPathInfo struct {
	Message          string `form:"message" json:"message" xml:"message"  binding:"required"`
	AssetDetail      Asset  `form:"asset" json:"asset" xml:"asset"  binding:"required"`
	PersonHash       string `form:"personHash" json:"personHash" xml:"personHash"  binding:"required"`
	MKTreePathDetail string `form:"path" json:"path" xml:"path"  binding:"required"`
}
