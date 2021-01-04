<?php
namespace Common\Database;
use Common\IDatabase;

class Mysqli implements IDatabase
{
    protected $conn;

    function connect($host, $user, $password, $dbName,$port)
    {
        $conn = mysqli_connect($host,$user,$password,$dbName,$port);
        $this->conn = $conn;
    }
    public function close()
    {
        mysqli_close($this->conn);
    }
    function query($sql):array
    {
        $result = mysqli_query($this->conn,$sql,1);

        while($row = mysqli_fetch_row($result)){
            return $row;
        }
        return [];
    }
}
