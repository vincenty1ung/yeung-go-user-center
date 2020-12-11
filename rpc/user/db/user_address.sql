CREATE TABLE `user_address`
(
    `id`      BIGINT AUTO_INCREMENT COMMENT '主键id',
    `user_id` BIGINT            NOT NULL COMMENT '用户id',
    `address` VARCHAR(2000)     NULL COMMENT '地址',
    `status`  TINYINT NOT NULL COMMENT '0:默认1:非默认',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT '用户地址表';

CREATE index user_address_user_id_index
    ON user_address (`user_id`);

