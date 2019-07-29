namespace go rpc
namespace php rpc

struct DataGiftPrize {
    1: i64 Id = 0
    2: string Title =""
    3: string Img =""
    4: i64 Displayorder = 0
    5: i64 Gtype = 0
    6: string Gdata = ""
}

struct DataResult {
    1: i64 code
    2: string Msg
    3: DataGiftPrize Gift
}

service LuckyService {
    DataResult DoLucky(1: i64 uid, 2: string username, 3: string ip, 4: i64 now, 5: string app, 6: string sign),
    list<DataGiftPrize> MyPrizeList(1: i64 uid, 2: string username, 3: string ip, 4: i64 now, 5: string app, 6: string sign)
}