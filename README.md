


--
-- 資料庫： `test`
--

-- --------------------------------------------------------
 

CREATE TABLE `test_table` (
  `field_key` varchar(64) NOT NULL DEFAULT '',
  `field_one` varchar(64) DEFAULT NULL,
  `field_two` tinyint(1) DEFAULT NULL,
  `field_thr` int(12) DEFAULT NULL,
  `field_fou` float DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

--
-- 資料表的匯出資料 `test_table`
--

INSERT INTO `test_table` (`field_key`, `field_one`, `field_two`, `field_thr`, `field_fou`) VALUES
('key001', 'one123', 1, 123455, 123.45),
('key002', 'one456', 0, 5678, 0.01),
('key003', 'one789', 1, 5678, 0.02);

--
-- 已匯出資料表的索引
--

--
-- 資料表索引 `test_table`
--
ALTER TABLE `test_table`
  ADD PRIMARY KEY (`field_key`) USING BTREE;
COMMIT;
