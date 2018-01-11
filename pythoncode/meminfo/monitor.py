#!/usr/bin/python
# -*- coding: UTF-8 -*-
import time
import MySQLdb as mysql

db = mysql.connect("127.0.0.1","root","devop","memory")
db.autocommit(True)
cur = db.cursor()

def getMem():
    with open('/proc/meminfo') as f:
        total = int(f.readline().split()[1])
        free = int(f.readline().split()[1])
        buffers = int(f.readline().split()[1])
        cache = int(f.readline().split()[1])
    mem_use = total-free-buffers-cache
    t = int(time.time())
    sql = 'insert into memory (memory,time) value (%s,%s)'%(mem_use/1024,t)
    cur.execute(sql)
    print mem_use/1024
    #print 'ok'
while True:
    time.sleep(1)
    getMem()
    #https://github.com/shengxinjing/my_blog/issues/1