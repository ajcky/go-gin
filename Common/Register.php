<?php
namespace Common;

class Register
{
    protected static $object;

    public static function set($alias, $obj)
    {
        self::$object[$alias] = $obj;
    }

    public static function get($name)
    {
        return self::$object[$name];
    }

    public static function _unset($alias)
    {
        unset(self::$object[$alias]);
    }
}