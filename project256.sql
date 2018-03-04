# 文章表
CREATE TABLE `essay` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `essay_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '文章id',
  `essay_title` varchar(32) NOT NULL DEFAULT '' COMMENT '文章标题',
  `essay_content` text NOT NULL COMMENT '文章内容',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态',
  `create_user` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `essay_id` (`essay_id`),
  KEY `create_user` (`create_user`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表';

# 心愿表
CREATE TABLE `wish` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `wish_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '心愿id',
  `wish_content` text NOT NULL COMMENT '心愿内容',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态',
  `create_user` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `wish_id` (`wish_id`),
  KEY `create_user` (`create_user`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='心愿表';

# 子心愿表
CREATE TABLE `little_wish` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `little_wish_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '子心愿id',
  `little_wish_content` text NOT NULL COMMENT '心愿内容',
  `wish_id`  bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '心愿id',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态',
  `create_user` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `little_wish_id` (`little_wish_id`),
  KEY `wish_id` (`wish_id`),
  KEY `create_user` (`create_user`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='子心愿表';

# 心愿文章表
CREATE TABLE `re_wish_essay` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `wish_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '心愿id',
  `little_wish_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '子心愿id',
  `essay_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '文章id',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT '状态',
  `create_user` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建人',
  `create_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `wish_id` (`wish_id`),
  KEY `little_wish_id` (`little_wish_id`),
  KEY `essay_id` (`essay_id`),
  KEY `create_user` (`create_user`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='心愿文章关联表';

