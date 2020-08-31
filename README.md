main.go

Remove.  Update Insert. QueryAll. QueryAll. QueryByKey.  

目前 函数只能  支持 一张表 一个结构体    
函数  改成 通用的 函数   可以多结构体对应 多表 都支持  
  
刚学go     还搞不懂    结构体 如何 弄到  函数    跟 函数返回特定结构体

意思就是   有 20张表  我 只要 建立 20个结构体 就可以  用这几个函数进行  修改删除新添加 和查询


以下是  数据库
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
