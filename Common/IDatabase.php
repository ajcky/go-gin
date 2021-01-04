<?php
namespace Common;
interface IDatabase
{
    function connect($host,$user,$password,$dbName,$port);
    function query($sql);
    function close();
}