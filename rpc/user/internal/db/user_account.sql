CREATE TABLE `user_account`
(
    `id`      BIGINT AUTO_INCREMENT COMMENT '主键id',
    `user_id` BIGINT       NOT NULL COMMENT '用户id',
    `balance` VARCHAR(225) NOT NULL COMMENT '余额',
    `loan`    VARCHAR(255) NULL COMMENT '贷款',
    CONSTRAINT user_account_user_id_uindex
        UNIQUE (`user_id`),
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT '用户账户表';
