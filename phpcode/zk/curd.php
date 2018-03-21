<?php
// 创建一个与服务器的连接
$zookeeper = new Zookeeper('115.159.215.179:2181');
$aclArray = array(
array(
'perms'  => Zookeeper::PERM_ALL,
'scheme' => 'world',
'id'     => 'anyone',
)
);
// 创建一个目录节点
$path = '/path';

if ( ! $zookeeper->exists($path))
$zookeeper->create($path, "parent", $aclArray);

$childPath = '/path/child';

if ( ! $zookeeper->exists($childPath))
$zookeeper->create($childPath,"child",$aclArray);


function callback(){
    echo "callback01";
}

// 判断目录是否存在
if ($zookeeper->exists($childPath)){
    // 取出子目录节点内容
    var_dump($zookeeper->get($childPath,call_user_func("callback")));
}

$zookeeper->set($childPath,'child01');

// 判断目录是否存在
if ($zookeeper->exists($childPath)){
    // 取出子目录节点内容
    var_dump($zookeeper->get($childPath));
}

// 删除目录节点
$zookeeper->delete($childPath);

// 判断目录是否存在
if ($zookeeper->exists($childPath)){
    var_dump("true");
} else {
    var_dump("false");
}
