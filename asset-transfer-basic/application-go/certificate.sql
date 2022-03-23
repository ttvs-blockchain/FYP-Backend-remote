CREATE TABLE `certificate` (
    `certid` int(20) NOT NULL AUTO_INCREMENT,
    `certno` VARCHAR(20) DEFAULT '',
    `personid` VARCHAR(20) DEFAULT '',
    `name` VARCHAR(20) DEFAULT '',
    `brand` VARCHAR(20) DEFAULT '',
    `numofdose` VARCHAR(20) DEFAULT '',
    `time` VARCHAR(20) DEFAULT '',
    `issuer` VARCHAR(100) DEFAULT '',
    `remark` VARCHAR(100) DEFAULT '',
    PRIMARY KEY(`certid`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

