<?php
namespace Common;

class Factory
{
    public static function createDB()
    {
        $db = Database::getInstance();
        Register::set('db',$db);
        return $db;
    }
}