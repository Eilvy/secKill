namespace go itemcenter

struct user{
    1: i64 id;
    2: string username;
    3: string password;
    4: bool status;
}
struct product{
    1: i64 id;
}
service Itemcenter{
    void Register(user u)
    user Login(user u)
    string ParseToken(string token,i64 typ)
    string CreateToken(user u,i64 typ)
    bool SecKill(user u)
}
