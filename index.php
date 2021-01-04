<?php
ini_set('display_errors','On');
error_reporting(E_ALL);

define('BASEDIR',__DIR__);
include BASEDIR.'/Common/Loader.php';
spl_autoload_register("\\Common\\Loader::autoload");

$db = new \Common\Database\Mysqli();
$db->connect('rm-uf63zn69702x935h5so.mysql.rds.aliyuncs.com','qingshu_mysql','Qingshu520','liaoke',3306);
$res = $db->query('select * from user limit 1');
var_dump($res);