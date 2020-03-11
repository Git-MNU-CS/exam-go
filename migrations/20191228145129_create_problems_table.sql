-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS `problems` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL COMMENT '题目名称',
  `describe` varchar(255) NOT NULL COMMENT '描述',
  `level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1、简单2、中等、3、困难、4非常困难',
  `course_id` int(11) NOT NULL COMMENT '所属课程ID',
  `type` enum('choice','fill-blank','program','short-answer') NOT NULL DEFAULT 'program',
  `status` enum('enable','disable') NOT NULL DEFAULT 'enable',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `problems`;
