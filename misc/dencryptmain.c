#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <math.h>

#include "dencrypt.h"

int main() {
    char hash_output[34 + 1]; 
    encrypt(hash_output, "123456");  //     # 密码加密
    printf("%s\n", hash_output);
    printf("%d\n", verify("123456", "$H$9wXDkkAvY1XanMapy6OgnEVNS0ifOu1")); //      # 密码验证
    printf("%d\n", verify("123456", hash_output)); //      # 密码验证
}
