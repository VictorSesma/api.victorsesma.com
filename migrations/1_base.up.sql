CREATE TABLE `curriculum_vitae`
(
    `id` INT NOT NULL,
    `section_type` VARCHAR (45) NULL,
    `name` VARCHAR (45) NULL,
    `show_order` INT NULL,
    `start_date` DATETIME NULL,
    `end_date` DATETIME NULL,
    `description` VARCHAR (600) NULL,
    `summary` VARCHAR (500) NULL,
    PRIMARY KEY (`id`)
);