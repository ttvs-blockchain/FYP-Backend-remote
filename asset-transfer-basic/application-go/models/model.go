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

type LocalChainInfo struct {
	LocalChainID        string `form:"localChainID" json:"localChainID" xml:"localChainID"  binding:"required"`
	LocalChainTxHash    string `form:"localChainTxHash" json:"localChainTxHash" xml:"localChainTxHash"  binding:"required"`
	LocalChainBlockNum  int64  `form:"localChainBlockNum" json:"localChainBlockNum" xml:"localChainBlockNum"  binding:"required"`
	LocalChainTimeStamp int64  `form:"localChainTimeStamp" json:"localChainTimeStamp" xml:"localChainTimeStamp"  binding:"required"`
}

type GlocalChainInfo struct {
	CertIDList           string `form:"localNodeID" json:"localNodeID" xml:"localNodeID"  binding:"required"`
	GlobalChainTxHash    string `form:"globalChainTxHash" json:"globalChainTxHash" xml:"globalChainTxHash"  binding:"required"`
	GlobalChainBlockNum  int64  `form:"globalChainBlockNum" json:"globalChainBlockNum" xml:"globalChainBlockNum"  binding:"required"`
	GlobalChainTimeStamp int64  `form:"globalChainTimeStamp" json:"globalChainTimeStamp" xml:"globalChainTimeStamp"  binding:"required"`
}

type MerkelTreePath struct {
	GlobalID    string
	CurrentHash string
	Path        []string
	Indexes     []int64
}
