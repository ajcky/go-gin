<?php
namespace Common;



class Database
{
    protected static $db;
    private function __construct()
    {

    }

    public static function getInstance()
    {
        if(self::$db){
            return self::$db;
        }else{
            return self::$db = new self();
        }
    }

    function where($where)
    {
        return $this;
    }

    function order($order)
    {
        return $this;
    }

    function limit($limit)
    {
        return $this;
    }
}