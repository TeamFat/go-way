<?php

/**
 * PHP Zookeeper
 * YuChao php access zookeeper library
 *
 * @category  Libraries
 * @package   EdjYuChao
 * @author    Yu Chao<yuchao@YuChao-inc.cn>
 * @copyright YuChao GPL
 * @license   http://www.php.net/license The PHP License, version 3.01
 */
class EdjZkWatcher extends Zookeeper
{

    private $path = '/YuChao';

    public function watcher()
    {
        echo "Insider Watcher ";
        
    }

    public function watch($path, $callback = 'default')
    {
        // path exist
        if ($callback == 'default') {
            $callback = array($this, 'watcher');
        } else {
            $callback = array($this, 'watch');
        }
        $YuChao = $this->get($path, $callback);
        echo "watch is " . $YuChao . "\n";
        return $YuChao;
    }

}

$zoo = new EdjZkWatcher('115.159.215.179:2181');
//$zoo->set("/YuChao","yuchaodemo");
//$YuChao = $zoo->get( '/YuChao', array($zoo, 'watcher' ) );
$watch_YuChao = $zoo->watch('/foo/001');
$other = $zoo->watch('/foo/002');
$zoo->set('/foo/002','1');

//echo "The YuChao is ".$YuChao;
echo "The YuChao is " . $watch_YuChao;
while (true) {
    echo '.';
    sleep(2);
}

?>