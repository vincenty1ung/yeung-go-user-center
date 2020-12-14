CREATE TABLE `user`
(
    `id`          BIGINT AUTO_INCREMENT COMMENT '主键id',
    `age`         TINYINT NOT NULL COMMENT '年龄',
    `status`      TINYINT NOT NULL COMMENT '0:初始化1:激活2:注销',
    `phone`       VARCHAR(11)       NULL COMMENT '手机号',
    `create_time` BIGINT            NOT NULL COMMENT '创建时间',
    `update_time` BIGINT            NOT NULL comment '更新时间',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
COMMENT '用户表';