-- +migrate Up
INSERT INTO `users` (`id`, `name`,`phone_number`) VALUES(1, 'hassan','09377561162');

INSERT INTO `vendors` (`id`, `name`) VALUES(1, 'chelopaz');
INSERT INTO `vendors` (`id`, `name`) VALUES(2, 'kababi');


INSERT INTO `agents` (`id`, `firstname`,`lastname`) VALUES(1, 'hossein','doulabi');

INSERT INTO `orders` (`id`, `user_id`,`vendor_id`,`delivery_time`) VALUES(1, 1,1,NOW());


-- +migrate Down
DELETE FROM `users` WHERE id =1;
DELETE FROM `vendors` WHERE id in (1,2);
DELETE FROM `agents` WHERE id =1;
