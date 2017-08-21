<?php

class stock{

    public static function httpGet($url) { 
        $curlHandle = curl_init(); 
        curl_setopt( $curlHandle , CURLOPT_URL, $url );
        curl_setopt( $curlHandle , CURLOPT_RETURNTRANSFER, 1 ); 
        curl_setopt( $curlHandle , CURLOPT_SSL_VERIFYPEER, false);
        curl_setopt( $curlHandle , CURLOPT_SSL_VERIFYHOST, false);
        curl_setopt( $curlHandle , CURLOPT_TIMEOUT, 10 ); 
        $content = curl_exec( $curlHandle );
        curl_close( $curlHandle ); 
        $content = preg_replace('/([A-Za-z]+)/', '"${1}"', $content); 
        return json_decode($content, true); 
    }

    public static function httpPost($url, $data){
        $data_string = json_encode($data);                                                                                                                                                           
        $ch = curl_init($url);                                                                      
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, "POST");                                                                     
        curl_setopt($ch, CURLOPT_POSTFIELDS, $data_string);                                                                  
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);                                                                      
        curl_setopt($ch, CURLOPT_HTTPHEADER, array(                                                                          
            'Content-Type: application/json',                                                                                
            'Content-Length: ' . strlen($data_string))                                                                       
        );                                                                                                                                                                                                                      
        $result = curl_exec($ch);
        curl_close( $ch );
        return json_decode($result, true); 
    }

    public static function getStockInfo($stockId, $scale, $dataLen){
        $url = "http://money.finance.sina.com.cn/quotes_service/api/json_v2.php/CN_MarketData.getKLineData?symbol={$stockId}&scale={$scale}&ma=no&datalen={$dataLen}";
        return self::httpGet($url);
    }

    public static function postStockInfo($datas, $stockId){
        $url = "http://localhost:9200/stock/{$stockId}";
        foreach($datas as $data){
            $res = self::httpPost($url, $data);
            var_dump($res['created']);
        }
    }
}
$stockId = 'sh601318';
$stockId = 'sz002236';
$result = stock::getStockInfo($stockId, 5, 10000);
var_dump($result);
stock::postStockInfo($result, $stockId);