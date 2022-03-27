CREATE TABLE `localCertificate` (
    `id` int(20) NOT NULL AUTO_INCREMENT,
    `certID` VARCHAR(255) DEFAULT '',
    `personID` VARCHAR(255) DEFAULT '',
    `name` VARCHAR(255) DEFAULT '',
    `brand` VARCHAR(255) DEFAULT '',
    `numOfDose` int DEFAULT 0,
    `issueTime` DATETIME ,
    `issuer` VARCHAR(255) DEFAULT '',
    `remark` VARCHAR(255) DEFAULT '',

    `localChainID` VARCHAR(255) DEFAULT '',
    `localChainTxHash`VARCHAR(255) DEFAULT '',
    `localChainBlockNum` int DEFAULT 0,
    `localChainTimeStamp` DateTime, 

    UNIQUE KEY(`certID`),
    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `globalChainInfo` (
    `id` int(20) NOT NULL AUTO_INCREMENT,
    `certIDList` TEXT ,

    `globalChainTxHash`VARCHAR(255) DEFAULT '',
    `globalChainBlockNum` int DEFAULT 0,
    `globalChainTimeStamp` DateTime, 

    PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

