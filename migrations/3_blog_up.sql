CREATE TABLE `blog`
(
  `postID` INT NOT NULL,
  `publishedOn` DATE NULL DEFAULT NULL,
  `postTitle` VARCHAR
(150) NULL DEFAULT NULL,
  `postContent` VARCHAR
(700) NULL DEFAULT NULL,
  PRIMARY KEY
(`postId`));