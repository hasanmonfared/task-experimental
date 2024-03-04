-- +migrate Up
INSERT INTO `users` (`id`, `name`,`phone_number`) VALUES(1, 'hassan','09377561162');

INSERT INTO `vendors` (`id`, `name`) VALUES(1, 'chelopaz');
INSERT INTO `vendors` (`id`, `name`) VALUES(2, 'kababi');
INSERT INTO `vendors` (`id`, `name`) VALUES(3, 'pitza');


INSERT INTO `agents` (`id`, `firstname`,`lastname`) VALUES(1, 'hossein','doulabi');
INSERT INTO `agents` (`id`, `firstname`,`lastname`) VALUES(2, 'hassan','doulabi');

INSERT INTO `orders` (`id`, `user_id`,`vendor_id`,`delivery_time`,`status`) VALUES(1, 1,1,2,'ready_to_send');
INSERT INTO `orders` (`id`, `user_id`,`vendor_id`,`delivery_time`,`status`) VALUES(2, 1,2,50,'ready_to_send');


-- +migrate Down
DELETE FROM `users` WHERE id =1;
DELETE FROM `vendors` WHERE id in (1,2);
DELETE FROM `agents` WHERE id =1;
