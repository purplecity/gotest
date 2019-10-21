//
// Created by ludongdong on 2019-10-09.
//

#ifndef QIHUO_MDUSERHANDLE_H
#define QIHUO_MDUSERHANDLE_H

#include "ctp/ThostFtdcMdApi.h"
#include "stdio.h"
#include "unistd.h"
#include <string>
#include <cstring>


class CMduserHandler : public CThostFtdcMdSpi {
private:
    CThostFtdcMdApi *m_mdApi;
public:
    void connect() {
        m_mdApi = CThostFtdcMdApi::CreateFtdcMdApi();
        m_mdApi->RegisterSpi(this);
        char url[] = "tcp://180.168.146.187:10110";
        m_mdApi->RegisterFront(url);
        m_mdApi->Init();
    }

    void login() {
        CThostFtdcReqUserLoginField t = {0};
        TThostFtdcBrokerIDType gBrokerID = "9999";
        TThostFtdcUserIDType	UserID = "";
        TThostFtdcPasswordType	Password = "";
        strcpy(t.BrokerID, gBrokerID);
        strcpy(t.UserID, UserID);
        strcpy(t.Password, Password);
        while (m_mdApi->ReqUserLogin(&t,1) != 0) {
            printf("login failed");
            sleep(1);
        }
    }

    void subscribe() {
        char **ppInstrument = new char *[50];
        char sy1[] = "IF1911";
        ppInstrument[0] = sy1;
        while(m_mdApi->SubscribeMarketData(ppInstrument,1) != 0 ) {
            printf("subscribe failed");
            sleep(1);
        }
    }

    void OnRtnDepthMarketData(CThostFtdcMarketDataField *pDepthMarketData) {
        printf("test--%lf\n",pDepthMarketData->LastPrice);
    }

};

#endif //QIHUO_MDUSERHANDLE_H
