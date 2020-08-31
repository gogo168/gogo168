<pre>

 
CREATE TABLE `test_table` (
  `field_key` varchar(64) NOT NULL DEFAULT '',
  `field_one` varchar(64) DEFAULT NULL,
  `field_two` tinyint(1) DEFAULT NULL,
  `field_thr` int(12) DEFAULT NULL,
  `field_fou` float DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
 

INSERT INTO `test_table` (`field_key`, `field_one`, `field_two`, `field_thr`, `field_fou`) VALUES
('key001', 'one123', 1, 123455, 123.45),
('key002', 'one456', 0, 5678, 0.01),
('key003', 'one789', 1, 5678, 0.02);
 
ALTER TABLE `test_table`
  ADD PRIMARY KEY (`field_key`) USING BTREE;
COMMIT;



<pre>
