SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `shop_user`;
DROP TABLE IF EXISTS `shop_product`;
DROP TABLE IF EXISTS `shop_product_img`;
DROP TABLE IF EXISTS `shop_category`;
DROP TABLE IF EXISTS `shop_product_info`;
DROP TABLE IF EXISTS `shop_order`;
DROP TABLE IF EXISTS `shop_order_item`;
DROP TABLE IF EXISTS `shop_user_recommend`;
SET FOREIGN_KEY_CHECKS = 1;

CREATE TABLE `shop_user` (
    `id` BIGINT(20) NOT NULL,
    `union_id` CHAR(20) NOT NULL,
    `nick` CHAR(20) NOT NULL,
    `phone` CHAR(11) NOT NULL,
    `balance` DECIMAL NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE (`union_id`)
);

CREATE TABLE `shop_product` (
    `id` BIGINT(20) NOT NULL,
    `category_id` BIGINT(20) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `price` DECIMAL NOT NULL,
    `img` VARCHAR(255) NOT NULL,
    `inventory` INTEGER(10) NOT NULL,
    `unit_name` CHAR(5) NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `shop_product_img` (
    `product_id` BIGINT(20) NOT NULL,
    `url` INTEGER NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL
);

CREATE TABLE `shop_category` (
    `id` BIGINT(20) NOT NULL,
    `name` VARCHAR(50) NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `shop_product_info` (
    `product_id` BIGINT(20) NOT NULL,
    `desc` TEXT NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL
);

CREATE TABLE `shop_order` (
    `id` BIGINT(20) NOT NULL,
    `orderNo` VARCHAR(32) NOT NULL,
    `paid` INTEGER(2) NOT NULL,
    `status` INTEGER(2) NOT NULL,
    `user_id` BIGINT(20) NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `shop_order_item` (
    `id` BIGINT(20) NOT NULL,
    `order_id` BIGINT(20) NOT NULL,
    `product_id` BIGINT(20) NOT NULL,
    `product_name` VARCHAR(50) NOT NULL,
    `product_price` DECIMAL NOT NULL,
    `num` INTEGER(10) NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `shop_user_recommend` (
    `id` BIGINT(20) NOT NULL,
    `user_id` BIGINT(20) NOT NULL,
    `recommend_user_id` BIGINT(20) NOT NULL,
    `created` DATETIME NOT NULL,
    `updated` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);
