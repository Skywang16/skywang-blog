SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for la_album
-- ----------------------------
DROP TABLE IF EXISTS `la_album`;
CREATE TABLE `la_album`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `cid`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类目ID',
    `aid`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
    `uid`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `type`        tinyint(2) UNSIGNED NOT NULL DEFAULT 10 COMMENT '文件类型: [10=图片, 20=视频]',
    `name`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件名称',
    `uri`         varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '文件路径',
    `ext`         varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '文件扩展',
    `size`        int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文件大小',
    `is_delete`   int(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX         `idx_cid`(`cid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_album_cate
-- ----------------------------
DROP TABLE IF EXISTS `la_album_cate`;
CREATE TABLE `la_album_cate`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `pid`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
    `type`        tinyint(2) UNSIGNED NOT NULL DEFAULT 10 COMMENT '类型: [10=图片, 20=视频]',
    `name`        varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
    `is_delete`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: [0=否, 1=是]',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '相册分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_article_category
-- ----------------------------
DROP TABLE IF EXISTS `la_article_category`;
CREATE TABLE `la_article_category`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
    `sort`        smallint(5) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序',
    `is_show`     tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否显示: 0=否, 1=是',
    `is_delete`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章分类表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_article_collect
-- ----------------------------
DROP TABLE IF EXISTS `la_article_collect`;
CREATE TABLE `la_article_collect`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_id`     int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
    `article_id`  int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '文章ID',
    `is_delete`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '文章收藏表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_decorate_page
-- ----------------------------
DROP TABLE IF EXISTS `la_decorate_page`;
CREATE TABLE `la_decorate_page`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `page_type`   tinyint(2) UNSIGNED NOT NULL DEFAULT 10 COMMENT '页面类型',
    `page_name`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '页面名称',
    `page_data`   text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '页面数据',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '页面装修表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_decorate_tabbar
-- ----------------------------
DROP TABLE IF EXISTS `la_decorate_tabbar`;
CREATE TABLE `la_decorate_tabbar`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '导航名称',
    `selected`    varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '未选图标',
    `unselected`  varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '已选图标',
    `link`        varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '链接地址',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '底部装修表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `la_dict_data`;
CREATE TABLE `la_dict_data`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `type_id`     int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型',
    `name`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '键名',
    `value`       varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '数值',
    `remark`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '备注',
    `sort`        smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
    `status`      tinyint(1) NOT NULL COMMENT '状态: 0=停用, 1=正常',
    `is_delete`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `la_dict_type`;
CREATE TABLE `la_dict_type`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `dict_name`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典名称',
    `dict_type`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典类型',
    `dict_remark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典备注',
    `dict_status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '字典状态: 0=停用, 1=正常',
    `is_delete`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_gen_table
-- ----------------------------
DROP TABLE IF EXISTS `la_gen_table`;
CREATE TABLE `la_gen_table`
(
    `id`             int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `table_name`     varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表名称',
    `table_comment`  varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表描述',
    `sub_table_name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关联表名称',
    `sub_table_fk`   varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关联表外键',
    `author_name`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '作者的名称',
    `entity_name`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '实体的名称',
    `module_name`    varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '生成模块名',
    `function_name`  varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '生成功能名',
    `tree_primary`   varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '树主键字段',
    `tree_parent`    varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '树父级字段',
    `tree_name`      varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '树显示字段',
    `gen_tpl`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT 'crud' COMMENT '生成模板方式: [crud=单表, tree=树表]',
    `gen_type`       tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '生成代码方式: [0=zip压缩包, 1=自定义路径]',
    `gen_path`       varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '/' COMMENT '生成代码路径: [不填默认项目路径]',
    `remarks`        varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
    `create_time`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成业务表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS `la_gen_table_column`;
CREATE TABLE `la_gen_table_column`
(
    `id`             int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '列主键',
    `table_id`       int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '表外键',
    `column_name`    varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '列名称',
    `column_comment` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '列描述',
    `column_length`  varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '0' COMMENT '列长度',
    `column_type`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '列类型 ',
    `java_type`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'JAVA类型',
    `java_field`     varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'JAVA字段',
    `is_pk`          tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否主键: [1=是, 0=否]',
    `is_increment`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否自增: [1=是, 0=否]',
    `is_required`    tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否必填: [1=是, 0=否]',
    `is_insert`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否插入字段: [1=是, 0=否]',
    `is_edit`        tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否编辑字段: [1=是, 0=否]',
    `is_list`        tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否列表字段: [1=是, 0=否]',
    `is_query`       tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否查询字段: [1=是, 0=否]',
    `query_type`     varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT 'EQ' COMMENT '查询方式: [等于、不等于、大于、小于、范围]',
    `html_type`      varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '显示类型: [文本框、文本域、下拉框、复选框、单选框、日期控件]',
    `dict_type`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '字典类型',
    `sort`           smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
    `create_time`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '代码生成字段表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_hot_search
-- ----------------------------
DROP TABLE IF EXISTS `la_hot_search`;
CREATE TABLE `la_hot_search`
(
    `id`   int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关键词',
    `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序号',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '热门搜索配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_notice_setting
-- ----------------------------
DROP TABLE IF EXISTS `la_notice_setting`;
CREATE TABLE `la_notice_setting`
(
    `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `scene`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景编号',
    `name`          varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '场景名称',
    `remarks`       varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '场景描述',
    `recipient`     tinyint(1) NOT NULL DEFAULT 1 COMMENT '接收人员: [1=用户, 2=平台]',
    `type`          tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '通知类型: [1=业务, 2=验证]',
    `system_notice` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '系统的通知设置',
    `sms_notice`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '短信的通知设置',
    `oa_notice`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '公众号通知设置',
    `mnp_notice`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '小程序通知设置',
    `is_delete`     tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
    `create_time`   int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`   int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time`   int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '消息通知设置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_official_reply
-- ----------------------------
DROP TABLE IF EXISTS `la_official_reply`;
CREATE TABLE `la_official_reply`
(
    `id`            int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`          varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名',
    `keyword`       varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关键词',
    `reply_type`    tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '回复类型: [1=关注回复 2=关键字回复, 3=默认回复]',
    `matching_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '匹配方式: [1=全匹配, 2=模糊匹配]',
    `content_type`  tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '内容类型: [1=文本]',
    `status`        tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '启动状态: [1=启动, 0=关闭]',
    `content`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '回复内容',
    `sort`          int(11) UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序编号',
    `is_delete`     tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
    `create_time`   int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`   int(10) UNSIGNED NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time`   int(10) UNSIGNED NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '公众号的回复表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_auth_admin
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_admin`;
CREATE TABLE `la_system_auth_admin`
(
    `id`              int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `dept_id`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门ID',
    `username`        varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '用户账号',
    `nickname`        varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '用户昵称',
    `password`        varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码',
    `avatar`          varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci          DEFAULT NULL COMMENT '用户头像',
    `role`            varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色主键',
    `salt`            varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '加密盐巴',
    `sort`            smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
    `is_multipoint`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '多端登录: 0=否, 1=是',
    `is_disable`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `is_delete`       tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `last_login_ip`   varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '最后登录IP',
    `last_login_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最后登录',
    `create_time`     int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`     int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time`     int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统管理成员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_auth_log
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_log`;
CREATE TABLE `la_system_auth_log`
(
    `create_time` DATE NOT NULL COMMENT '创建时间',
    `today_visits` bigint(20) NOT NULL DEFAULT 0 COMMENT '全局当天访问量',
    `total_visits` bigint(20) NOT NULL DEFAULT 0 COMMENT '总访问量',
    `today_users`  bigint(20) NOT NULL DEFAULT 0 COMMENT '全局当天新增用户量',
    `update_time` DATE NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`create_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系统管理成员表' ROW_FORMAT=Dynamic;

-- ----------------------------
-- Table structure for la_system_auth_dept
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_dept`;
CREATE TABLE `la_system_auth_dept`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `pid`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级主键',
    `name`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门名称',
    `duty`        varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '负责人名',
    `mobile`      varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '联系电话',
    `sort`        smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
    `is_stop`     tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `is_delete`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统部门管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_menu`;
CREATE TABLE `la_system_auth_menu`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `pid`         int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级菜单',
    `menu_type`   char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci      NOT NULL DEFAULT '' COMMENT '权限类型: M=目录，C=菜单，A=按钮',
    `menu_name`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
    `menu_icon`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
    `menu_sort`   smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单排序',
    `perms`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '权限标识',
    `paths`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
    `component`   varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '前端组件',
    `selected`    varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '选中路径',
    `params`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由参数',
    `is_cache`    tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否缓存: 0=否, 1=是',
    `is_show`     tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否显示: 0=否, 1=是',
    `is_disable`  tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 775 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统菜单管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_auth_perm
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_perm`;
CREATE TABLE `la_system_auth_perm`
(
    `id`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主键',
    `role_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
    `menu_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单ID',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_auth_role
-- ----------------------------
DROP TABLE IF EXISTS `la_system_auth_role`;
CREATE TABLE `la_system_auth_role`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
    `remark`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
    `sort`        smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色排序',
    `is_disable`  tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统角色管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_config
-- ----------------------------
DROP TABLE IF EXISTS `la_system_config`;
CREATE TABLE `la_system_config`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `type`        varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '类型',
    `name`        varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '键',
    `value`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '值',
    `create_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统全局配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_log_login
-- ----------------------------
DROP TABLE IF EXISTS `la_system_log_login`;
CREATE TABLE `la_system_log_login`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '注解',
    `admin_id`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
    `username`    varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '登录账号',
    `ip`          varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL COMMENT '登录地址',
    `os`          varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
    `browser`     varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器',
    `status`      tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '操作状态: 1=成功, 2=失败',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统登录日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_log_operate
-- ----------------------------
DROP TABLE IF EXISTS `la_system_log_operate`;
CREATE TABLE `la_system_log_operate`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `admin_id`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作人ID',
    `type`        varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '请求类型: GET/POST/PUT',
    `title`       varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作标题',
    `ip`          varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '请求IP',
    `url`         varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求接口',
    `method`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '请求方法',
    `args`        text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数',
    `error`       text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '错误信息',
    `status`      tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '执行状态: 1=成功, 2=失败',
    `start_time`  int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '开始时间',
    `end_time`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '结束时间',
    `task_time`   int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '执行耗时',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统操作日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_system_log_sms
-- ----------------------------
DROP TABLE IF EXISTS `la_system_log_sms`;
CREATE TABLE `la_system_log_sms`
(
    `id`          int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
    `scene`       int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景编号',
    `mobile`      varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '手机号码',
    `content`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '发送内容',
    `status`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送状态：[0=发送中, 1=发送成功, 2=发送失败]',
    `results`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '短信结果',
    `send_time`   int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '发送时间',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统短信日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_article_news
-- ----------------------------
DROP TABLE IF EXISTS `la_article_news`;
CREATE TABLE `la_article_news`
(
    `id`           int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `title`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资讯标题',
    `desc`         varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资讯简介',
    `cid`          varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '资讯作者',
    `image`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '资讯封面',
    `content`      text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资讯内容',
    `sort`         smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
    `status`       tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态: 0=否, 1=是',
    `recommended`  int(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否推荐: 0=否, 1=是',
    `publish_time` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '发布时间',
    `is_stop`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `is_delete`    tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time`  int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`  int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time`  int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '资讯管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_information_weapons
-- ----------------------------
DROP TABLE IF EXISTS `la_information_weapons`;
CREATE TABLE `la_information_weapons`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '武器名称',
    `types`       varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '武器类型',
    `level`       varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '武器等级',
    `image`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '武器图片',
    `price`       varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '武器价格',
    `series`      varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '武器系列',
    `status`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态: 0=否, 1=是',
    `is_stop`     tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `is_delete`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '武器表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_information_weapons_series
-- ----------------------------
DROP TABLE IF EXISTS `la_information_weapons_series`;
CREATE TABLE `la_information_weapons_series`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '系列名称',
    `remark`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
    `sort`        smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '系列排序',
    `is_disable`  tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '武器系列管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_information_hero
-- ----------------------------
DROP TABLE IF EXISTS `la_information_hero`;
CREATE TABLE `la_information_hero`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
    `types`       varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色定位',
    `image`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色图片',
    `status`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态: 0=否, 1=是',
    `is_stop`     tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `is_delete`   tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '武器表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_information_hero_type
-- ----------------------------
DROP TABLE IF EXISTS `la_information_hero_type`;
CREATE TABLE `la_information_hero_type`
(
    `id`          int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '系列名称',
    `remark`      varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注信息',
    `sort`        smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '系列排序',
    `is_disable`  tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '武器系列管理表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for la_information_maps
-- ----------------------------
DROP TABLE IF EXISTS `la_information_maps`;
CREATE TABLE `la_information_maps`
(
    `id`             int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    `name`           varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地图名称',
    `image`          varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地图图片',
    `locations`      varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地图地点',
    `mechanics`      varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地图机制',
    `thumbnail`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '缩略图',
    `plant_quantity` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '包点数量',
    `status`         tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态: 0=否, 1=是',
    `is_stop`        tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否禁用: 0=否, 1=是',
    `is_delete`      tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除: 0=否, 1=是',
    `create_time`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
    `delete_time`    int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '删除时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '武器表' ROW_FORMAT = Dynamic;

SET
FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- INSERT
-- ----------------------------
BEGIN;
INSERT INTO `la_system_auth_dept`
VALUES (1, 0, '默认部门', 'wjj', '19858851348', 10, 0, 0, 1649841995, 1660190949, 0);
INSERT INTO `la_system_auth_admin`
VALUES (1, 0, 'admin', 'admin', '7fac2474740becfaf1ecbdd6cc8fb076', '/api/uploads/image/20242605/79452e14b6dd433585760c4e22c9c2cc.jpg', '0', '5Xar0', 0,
        1, 0, 0, '127.0.0.1', 1660641347, 1642321599, 1660287325, 0);
INSERT INTO `la_system_auth_admin`
VALUES (2, 1, 'user', 'user', 'a7dafd36071a23776a992822e4617b12', '/api/static/backend_avatar.png', '2', 'A5ENl', 1,
        1, 0, 0, '::1', 1712713471, 1712713453, 1712713471, 0);
INSERT INTO `la_system_auth_role`
VALUES (1, '审核员', '审核数据', 999, 0, 1668679451, 1668679468);
INSERT INTO `la_system_auth_role`
VALUES (2, '用户', '用户', 998, 0, 1711078383, 1711078383);
INSERT INTO `la_system_auth_role`
VALUES (3, '测试人员', '测试', 997, 0, 1711078383, 1711078383);
COMMIT;

INSERT INTO `la_system_auth_log` VALUES ('2024-04-10', 2, 2, 0, '2024-04-10');
INSERT INTO `la_system_auth_log` VALUES ('2024-04-11', 9, 11, 0, '2024-04-11');
INSERT INTO `la_system_auth_log` VALUES ('2024-04-12', 24, 25, 1, '2024-04-12');
INSERT INTO `la_system_auth_log` VALUES ('2024-04-13', 10, 36, 0, '2024-04-13');
INSERT INTO `la_system_auth_log` VALUES ('2024-04-14', 2, 37, 0, '2024-04-14');
INSERT INTO `la_system_auth_log` VALUES ('2024-04-15', 22, 58, 0, '2024-04-15');
INSERT INTO `la_system_auth_log` VALUES ('2024-04-17', 18, 77, 0, '2024-04-17');


BEGIN;
INSERT INTO `la_system_config`
VALUES (1, 'storage', 'default', 'local', 1660620367, 1662620927);
INSERT INTO `la_system_config`
VALUES (2, 'storage', 'local', '{\"name\":\"本地存储\"}', 1660620367, 1662620927);
INSERT INTO `la_system_config`
VALUES (3, 'storage', 'qiniu',
        '{\"name\":\"七牛云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', 1660620367,
        1660620367);
INSERT INTO `la_system_config`
VALUES (4, 'storage', 'aliyun',
        '{\"name\":\"阿里云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\"}', 1660620367,
        1662620071);
INSERT INTO `la_system_config`
VALUES (5, 'storage', 'qcloud',
        '{\"name\":\"腾讯云存储\",\"bucket\":\"\",\"secretKey\":\"\",\"accessKey\":\"\",\"domain\":\"\",\"region\":\"\"}',
        1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (6, 'sms', 'default', 'aliyun', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (7, 'sms', 'aliyun',
        '{\"name\":\"阿里云短信\",\"alias\":\"aliyun\",\"sign\":\"\",\"appKey\":\"\",\"secretKey\":\"\"}', 1660620367,
        1660620367);
INSERT INTO `la_system_config`
VALUES (8, 'sms', 'tencent',
        '{\"name\":\"腾讯云短信\",\"alias\":\"tencent\",\"sign\":\"\",\"appId\":\"\",\"secretId\":\"\",\"secretKey\":\"\"}',
        1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (9, 'sms', 'huawei', '{\"name\":\"华为云短信\",\"alias\":\"huawei\"}', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (10, 'website', 'name', '无畏契约后台管理', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (11, 'website', 'logo', '/api/static/logo.png', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (12, 'website', 'favicon', '/api/static/favicon.png', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (13, 'website', 'backdrop', '/api/static/backdrop.png', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (14, 'website', 'copyright', '[{\"name\":\"无畏契约后台管理系统\",\"link\":\"http://www.beian.gov.cn\"}]',
        1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (15, 'website', 'homeBanner', '/api/static/banner0.jpg,/api/static/banner1.jpg', 1631255140, 1631255140);
INSERT INTO `la_system_config`
VALUES (17, 'protocol', 'service', '{\"name\":\"服务协议\",\"content\":\"\"}', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (18, 'protocol', 'privacy', '{\"name\":\"隐私协议\",\"content\":\"\"}', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (19, 'tabbar', 'style', '{\"defaultColor\":\"#4A5DFF\",\"selectedColor\":\"#EA5455\"}', 1660620367, 1662544900);
INSERT INTO `la_system_config`
VALUES (20, 'search', 'isHotSearch', '0', 1660620367, 1662546997);
INSERT INTO `la_system_config`
VALUES (30, 'h5_channel', 'status', '1', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (31, 'h5_channel', 'close', '0', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (32, 'h5_channel', 'url', '', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (40, 'mp_channel', 'name', '', 1660620367, 1662551403);
INSERT INTO `la_system_config`
VALUES (41, 'mp_channel', 'primaryId', '', 1660620367, 1662551403);
INSERT INTO `la_system_config`
VALUES (42, 'mp_channel', 'appId', '', 1660620367, 1662551403);
INSERT INTO `la_system_config`
VALUES (43, 'mp_channel', 'appSecret', '', 1660620367, 1662551403);
INSERT INTO `la_system_config`
VALUES (44, 'mp_channel', 'qrCode', '', 1660620367, 1662551403);
INSERT INTO `la_system_config`
VALUES (50, 'wx_channel', 'appId', '', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (51, 'wx_channel', 'appSecret', '', 1660620367, 1660620367);
INSERT INTO `la_system_config`
VALUES (55, 'oa_channel', 'name', '', 1660620367, 1662551337);
INSERT INTO `la_system_config`
VALUES (56, 'oa_channel', 'primaryId', ' ', 1660620367, 1662551337);
INSERT INTO `la_system_config`
VALUES (57, 'oa_channel', 'qrCode', '', 1662551337, 1662551337);
INSERT INTO `la_system_config`
VALUES (58, 'oa_channel', 'appId', '', 1660620367, 1662551337);
INSERT INTO `la_system_config`
VALUES (59, 'oa_channel', 'appSecret', '', 1660620367, 1662551337);
INSERT INTO `la_system_config`
VALUES (60, 'oa_channel', 'url', '', 1660620367, 1662551337);
INSERT INTO `la_system_config`
VALUES (61, 'oa_channel', 'token', '', 1660620367, 1662551337);
INSERT INTO `la_system_config`
VALUES (62, 'oa_channel', 'encodingAesKey', '', 1660620367, 1662551337);
INSERT INTO `la_system_config`
VALUES (63, 'oa_channel', 'encryptionType', '1', 1660620367, 1662551337);
INSERT INTO `la_system_config`
VALUES (64, 'oa_channel', 'menus', '[]', 1631255140, 1663118712);
INSERT INTO `la_system_config`
VALUES (70, 'login', 'loginWay', '1,2', 1660620367, 1662538771);
INSERT INTO `la_system_config`
VALUES (71, 'login', 'forceBindMobile', '0', 1660620367, 1662538771);
INSERT INTO `la_system_config`
VALUES (72, 'login', 'openAgreement', '1', 1660620367, 1662538771);
INSERT INTO `la_system_config`
VALUES (73, 'login', 'openOtherAuth', '1', 1660620367, 1662538771);
INSERT INTO `la_system_config`
VALUES (74, 'login', 'autoLoginAuth', '1,2', 1660620367, 1662538771);
INSERT INTO `la_system_config`
VALUES (80, 'user', 'defaultAvatar', '/api/static/default_avatar.png', 1660620367, 1662535156);
COMMIT;

BEGIN;
INSERT INTO `la_notice_setting`
VALUES (1, 101, '登录验证码', '用户手机号码登录时发送', 1, 2, '{}',
        '{\"type\":\"sms\",\"templateId\":\"SMS_222458159\",\"content\":\"您正在登录，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"tips\":[\"可选变量 验证码:code\",\"示例：您正在登录，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"生效条件：1、管理后台完成短信设置。2、第三方短信平台申请模板。\"],\"status\":\"1\"}',
        '{}', '{}', 0, 1648696695, 1648696695, 0);
INSERT INTO `la_notice_setting`
VALUES (2, 102, '绑定手机验证码', '用户绑定手机号码时发送', 1, 2, '{}',
        '{\"type\":\"sms\",\"templateId\":\"SMS_175615069\",\"content\":\"您正在绑定手机号，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"tips\":[\"可选变量 验证码:code\",\"示例：您正在绑定手机号，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"生效条件：1、管理后台完成短信设置。2、第三方短信平台申请模板。\"],\"status\":\"1\"}',
        '{}', '{}', 0, 1648696695, 1648696695, 0);
INSERT INTO `la_notice_setting`
VALUES (3, 103, '变更手机验证码', '用户变更手机号码时发送', 1, 2, '{}',
        '{\"type\":\"sms\",\"templateId\":\"SMS_207952628\",\"content\":\"您正在变更手机号，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"tips\":[\"可选变量 验证码:code\",\"示例：您正在变更手机号，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"生效条件：1、管理后台完成短信设置。2、第三方短信平台申请模板。\"],\"status\":\"1\"}',
        '{}', '{}', 0, 1648696695, 1648696695, 0);
INSERT INTO `la_notice_setting`
VALUES (4, 104, '找回登录密码验证码', '用户找回登录密码号码时发送', 1, 2, '{}',
        '{\"type\":\"sms\",\"templateId\":\"SMS_175615069\",\"content\":\"您正在找回登录密码，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"tips\":[\"可选变量 验证码:code\",\"示例：您正在找回登录密码，验证码${code}，切勿将验证码泄露于他人，本条验证码有效期5分钟。\",\"条验证码有效期5分钟。\"],\"status\":\"1\"}',
        '{}', '{}', 0, 1648696695, 1648696695, 0);
COMMIT;

BEGIN;
INSERT INTO `la_decorate_tabbar`
VALUES (13, '首页', '/api/static/tabbar_home_sel.png', '/api/static/tabbar_home.png',
        '{\"path\":\"/pages/index/index\",\"name\":\"商城首页\",\"type\":\"shop\"}', 1662688157, 1662688157);
INSERT INTO `la_decorate_tabbar`
VALUES (14, '资讯', '/api/static/tabbar_text_sel.png', '/api/static/tabbar_text.png',
        '{\"path\":\"/pages/news/news\",\"name\":\"文章资讯\",\"type\":\"shop\"}', 1662688157, 1662688157);
INSERT INTO `la_decorate_tabbar`
VALUES (15, '我的', '/api/static/tabbar_me_sel.png', '/api/static/tabbar_me.png',
        '{\"path\":\"/pages/user/user\",\"name\":\"个人中心\",\"type\":\"shop\"}', 1662688157, 1662688157);
COMMIT;

BEGIN;
INSERT INTO `la_decorate_page`
VALUES (1, 1, '商城首页',
        '[{\"title\":\"搜索\",\"name\":\"search\",\"disabled\":1,\"content\":{},\"styles\":{}},{\"title\":\"首页轮播图\",\"name\":\"banner\",\"content\":{\"enabled\":1,\"data\":[{\"image\":\"/api/static/banner01.png\",\"name\":\"\",\"link\":{\"path\":\"/pages/index/index\",\"name\":\"商城首页\",\"type\":\"shop\"}},{\"image\":\"/api/static/banner02.png\",\"name\":\"\",\"link\":{}}]},\"styles\":{}},{\"title\":\"导航菜单\",\"name\":\"nav\",\"content\":{\"enabled\":1,\"data\":[{\"image\":\"/api/static/nav01.png\",\"name\":\"资讯中心\",\"link\":{\"path\":\"/pages/news/news\",\"name\":\"文章资讯\",\"type\":\"shop\"}},{\"image\":\"/api/static/nav02.png\",\"name\":\"我的收藏\",\"link\":{\"path\":\"/pages/collection/collection\",\"name\":\"我的收藏\",\"type\":\"shop\"}},{\"image\":\"/api/static/nav03.png\",\"name\":\"个人设置\",\"link\":{\"path\":\"/pages/user_set/user_set\",\"name\":\"个人设置\",\"type\":\"shop\"}},{\"image\":\"/api/static/nav04.png\",\"name\":\"联系客服\",\"link\":{\"path\":\"/pages/customer_service/customer_service\",\"name\":\"联系客服\",\"type\":\"shop\"}},{\"image\":\"/api/static/nav05.png\",\"name\":\"关于我们\",\"link\":{\"path\":\"/pages/as_us/as_us\",\"name\":\"关于我们\",\"type\":\"shop\"}}]},\"styles\":{}},{\"id\":\"l84almsk2uhyf\",\"title\":\"资讯\",\"name\":\"news\",\"disabled\":1,\"content\":{},\"styles\":{}}]',
        1661757188, 1663321380);
INSERT INTO `la_decorate_page`
VALUES (2, 2, '个人中心',
        '[{\"title\":\"用户信息\",\"name\":\"user-info\",\"disabled\":1,\"content\":{},\"styles\":{}},{\"title\":\"我的服务\",\"name\":\"my-service\",\"content\":{\"style\":2,\"title\":\"服务中心\",\"data\":[{\"image\":\"/api/static/user_collect.png\",\"name\":\"我的收藏\",\"link\":{\"path\":\"/pages/collection/collection\",\"name\":\"我的收藏\",\"type\":\"shop\"}},{\"image\":\"/api/static/user_setting.png\",\"name\":\"个人设置\",\"link\":{\"path\":\"/pages/user_set/user_set\",\"name\":\"个人设置\",\"type\":\"shop\"}},{\"image\":\"/api/static/user_kefu.png\",\"name\":\"联系客服\",\"link\":{\"path\":\"/pages/customer_service/customer_service\",\"name\":\"联系客服\",\"type\":\"shop\"}}]},\"styles\":{}},{\"title\":\"个人中心广告图\",\"name\":\"user-banner\",\"content\":{\"enabled\":1,\"data\":[{\"image\":\"/api/static/ad01.jpg\",\"name\":\"\",\"link\":{}}]},\"styles\":{}}]',
        1661757188, 1663320728);
INSERT INTO `la_decorate_page`
VALUES (3, 3, '客服设置',
        '[{\"title\":\"客服设置\",\"name\":\"customer-service\",\"content\":{\"title\":\"添加客服二维码\",\"time\":\"早上 9:00 - 22:00\",\"mobile\":\"13800138000\",\"qrcode\":\"\"},\"styles\":{}}]',
        1661757188, 1662689155);
COMMIT;

BEGIN;
INSERT INTO `la_system_auth_menu`
VALUES (1, 0, 'C', '工作台', 'el-icon-Monitor', 999, 'index:console', 'workbench', 'workbench/index', '', '', 1, 1, 0,
        1650341765, 1668672757);
INSERT INTO `la_system_auth_menu`
VALUES (100, 0, 'M', '权限管理', 'el-icon-Lock', 600, '', 'permission', '', '', '', 0, 1, 0, 1650341765, 1662626201);
INSERT INTO `la_system_auth_menu`
VALUES (101, 100, 'C', '管理员', 'local-icon-wode', 0, 'system:admin:list', 'admin', 'permission/admin/index', '', '',
        1, 1, 0, 1650341765, 1663301404);
INSERT INTO `la_system_auth_menu`
VALUES (102, 101, 'A', '管理员详情', '', 0, 'system:admin:detail', '', '', '', '', 0, 1, 0, 1650341765, 1660201785);
INSERT INTO `la_system_auth_menu`
VALUES (103, 101, 'A', '管理员新增', '', 0, 'system:admin:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (104, 101, 'A', '管理员编辑', '', 0, 'system:admin:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (105, 101, 'A', '管理员删除', '', 0, 'system:admin:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (106, 101, 'A', '管理员状态', '', 0, 'system:admin:disable', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (110, 100, 'C', '角色管理', 'el-icon-Female', 0, 'system:role:list', 'role', 'permission/role/index', '', '', 1,
        1, 0, 1650341765, 1663301451);
INSERT INTO `la_system_auth_menu`
VALUES (111, 110, 'A', '角色详情', '', 0, 'system:role:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (112, 110, 'A', '角色新增', '', 0, 'system:role:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (113, 110, 'A', '角色编辑', '', 0, 'system:role:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (114, 110, 'A', '角色删除', '', 0, 'system:role:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (120, 100, 'C', '菜单管理', 'el-icon-Operation', 0, 'system:menu:list', 'menu', 'permission/menu/index', '', '',
        1, 1, 0, 1650341765, 1663301388);
INSERT INTO `la_system_auth_menu`
VALUES (121, 120, 'A', '菜单详情', '', 0, 'system:menu:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (122, 120, 'A', '菜单新增', '', 0, 'system:menu:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (123, 120, 'A', '菜单编辑', '', 0, 'system:menu:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (124, 120, 'A', '菜单删除', '', 0, 'system:menu:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (130, 0, 'M', '组织管理', 'el-icon-OfficeBuilding', 700, '', 'organization', '', '', '', 0, 1, 0, 1650341765,
        1664416715);
INSERT INTO `la_system_auth_menu`
VALUES (131, 130, 'C', '部门管理', 'el-icon-Coordinate', 0, 'system:dept:list', 'department',
        'organization/department/index', '', '', 1, 1, 0, 1650341765, 1660201994);
INSERT INTO `la_system_auth_menu`
VALUES (132, 131, 'A', '部门详情', '', 0, 'system:dept:detail', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (133, 131, 'A', '部门新增', '', 0, 'system:dept:add', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (134, 131, 'A', '部门编辑', '', 0, 'system:dept:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (135, 131, 'A', '部门删除', '', 0, 'system:dept:del', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (200, 0, 'M', '其它管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1660636870);
INSERT INTO `la_system_auth_menu`
VALUES (201, 200, 'M', '图库管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (202, 201, 'A', '文件列表', '', 0, 'albums:albumList', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (203, 201, 'A', '文件命名', '', 0, 'albums:albumRename', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (204, 201, 'A', '文件移动', '', 0, 'albums:albumMove', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (205, 201, 'A', '文件删除', '', 0, 'albums:albumDel', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (206, 201, 'A', '分类列表', '', 0, 'albums:cateList', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (207, 201, 'A', '分类新增', '', 0, 'albums:cateAdd', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (208, 201, 'A', '分类命名', '', 0, 'albums:cateRename', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (209, 201, 'A', '分类删除', '', 0, 'albums:cateDel', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (215, 200, 'M', '上传管理', '', 0, '', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (216, 215, 'A', '上传图片', '', 0, 'upload:image', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (217, 215, 'A', '上传视频', '', 0, 'upload:video', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (500, 0, 'M', '系统设置', 'el-icon-Setting', 0, '', 'setting', '', '', '', 0, 1, 0, 1650341765, 1662626322);
INSERT INTO `la_system_auth_menu`
VALUES (501, 500, 'M', '网站设置', 'el-icon-Basketball', 10, '', 'website', '', '', '', 0, 1, 0, 1650341765,
        1663233572);
INSERT INTO `la_system_auth_menu`
VALUES (502, 501, 'C', '网站信息', '', 0, 'setting:website:detail', 'information', 'setting/website/information', '',
        '', 0, 1, 0, 1650341765, 1660202218);
INSERT INTO `la_system_auth_menu`
VALUES (503, 502, 'A', '保存配置', '', 0, 'setting:website:save', '', '', '', '', 0, 0, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (510, 501, 'C', '政策协议', '', 0, 'setting:protocol:detail', 'protocol', 'setting/website/protocol', '', '', 0,
        1, 0, 1660027606, 1660202312);
INSERT INTO `la_system_auth_menu`
VALUES (511, 510, 'A', '协议保存', '', 0, 'setting:protocol:save', '', '', '', '', 0, 0, 0, 1660027606, 1663670865);
INSERT INTO `la_system_auth_menu`
VALUES (515, 600, 'C', '字典管理', 'el-icon-Box', 0, 'setting:dict:type:list', 'dict', 'setting/dict/type/index', '',
        '', 0, 1, 0, 1660035436, 1663226087);
INSERT INTO `la_system_auth_menu`
VALUES (516, 515, 'A', '字典类型新增', '', 0, 'setting:dict:type:add', '', '', '', '', 0, 1, 0, 1660202761, 1660202761);
INSERT INTO `la_system_auth_menu`
VALUES (517, 515, 'A', '字典类型编辑', '', 0, 'setting:dict:type:edit', '', '', '', '', 0, 1, 0, 1660202842,
        1660202842);
INSERT INTO `la_system_auth_menu`
VALUES (518, 515, 'A', '字典类型删除', '', 0, 'setting:dict:type:del', '', '', '', '', 0, 1, 0, 1660202903, 1660202903);
INSERT INTO `la_system_auth_menu`
VALUES (519, 600, 'C', '字典数据管理', '', 0, 'setting:dict:data:list', 'dict/data', 'setting/dict/data/index',
        '/dev_tools/dict', '', 0, 0, 0, 1660202948, 1663309252);
INSERT INTO `la_system_auth_menu`
VALUES (520, 515, 'A', '字典数据新增', '', 0, 'setting:dict:data:add', '', '', '', '', 0, 1, 0, 1660203117, 1660203117);
INSERT INTO `la_system_auth_menu`
VALUES (521, 515, 'A', '字典数据编辑', '', 0, 'setting:dict:data:edit', '', '', '', '', 0, 1, 0, 1660203142,
        1660203142);
INSERT INTO `la_system_auth_menu`
VALUES (522, 515, 'A', '字典数据删除', '', 0, 'setting:dict:data:del', '', '', '', '', 0, 1, 0, 1660203159, 1660203159);
INSERT INTO `la_system_auth_menu`
VALUES (550, 500, 'M', '系统维护', 'el-icon-SetUp', 0, '', 'system', '', '', '', 0, 1, 0, 1650341765, 1660202466);
INSERT INTO `la_system_auth_menu`
VALUES (551, 550, 'C', '系统环境', '', 0, 'monitor:server', 'environment', 'setting/system/environment', '', '', 0, 1,
        0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (552, 550, 'C', '系统缓存', '', 0, 'monitor:cache', 'cache', 'setting/system/cache', '', '', 0, 1, 0, 1650341765,
        1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (553, 550, 'C', '系统日志', '', 0, 'system:log:operate', 'journal', 'setting/system/journal', '', '', 0, 1, 0,
        1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (555, 500, 'C', '存储设置', 'el-icon-FolderOpened', 6, 'setting:storage:list', 'storage',
        'setting/storage/index', '', '', 0, 1, 0, 1650341765, 1663312996);
INSERT INTO `la_system_auth_menu`
VALUES (556, 555, 'A', '保存配置', '', 0, 'setting:storage:edit', '', '', '', '', 0, 1, 0, 1650341765, 1650341765);
INSERT INTO `la_system_auth_menu`
VALUES (600, 0, 'M', '开发工具', 'el-icon-EditPen', 0, '', 'dev_tools', '', '', '', 0, 1, 0, 1660027606, 1664335701);
INSERT INTO `la_system_auth_menu`
VALUES (700, 0, 'M', '素材管理', 'el-icon-Picture', 43, '', 'material', '', '', '', 0, 1, 0, 1660203293, 1663300847);
INSERT INTO `la_system_auth_menu`
VALUES (701, 700, 'C', '素材中心', 'el-icon-PictureRounded', 0, '', 'index', 'material/index', '', '', 0, 1, 0,
        1660203402, 1663301493);
INSERT INTO `la_system_auth_menu`
VALUES (703, 0, 'M', '文章资讯', 'el-icon-ChatLineSquare', 900, '', 'article', '', '', '', 0, 1, 0, 1661757636,
        1664416659);
INSERT INTO `la_system_auth_menu`
VALUES (704, 703, 'C', '文章管理', 'el-icon-ChatDotSquare', 3, 'article:list', 'lists', 'article/lists/index', '', '',
        1, 1, 0, 1661757743, 1663658220);
-- INSERT INTO `la_system_auth_menu` VALUES (705, 703, 'C', '文章栏目', 'el-icon-CollectionTag', 0, 'article:cate:list', 'column', 'article/column/index', '', '', 1, 1, 0, 1661759218, 1663578137);
-- INSERT INTO `la_system_auth_menu` VALUES (706, 0, 'M', '渠道设置', 'el-icon-Message', 46, '', 'channel', '', '', '', 0, 1, 0, 1661767630, 1664416682);
-- INSERT INTO `la_system_auth_menu` VALUES (707, 706, 'C', 'H5设置', 'el-icon-Cellphone', 0, 'channel:h5:detail', 'h5', 'channel/h5', '', '', 0, 1, 0, 1661768566, 1662626123);
-- INSERT INTO `la_system_auth_menu` VALUES (708, 706, 'M', '微信公众号', 'local-icon-dingdan', 0, '', 'wx_oa', '', '', '', 0, 1, 0, 1661769386, 1663301237);
-- INSERT INTO `la_system_auth_menu` VALUES (709, 708, 'C', '公众号配置', '', 0, 'channel:oa:detail', 'config', 'channel/wx_oa/config', '', '', 0, 1, 0, 1661769457, 1662638440);
-- INSERT INTO `la_system_auth_menu` VALUES (710, 706, 'C', '微信小程序', 'local-icon-weixin', 0, 'channel:mp:detail', 'weapp', 'channel/weapp', '', '', 0, 1, 0, 1661823746, 1663301266);
-- INSERT INTO `la_system_auth_menu` VALUES (711, 706, 'C', '微信开发平台', 'el-icon-DataBoard', 0, 'channel:wx:detail', 'wx_dev', 'channel/wx_dev', '', '', 0, 0, 0, 1661824989, 1663310675);
INSERT INTO `la_system_auth_menu`
VALUES (712, 0, 'M', '用户管理', 'el-icon-User', 800, '', 'consumer', '', '', '', 0, 1, 0, 1661832966, 1663294141);
INSERT INTO `la_system_auth_menu`
VALUES (713, 712, 'C', '用户列表', 'el-icon-User', 0, 'user:list', 'lists', 'consumer/lists/index', '', '', 0, 1, 0,
        1661839365, 1663301092);
INSERT INTO `la_system_auth_menu`
VALUES (714, 714, 'A', '用户编辑', '', 0, 'user:edit', 'detail', 'consumer/lists/detail', '/consumer/lists', '', 0, 0,
        0, 1661840502, 1662627718);
-- INSERT INTO `la_system_auth_menu` VALUES (715, 600, 'C', '编辑数据表', '', 0, 'gen:editTable', 'code/edit', 'dev_tools/code/edit', '/dev_tools/code', '', 0, 0, 0, 1661843525, 1661843615);
-- INSERT INTO `la_system_auth_menu` VALUES (716, 705, 'A', '栏目详情', '', 0, 'article:cate:detail', 'lists/edit', 'article/lists/edit', '/article/lists', '', 0, 0, 0, 1661844126, 1662626009);
-- INSERT INTO `la_system_auth_menu` VALUES (717, 0, 'M', '装修管理', 'el-icon-Brush', 47, '', 'decoration', '', '', '', 0, 1, 0, 1661845634, 1664416675);
-- INSERT INTO `la_system_auth_menu` VALUES (718, 717, 'C', '页面装修', 'el-icon-CopyDocument', 0, 'decorate:pages:detail', 'pages', 'decoration/pages/index', '', '', 0, 1, 0, 1661845678, 1663294313);
-- INSERT INTO `la_system_auth_menu` VALUES (719, 717, 'C', '底部导航', 'el-icon-Position', 0, 'decorate:tabbar:detail', 'tabbar', 'decoration/tabbar', '', '', 0, 1, 0, 1661845811, 1663294354);
-- INSERT INTO `la_system_auth_menu` VALUES (720, 500, 'M', '消息通知', 'el-icon-Message', 9, '', 'message', '', '', '', 0, 1, 0, 1661848742, 1662626364);
-- INSERT INTO `la_system_auth_menu` VALUES (721, 720, 'C', '通知设置', '', 0, 'setting:notice:list', 'notice', 'message/notice/index', '', '', 0, 1, 0, 1661848772, 1662638112);
-- INSERT INTO `la_system_auth_menu` VALUES (722, 720, 'C', '通知详情', '', 0, 'setting:notice:detail', 'notice/edit', 'message/notice/edit', '/setting/message/notice', '', 0, 0, 0, 1661848944, 1663142853);
-- INSERT INTO `la_system_auth_menu` VALUES (723, 720, 'C', '短信设置', '', 0, 'setting:sms:list', 'short_letter', 'message/short_letter/index', '', '', 0, 1, 0, 1661848995, 1662638165);
-- INSERT INTO `la_system_auth_menu` VALUES (724, 500, 'M', '用户设置', 'local-icon-keziyuyue', 8, '', 'user', '', '', '', 0, 1, 0, 1662455407, 1663301570);
-- INSERT INTO `la_system_auth_menu` VALUES (725, 724, 'C', '用户设置', '', 0, 'setting:user:detail', 'setup', 'setting/user/setup', '', '', 0, 1, 0, 1662455555, 1663312225);
-- INSERT INTO `la_system_auth_menu` VALUES (726, 724, 'C', '登录注册', '', 0, 'setting:login:detail', 'login_register', 'setting/user/login_register', '', '', 0, 1, 0, 1662456475, 1663312263);
-- INSERT INTO `la_system_auth_menu` VALUES (728, 500, 'C', '热门搜索', 'el-icon-Search', 7, 'setting:search:detail', 'search', 'setting/search/index', '', '', 0, 1, 0, 1662540429, 1663312392);
-- INSERT INTO `la_system_auth_menu` VALUES (730, 704, 'A', '文章新增', '', 0, 'article:add', '', '', '', '', 0, 1, 0, 1662625870, 1662625870);
-- INSERT INTO `la_system_auth_menu` VALUES (732, 704, 'A', '文章删除', '', 0, 'article:del', '', '', '', '', 0, 1, 0, 1662625894, 1662625894);
-- INSERT INTO `la_system_auth_menu` VALUES (733, 704, 'A', '文章状态', '', 0, 'article:change', '', '', '', '', 0, 1, 0, 1662625909, 1662625909);
-- INSERT INTO `la_system_auth_menu` VALUES (734, 705, 'A', '栏目新增', '', 0, 'article:cate:add', '', '', '', '', 0, 1, 0, 1662626024, 1662626024);
-- INSERT INTO `la_system_auth_menu` VALUES (735, 705, 'A', '栏目编辑', '', 0, 'article:cate:edit', '', '', '', '', 0, 1, 0, 1662626044, 1662626044);
-- INSERT INTO `la_system_auth_menu` VALUES (736, 705, 'A', '栏目删除', '', 0, 'article:cate:del', '', '', '', '', 0, 1, 0, 1662626060, 1662626060);
-- INSERT INTO `la_system_auth_menu` VALUES (737, 705, 'A', '栏目状态', '', 0, 'article:cate:change', '', '', '', '', 0, 1, 0, 1662626077, 1662626077);
INSERT INTO `la_system_auth_menu`
VALUES (738, 703, 'A', '文章编辑', '', 0, 'article:edit', 'lists/edit', 'article/lists/edit', '', '', 0, 0, 0,
        1662626554, 1663309550);
INSERT INTO `la_system_auth_menu`
VALUES (739, 712, 'C', '用户详情', '', 0, 'user:detail', 'detail', 'consumer/lists/detail', '/consumer/lists', '', 0, 0,
        0, 1662628049, 1662628049);
INSERT INTO `la_system_auth_menu`
VALUES (740, 739, 'A', '用户编辑', '', 0, 'user:edit', '', '', '', '', 0, 1, 0, 1662628085, 1662628085);
-- INSERT INTO `la_system_auth_menu` VALUES (741, 721, 'A', '设置保存', '', 0, 'setting:notice:save', '', '', '', '', 0, 1, 0, 1662638049, 1662638049);
-- INSERT INTO `la_system_auth_menu` VALUES (742, 723, 'A', '短信详情', '', 0, 'setting:sms:detail', '', '', '', '', 0, 1, 0, 1662638180, 1662638180);
-- INSERT INTO `la_system_auth_menu` VALUES (743, 723, 'A', '保存设置', '', 0, 'setting:sms:save', '', '', '', '', 0, 1, 0, 1662638196, 1662638196);
-- INSERT INTO `la_system_auth_menu` VALUES (744, 707, 'A', '设置保存', '', 0, 'channel:h5:save', '', '', '', '', 0, 1, 0, 1662638326, 1662638326);
-- INSERT INTO `la_system_auth_menu` VALUES (745, 710, 'A', '设置保存', '', 0, 'channel:mp:detail', '', '', '', '', 0, 1, 0, 1662638359, 1662638359);
-- INSERT INTO `la_system_auth_menu` VALUES (746, 711, 'A', '保存设置', '', 0, 'channel:wx:save', '', '', '', '', 0, 1, 0, 1662638410, 1662638410);
-- INSERT INTO `la_system_auth_menu` VALUES (747, 709, 'A', '保存', '', 0, 'channel:oa:save', '', '', '', '', 0, 1, 0, 1662638459, 1663310514);
-- INSERT INTO `la_system_auth_menu` VALUES (748, 708, 'C', '菜单管理', '', 0, '', 'menu', 'channel/wx_oa/menu', '', '', 0, 1, 0, 1663050714, 1663050714);
-- INSERT INTO `la_system_auth_menu` VALUES (750, 708, 'C', '关注回复', '', 0, 'channel:oaReplyFollow:list', 'follow', 'channel/wx_oa/reply/follow_reply', '', '', 0, 1, 0, 1663149592, 1664511108);
-- INSERT INTO `la_system_auth_menu` VALUES (751, 708, 'C', '关键字回复', '', 0, 'channel:oaReplyKeyword:list', 'keyword', 'channel/wx_oa/reply/keyword_reply', '', '', 0, 1, 0, 1663149622, 1664511241);
-- INSERT INTO `la_system_auth_menu` VALUES (752, 708, 'C', '默认回复', '', 0, 'channel:oaReplyDefault:list', 'default', 'channel/wx_oa/reply/default_reply', '', '', 0, 1, 0, 1663149650, 1664517685);
-- INSERT INTO `la_system_auth_menu` VALUES (753, 718, 'A', '保存', '', 0, 'decorate:pages:save', '', '', '', '', 0, 1, 0, 1663236648, 1663236648);
-- INSERT INTO `la_system_auth_menu` VALUES (754, 719, 'A', '保存', '', 0, 'decorate:tabbar:save', '', '', '', '', 0, 1, 0, 1663236675, 1663236675);
-- INSERT INTO `la_system_auth_menu` VALUES (755, 704, 'A', '文章详情', '', 0, 'article:detail', '', '', '', '', 0, 1, 0, 1663310241, 1663310252);
-- INSERT INTO `la_system_auth_menu` VALUES (756, 748, 'A', '发布', '', 0, 'channel:oaMenu:publish', '', '', '', '', 0, 1, 0, 1663310379, 1663310525);
-- INSERT INTO `la_system_auth_menu` VALUES (757, 748, 'A', '保存', '', 0, 'channel:oaMenu:save', '', '', '', '', 0, 1, 0, 1663310556, 1663310556);
-- INSERT INTO `la_system_auth_menu` VALUES (758, 725, 'A', '保存', '', 0, 'setting:user:save', '', '', '', '', 0, 1, 0, 1663312193, 1663312193);
-- INSERT INTO `la_system_auth_menu` VALUES (759, 726, 'A', '保存', '', 0, 'setting:login:save', '', '', '', '', 0, 1, 0, 1663312289, 1663312289);
-- INSERT INTO `la_system_auth_menu` VALUES (760, 728, 'A', '保存', '', 0, 'setting:search:save', '', '', '', '', 0, 1, 0, 1663312423, 1663312423);
-- INSERT INTO `la_system_auth_menu` VALUES (762, 750, 'A', '新增', '', 0, 'channel:oaReplyFollow:add', '', '', '', '', 1, 1, 0, 1664511131, 1664511131);
-- INSERT INTO `la_system_auth_menu` VALUES (763, 750, 'A', '状态', '', 0, 'channel:oaReplyFollow:status', '', '', '', '', 1, 1, 0, 1664511160, 1664511160);
-- INSERT INTO `la_system_auth_menu` VALUES (764, 750, 'A', '编辑', '', 0, 'channel:oaReplyFollow:edit', '', '', '', '', 1, 1, 0, 1664511177, 1664511190);
-- INSERT INTO `la_system_auth_menu` VALUES (765, 750, 'A', '删除', '', 0, 'channel:oaReplyFollow:del', '', '', '', '', 1, 1, 0, 1664511208, 1664511208);
-- INSERT INTO `la_system_auth_menu` VALUES (766, 751, 'A', '新增', '', 0, 'channel:oaReplyKeyword:add', '', '', '', '', 1, 1, 0, 1664511264, 1664511264);
-- INSERT INTO `la_system_auth_menu` VALUES (767, 751, 'A', '状态', '', 0, 'channel:oaReplyKeyword:status', '', '', '', '', 1, 1, 0, 1664511295, 1664511295);
-- INSERT INTO `la_system_auth_menu` VALUES (768, 751, 'A', '编辑', '', 0, 'channel:oaReplyKeyword:edit', '', '', '', '', 1, 1, 0, 1664511312, 1664511312);
-- INSERT INTO `la_system_auth_menu` VALUES (769, 751, 'A', '删除', '', 0, 'channel:oaReplyKeyword:del', '', '', '', '', 1, 1, 0, 1664511327, 1664511327);
-- INSERT INTO `la_system_auth_menu` VALUES (770, 752, 'A', '新增', '', 0, 'channel:oaReplyDefault:add', '', '', '', '', 1, 1, 0, 1664517709, 1664517709);
-- INSERT INTO `la_system_auth_menu` VALUES (771, 752, 'A', '编辑', '', 0, 'channel:oaReplyDefault:edit', '', '', '', '', 1, 1, 0, 1664517725, 1664517725);
-- INSERT INTO `la_system_auth_menu` VALUES (772, 752, 'A', '状态', '', 0, 'channel:oaReplyDefault:status', '', '', '', '', 1, 1, 0, 1664517757, 1664517757);
-- INSERT INTO `la_system_auth_menu` VALUES (773, 752, 'A', '删除', '', 0, 'channel:oaReplyDefault:del', '', '', '', '', 1, 1, 0, 1664517778, 1664517778);
-- INSERT INTO `la_system_auth_menu` VALUES (774, 610, 'A', '导入数据表列表', '', 0, 'gen:db', '', '', '', '', 1, 1, 0, 1665646316, 1665646316);
INSERT INTO `la_system_auth_menu`
VALUES (775, 703, 'C', '文章添加/编辑', '', 0, 'article:add/edit', 'lists/edit', 'article/lists/edit', '/article/lists',
        '', 0, 0, 0, 1668677477, 1668677477);
INSERT INTO `la_system_auth_menu` VALUES (776, 0, 'M', '资料管理', 'el-icon-Grid', 950, '', 'information', 'information', '', '', 1, 1, 0, 1711439525, 1711439525);
INSERT INTO `la_system_auth_menu` VALUES (777, 776, 'M', '武器管理', '', 3, '', 'weapons', 'information/weapons/index', '', '', 1, 1, 0, 1711417265, 1711417265);
INSERT INTO `la_system_auth_menu` VALUES (778, 777, 'A', '武器编辑', '', 0, 'information:weapons:edit', '', '', '', '', 1, 1, 0, 1711417662, 1711417662);
INSERT INTO `la_system_auth_menu` VALUES (779, 777, 'C', '武器添加/编辑', '', 0, 'weapons:add/edit', 'edit', 'information/weapons/edit', '', '', 0, 0, 0, 1711418200, 1711418200);
INSERT INTO `la_system_auth_menu` VALUES (780, 777, 'C', '武器列表', '', 3, 'information:weapons:list', 'weapons', 'information/weapons/index', '', '', 1, 1, 0, 1711433509, 1711433509);
INSERT INTO `la_system_auth_menu` VALUES (781, 777, 'C', '系列管理', '', 0, 'information:weapons:series:list', 'series/index', 'information/weapons/series/index', '', '', 1, 1, 0, 1711693863, 1711693863);
INSERT INTO `la_system_auth_menu` VALUES (782, 776, 'M', '英雄管理', '', 0, '', 'hero', 'information/hero', '', '', 1, 1, 0, 1713433403, 1713433403);
INSERT INTO `la_system_auth_menu` VALUES (783, 782, 'C', '英雄定位', '', 0, 'information:hero:type:list', 'type/index', 'information/hero/type/index', '', '', 1, 1, 0, 1713433480, 1713433480);
INSERT INTO `la_system_auth_menu` VALUES (784, 782, 'A', '英雄编辑', '', 0, 'information:hero:edit', '', '', '', '', 1, 1, 0, 1713433510, 1713433510);
INSERT INTO `la_system_auth_menu` VALUES (785, 782, 'C', '英雄添加编辑', '', 0, 'hero:add/edit', 'edit', 'information/hero/edit', '/information/hero/index', '', 0, 0, 0, 1713433605, 1713433605);
INSERT INTO `la_system_auth_menu` VALUES (786, 782, 'C', '英雄列表', '', 3, 'information:hero:list', 'index', 'information/hero/index', '', '', 1, 1, 0, 1713433699, 1713433699);
INSERT INTO `la_system_auth_menu` VALUES (787, 776, 'M', '地图管理', '', 0, '', 'maps', 'information/maps/index', '', '', 1, 1, 0, 1713514039, 1713514039);
INSERT INTO `la_system_auth_menu` VALUES (788, 787, 'C', '地图列表', '', 0, 'information:maps:list', 'index', 'information/maps/index', '', '', 1, 1, 0, 1713514205, 1713514205);
INSERT INTO `la_system_auth_menu` VALUES (789, 787, 'A', '地图编辑', '', 0, 'informatiion:maps:edit', '', '', '', '', 1, 1, 0, 1713514260, 1713514260);
INSERT INTO `la_system_auth_menu` VALUES (790, 787, 'C', '地图添加/编辑', '', 0, 'maps:add/edit', 'edit', 'information/maps/edit', 'information/maps/index', '', 0, 0, 0, 1713514373, 1713514373);
COMMIT;

