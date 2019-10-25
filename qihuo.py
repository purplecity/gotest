# -*- coding:utf-8 -*-
import socket

__author__ = 'weijie'

from EmQuantAPI import *
import traceback
from multiprocessing import Manager, Pool
from readerwriterlock import rwlock
import threading
import queue


NotifyQueue = Manager().Queue()
count = 0
num = 1
HOST = "127.0.0.1"
whiteIP1 = "47.244.212.51"
whiteIP2 = "47.244.217.66"


def mainCallback(quantdata):
    """
    mainCallback 是主回调函数，可捕捉如下错误
    在start函数第三个参数位传入，该函数只有一个为c.EmQuantData类型的参数quantdata
    :param quantdata:c.EmQuantData
    :return:
    """
    print("mainCallback", str(quantdata))
    # 登录掉线或者 登陆数达到上线（即登录被踢下线） 这时所有的服务都会停止
    if str(quantdata.ErrorCode) == "10001011" or str(quantdata.ErrorCode) == "10001009":
        print("Your account is disconnect. You can force login automatically here if you need.")
    # 行情登录验证失败（每次连接行情服务器时需要登录验证）或者行情流量验证失败时，会取消所有订阅，用户需根据具体情况处理
    elif str(quantdata.ErrorCode) == "10001021" or str(quantdata.ErrorCode) == "10001022":
        print("Your all csq subscribe have stopped.")
    # 行情服务器断线自动重连连续6次失败（1分钟左右）不过重连尝试还会继续进行直到成功为止，遇到这种情况需要确认两边的网络状况
    elif str(quantdata.ErrorCode) == "10002009":
        print("Your all csq subscribe have stopped, reconnect 6 times fail.")
        # 行情订阅遇到一些错误(这些错误会导致重连，错误原因通过日志输出，统一转换成EQERR_QUOTE_RECONNECT在这里通知)，正自动重连并重新订阅,可以做个监控
    elif str(quantdata.ErrorCode) == "10002012":
        print("csq subscribe break on some error, reconnect and request automatically.")
        # 资讯服务器断线自动重连连续6次失败（1分钟左右）不过重连尝试还会继续进行直到成功为止，遇到这种情况需要确认两边的网络状况
    elif str(quantdata.ErrorCode) == "10002014":
        print("Your all cnq subscribe have stopped, reconnect 6 times fail.")
    # 资讯订阅遇到一些错误(这些错误会导致重连，错误原因通过日志输出，统一转换成EQERR_INFO_RECONNECT在这里通知)，正自动重连并重新订阅,可以做个监控
    elif str(quantdata.ErrorCode) == "10002013":
        print("cnq subscribe break on some error, reconnect and request automatically.")
    # 资讯登录验证失败（每次连接资讯服务器时需要登录验证）或者资讯流量验证失败时，会取消所有订阅，用户需根据具体情况处理
    elif str(quantdata.ErrorCode) == "10001024" or str(quantdata.ErrorCode) == "10001025":
        print("Your all cnq subscribe have stopped.")
    else:
        pass


def login():
    loginResult = c.start("ForceLogin=1", '', mainCallback)
    if (loginResult.ErrorCode != 0):
        print("login in fail")
        exit()


# IF1911
IF1911Queue = Manager().Queue()
IF1911 = "IF1911.CFE"
IF1911PORT = 50000
IF1911Lock = rwlock.RWLockFair()

def IF1911Get():
    while True:
        p = IF1911Queue.get()


def IF1911Process():
    numQueue = Queue.Queue()
    getThread = threading.Thread(target=IF1911Get())
    getThread.start()
    with socket.socket(socket.AF_INET,socket.SOCK_STREAM) as s:
        s.bind(HOST,IF1911PORT)
        s.listen(2)  #测试下第三个连接来会不会被拒绝
        while True:
            conn,addr = s.accept()
            with conn:
                print("connected by",addr)
                if addr != whiteIP1 and addr != whiteIP2:
                    conn.close()
                else:




def csqIF1911Callback(quantdata):
    """
    csqCallback 是csq订阅时提供的回调函数模板。该函数只有一个为c.EmQuantData类型的参数quantdata
    :param quantdata:c.EmQuantData
    :return:
    """
    IF1911Queue.put(quantdata.Data[IF1911][1])

def subIF1911():
    try:
        # 实时行情订阅使用范例
        data = c.csq(IF1911, 'TIME,Now', 'Pushtype=1', csqIF1911Callback)
        if (data.ErrorCode != 0):
            print("request csq Error, ", data.ErrorMsg)
            NotifyQueue.put(IF1911)

    except Exception as ee:
        print("error >>>", ee)
        traceback.print_exc()
    else:
        print("sub {} success".format(IF1911))


if __name__ == '__main__':
    login()
    with Pool(processes=num) as pool:
        pool.apply_async(IF1911Process)

    subIF1911()
    while count < num:
        NotifyQueue.get()
        count += 1
    pool.close()
    pool.join()
