CREATE DATABASE banking;
USE banking;

DROP TABLE IF EXISTS customers
CREATE TABLE customers (
    `customer_id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `date_of_birth` date NOT NULL,
    `city` varchar(100) NOT NULL,
    `zip_code` varchar(100) NOT NULL,
    `status` tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `customers` VALUES 
(2000, 'Amir', '1991-08-15', 'Qazvin', 'BN11AA', 1),
(2001, 'Sana', '1991-08-15', 'Qazvin', 'BN11AA', 1),
(2002, 'Radmehr', '1991-08-15', 'Qazvin', 'BN11AA', 1),
(2003, 'Ali', '1991-08-15', 'Qazvin', 'BN11AA', 0),
(2004, 'Javad', '1991-08-15', 'Qazvin', 'BN11AA', 1);


DROP TABLE IF EXISTS `accounts`
CREATE TABLE `accounts` (
    `account_id` int(11) NOT NULL AUTO_INCREMENT,
    `customer_id` int(11) NOT NULL ,
    `opening_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `account_type` varchar(10) NOT NULL,
    `pin` varchar(10) NOT NULL,
    `status` tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY (`account_id`),
    KEY `accounts_FK` (`customer_id`),
    CONSTRAINT `accounts_FK` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`)
)ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;

INSERT INTO `accounts` VALUES 
(2000, 2000, '1991-08-15', 'saving', '123', 1),
(2001, 2001, '1991-08-15', 'saving', '2323', 1),
(2002, 2002, '1991-08-15', 'saving', '345', 1),
(2003, 2003, '1991-08-15', 'checking', '345', 1),
(2004, 2004, '1991-08-15', 'checking', '454', 0);


DROP TABLE IF EXISTS `transactions`
CREATE TABLE `transactions` (
    `transaction_id` int(11) NOT NULL AUTO_INCREMENT,
    `account_id` int(11) NOT NULL ,
    `amount` int(11) NOT NULL ,
    `transaction_type` varchar(10) NOT NULL,
    `transaction_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`transaction_id`),
    KEY `transaction_FK` (`account_id`),
    CONSTRAINT `transaction_FK` FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`)
)ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;
















----------------------------------------------------------------------------------------







CREATE DATABASE banking;
USE banking;

CREATE TABLE customers (
	customer_id int(11) AUTO_INCREMENT,
	name VARCHAR(20), city VARCHAR(20),
       zip_code VARCHAR(20), date_of_birth DATE, status tinyint(1) DEFAULT '1',PRIMARY KEY (customer_id));
       
CREATE TABLE accounts (
	account_id int(11)  AUTO_INCREMENT,
	customer_id int(11) NOT NULL ,
    opening_date datetime DEFAULT CURRENT_TIMESTAMP,
    account_type varchar(10),
    pin varchar(10),
    status tinyint(1) DEFAULT '1',
   PRIMARY KEY (account_id),
   CONSTRAINT `accounts_FK` FOREIGN KEY (customer_id) REFERENCES customers (customer_id)
  
   );
       
CREATE TABLE transactions ( 
	transaction_id int(11) AUTO_INCREMENT,
	account_id int(11),
	amount int(11),
    transaction_type varchar(10),
    transaction_date datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (transaction_id),
    CONSTRAINT `transaction_FK` FOREIGN KEY (account_id) REFERENCES accounts (account_id)
    );