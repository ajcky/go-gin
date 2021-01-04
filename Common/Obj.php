<?php
/**
 * 命名空间必须和目录一致 Common
 * 类名必须和文件名一致 Obj
**/

namespace Common;

class Obj
{
    protected $array = [];
    public static function index()
    {

    }
    function __set($name, $value)
    {
        $this->array[$name] = $value;
    }

    function __get($name)
    {
        return $this->array[$name];
    }

    function __call($name, $arguments)
    {
        return 'function:'.__METHOD__;
    }

    function __toString()
    {
        return __CLASS__;
    }

    function __invoke($param)
    {
        return 'invoke';
    }

}