使用postman向 127.0.0.1:8082

raw项：输入测试用例

{
    "long_url": "https://www.google.com/search?q=go正则表达式"
}
##注意：启动redis##


成功会获得:
{
    "original_url": "https://www.google.com/search?q=%E8%BF%99%E6%98%AF%E4%B8%80%E4%B8%AA%E9%87%8D%E5%AE%9A%E5%90%91%E9%A1%B5%E9%9D%A2&sca_esv=66df0cea068f4c14&sxsrf=ANbL-n6zicq5U65FNzHwGDMPbYN9HdmmYg%3A1776328101631&ei=pZ3gaYagJtypruEPks_C0AQ&biw=1745&bih=866&ved=0ahUKEwjGhoLs-fGTAxXclCsGHZKnEEoQ4dUDCBE&uact=5&oq=%E8%BF%99%E6%98%AF%E4%B8%80%E4%B8%AA%E9%87%8D%E5%AE%9A%E5%90%91%E9%A1%B5%E9%9D%A2&gs_lp=Egxnd3Mtd2l6LXNlcnAiG-i_meaYr-S4gOS4qumHjeWumuWQkemhtemdojIFEAAY7wUyBRAAGO8FMgUQABjvBTIFEAAY7wUyBRAAGO8FSJYiUABYiyFwAHgBkAEBmAG0AaAB9xqqAQQwLjI1uAEDyAEA-AEBmAIXoALNGcICDBAjGIAEGBMYJxiKBcICBRAAGIAEwgIFEC4YgATCAgcQABiABBgKwgIKEC4YgAQYQxiKBcICGRAuGIAEGEMYigUYlwUY3AQY3gQY4ATYAQHCAhQQLhiABBiXBRjcBBjeBBjfBNgBAcICCBAAGIAEGKIEwgIHEAAYgAQYDcICBxAuGIAEGA3CAgcQABiABBgMmAMAugYGCAEQARgUkgcEMC4yM6AH32ayBwQwLjIzuAfNGcIHCDAuMTMuNS41yAdlgAgA&sclient=gws-wiz-serp",
    "short_code": "2XvrLW5O",
    "short_url": "https://www.google.com/2XvrLW5O"
}

使用GET/在阅览器中输入 127.0.0.1:8082/2XvrLW5O 
重定向回 “https://www.google.com/search?q=go正则表达式”
