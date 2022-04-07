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
	CertDetail     Asset  `form:"CertDetail" json:"CertDetail" xml:"CertDetail"  binding:"required"`
	PersonInfoHash string `form:"PersonInfoHash" json:"PersonInfoHash" xml:"PersonInfoHash"  binding:"required"`
	Key            string `form:"Key" json:"Key" xml:"Key"  binding:"required"`
}

type LocalChainInfo struct {
	LocalChainID         string `form:"localChainID" json:"localChainID" xml:"localChainID"  binding:""`
	MerkleTreePathDetail string `form:"merkleTreePathDetail" json:"merkleTreePathDetail" xml:"merkleTreePathDetail"  binding:"required"`
	LocalChainTxHash     string `form:"localChainTxHash" json:"localChainTxHash" xml:"localChainTxHash"  binding:""`
	LocalChainBlockNum   int64  `form:"localChainBlockNum" json:"localChainBlockNum" xml:"localChainBlockNum"  binding:""`
	LocalChainTimeStamp  int64  `form:"localChainTimeStamp" json:"localChainTimeStamp" xml:"localChainTimeStamp"  binding:""`
}

type GlobalChainInfo struct {
	CertIDList           string `form:"certIDList" json:"certIDList" xml:"certIDList"  binding:"required"`
	MerkleTreeRoot       string `form:"merkleTreeRoot" json:"merkleTreeRoot" xml:"merkleTreeRoot"  binding:"required"`
	GlobalChainTxHash    string `form:"globalChainTxHash" json:"globalChainTxHash" xml:"globalChainTxHash"  binding:""`
	GlobalChainBlockNum  int64  `form:"globalChainBlockNum" json:"globalChainBlockNum" xml:"globalChainBlockNum"  binding:""`
	GlobalChainTimeStamp int64  `form:"globalChainTimeStamp" json:"globalChainTimeStamp" xml:"globalChainTimeStamp"  binding:""`
}

type MerkleTreePath struct {
	GlobalID    string
	CurrentHash string
	Path        []string
	Indexes     []int64
}
