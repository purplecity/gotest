#include "mduserhandle.h"

int main() {
    CMduserHandler *mduser = new CMduserHandler;
    mduser->connect();
    mduser->login();
    mduser->subscribe();
    getchar();
    return 0;
}