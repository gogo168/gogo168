<?php

include __DIR__.'/MysqlStructSync.php'; 

 
$local_database_config=['host'=>'127.0.0.1','username'=>'root','passwd'=>'root','dbname'=>'test','port'=>3306];

$develop_database_config=['host'=>'127.0.0.1','username'=>'root','passwd'=>'root','dbname'=>'sakila','port'=>3306];

//把local数据库结构更新为develop数据库结构
$compare=new \linge\MysqlStructSync($local_database_config,$develop_database_config);

$compare->removeAutoIncrement();

$compare->baseDiff(); //TABLE COLUMNS(ADD,DROP,MODIFY) CONSTRAINTS(PK,FK,index, ... etc)

$compare->advanceDiff(); //VIEW TRIGGER EVENT FUNCTION PROCEDURE (ADD,DROP)

$diff_sql=$compare->getDiffSql();
//print_r($diff_sql);

/*******************************************/
//用法一:自动执行全部差异语句,更新结构
//$execute_sql_stat=$compare->execute();
//print_r($execute_sql_stat);


//用法二:手动选择要执行的差异语句,记住:选择储存过程，函数等请确保数据库表已经同步
$compare->manuallySelectUpdates();
?>
